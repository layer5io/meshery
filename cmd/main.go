package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path"
	"time"

	"github.com/layer5io/meshery/helpers"
	"github.com/layer5io/meshkit/utils"
	"github.com/layer5io/meshsync/pkg/broker"
	"github.com/layer5io/meshsync/pkg/broker/nats"
	"github.com/layer5io/meshsync/pkg/model"

	"github.com/layer5io/meshery/handlers"
	"github.com/layer5io/meshery/models"
	"github.com/layer5io/meshery/router"
	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"

	"github.com/vmihailenco/taskq/v3"
	"github.com/vmihailenco/taskq/v3/memqueue"
)

var (
	globalTokenForAnonymousResults string
	version                        = "Not Set"
	commitsha                      = "Not Set"
)

const (
	// DefaultProviderURL is the provider url for the "none" provider
	DefaultProviderURL = "https://meshery.layer5.io"
)

func main() {
	if globalTokenForAnonymousResults != "" {
		models.GlobalTokenForAnonymousResults = globalTokenForAnonymousResults
	}

	ctx := context.Background()

	viper.AutomaticEnv()

	viper.SetDefault("PORT", 8080)
	viper.SetDefault("ADAPTER_URLS", "")
	viper.SetDefault("BUILD", version)
	viper.SetDefault("COMMITSHA", commitsha)

	home, err := os.UserHomeDir()
	if viper.GetString("USER_DATA_FOLDER") == "" {
		if err != nil {
			logrus.Fatalf("unable to retrieve the user's home directory: %v", err)
		}
		viper.SetDefault("USER_DATA_FOLDER", path.Join(home, ".meshery", "config"))
	}
	logrus.Infof("Using '%s' to store user data", viper.GetString("USER_DATA_FOLDER"))

	if viper.GetString("KUBECONFIG_FOLDER") == "" {
		if err != nil {
			logrus.Fatalf("unable to retrieve the user's home directory: %v", err)
		}
		viper.SetDefault("KUBECONFIG_FOLDER", path.Join(home, ".kube"))
	}
	logrus.Infof("Using '%s' as the folder to look for kubeconfig file", viper.GetString("KUBECONFIG_FOLDER"))

	if viper.GetBool("DEBUG") {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.Infof("Log level: %s", logrus.GetLevel())

	adapterURLs := viper.GetStringSlice("ADAPTER_URLS")

	adapterTracker := helpers.NewAdaptersTracker(adapterURLs)
	queryTracker := helpers.NewUUIDQueryTracker()

	// Uncomment line below to generate a new UUID and force the user to login every time Meshery is started.
	// fileSessionStore := sessions.NewFilesystemStore("", []byte(uuid.NewV4().Bytes()))
	// fileSessionStore := sessions.NewFilesystemStore("", []byte("Meshery"))
	// fileSessionStore.MaxLength(0)

	QueueFactory := memqueue.NewFactory()
	mainQueue := QueueFactory.RegisterQueue(&taskq.QueueOptions{
		Name: "loadTestReporterQueue",
	})

	provs := map[string]models.Provider{}

	preferencePersister, err := models.NewMapPreferencePersister()
	if err != nil {
		logrus.Fatal(err)
	}
	defer preferencePersister.ClosePersister()

	smiResultPersister, err := models.NewBitCaskSmiResultsPersister(viper.GetString("USER_DATA_FOLDER"))
	if err != nil {
		logrus.Fatal(err)
	}
	defer smiResultPersister.CloseResultPersister()

	resultPersister, err := models.NewBitCaskResultsPersister(viper.GetString("USER_DATA_FOLDER"))
	if err != nil {
		logrus.Fatal(err)
	}
	defer resultPersister.CloseResultPersister()

	testConfigPersister, err := models.NewBitCaskTestProfilesPersister(viper.GetString("USER_DATA_FOLDER"))
	if err != nil {
		logrus.Fatal(err)
	}
	defer testConfigPersister.CloseTestConfigsPersister()

	lProv := &models.DefaultLocalProvider{
		ProviderBaseURL:        DefaultProviderURL,
		MapPreferencePersister: preferencePersister,
		ResultPersister:        resultPersister,
		SmiResultPersister:     smiResultPersister,
		TestProfilesPersister:  testConfigPersister,
	}
	lProv.Initialize()
	provs[lProv.Name()] = lProv

	cPreferencePersister, err := models.NewBitCaskPreferencePersister(viper.GetString("USER_DATA_FOLDER"))
	if err != nil {
		logrus.Fatal(err)
	}
	defer preferencePersister.ClosePersister()

	RemoteProviderURLs := viper.GetStringSlice("PROVIDER_BASE_URLS")
	for _, providerurl := range RemoteProviderURLs {
		parsedURL, err := url.Parse(providerurl)
		if err != nil {
			logrus.Error(providerurl, "is invalid url skipping provider")
			continue
		}
		cp := &models.RemoteProvider{
			RemoteProviderURL:          parsedURL.String(),
			RefCookieName:              parsedURL.Host + "_ref",
			SessionName:                parsedURL.Host,
			TokenStore:                 make(map[string]string),
			LoginCookieDuration:        1 * time.Hour,
			BitCaskPreferencePersister: cPreferencePersister,
			ProviderVersion:            "v0.3.14",
			SmiResultPersister:         smiResultPersister,
		}

		cp.Initialize()

		cp.SyncPreferences()
		defer cp.StopSyncPreferences()
		provs[cp.Name()] = cp
	}

	h := handlers.NewHandlerInstance(&models.HandlerConfig{
		Providers:              provs,
		ProviderCookieName:     "meshery-provider",
		ProviderCookieDuration: 30 * 24 * time.Hour,

		AdapterTracker: adapterTracker,
		QueryTracker:   queryTracker,

		Queue: mainQueue,

		KubeConfigFolder: viper.GetString("KUBECONFIG_FOLDER"),

		GrafanaClient:         models.NewGrafanaClient(),
		GrafanaClientForQuery: models.NewGrafanaClientWithHTTPClient(&http.Client{Timeout: time.Second}),

		PrometheusClient:         models.NewPrometheusClient(),
		PrometheusClientForQuery: models.NewPrometheusClientWithHTTPClient(&http.Client{Timeout: time.Second}),
	})

	port := viper.GetInt("PORT")
	r := router.NewRouter(ctx, h, port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// subscribing to nats
	natsClient, err := nats.New("nats://127.0.0.1:4222")
	if err != nil {
		logrus.Printf("Nats client create error %v", err)
	} else {
		msgch := make(chan *broker.Message)
		err = natsClient.SubscribeWithChannel("meshsync", "meshery", msgch)
		if err != nil {
			logrus.Printf("Error subscribing to meshsync %v", err)
		}

		go func() {
			for {
				select {
				case msg := <-msgch:
					objectJSON, _ := json.Marshal(msg.Object)
					var object model.Object
					_ = utils.Unmarshal(string(objectJSON), &object)
					// persist the object
					log.Printf("%+v\n", object)
				}
			}
		}()
	}

	go func() {
		logrus.Infof("Starting Server listening on :%d", port)
		if err := r.Run(); err != nil {
			logrus.Fatalf("ListenAndServe Error: %v", err)
		}
	}()
	<-c
	logrus.Info("Shutting down Meshery")
}
