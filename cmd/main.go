package main

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/layer5io/meshery/helpers"
	"net/http"
	"os"
	"os/signal"
	"path"
	"time"

	"github.com/gorilla/sessions"
	"github.com/layer5io/meshery/handlers"
	"github.com/layer5io/meshery/models"
	"github.com/layer5io/meshery/router"
	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"

	"github.com/vmihailenco/taskq"
	"github.com/vmihailenco/taskq/memqueue"
)

var globalTokenForAnonymousResults string

func main() {
	if globalTokenForAnonymousResults != "" {
		models.GlobalTokenForAnonymousResults = globalTokenForAnonymousResults
	}

	ctx := context.Background()

	viper.AutomaticEnv()

	viper.SetDefault("PORT", 8080)
	viper.SetDefault("ADAPTER_URLS", "")

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

	// Uncomment line below to generate a new UID and force the user to login every time Meshery is started.
	// fileSessionStore := sessions.NewFilesystemStore("", []byte(uuid.NewV4().Bytes()))
	// fileSessionStore := sessions.NewFilesystemStore("", []byte("Meshery"))
	// fileSessionStore.MaxLength(0)

	queueFactory := memqueue.NewFactory()
	mainQueue := queueFactory.NewQueue(&taskq.QueueOptions{
		Name: "loadTestReporterQueue",
	})

	// sessionPersister := helpers.NewFileSessionPersister(viper.GetString("USER_DATA_FOLDER"))
	// sessionPersister, err := helpers.NewBadgerSessionPersister(viper.GetString("USER_DATA_FOLDER"))
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	// sessionPersister, _ := helpers.NewMapSessionPersister()

	var prov models.Provider
	var cookieSessionStore *sessions.CookieStore

	if viper.GetBool("NO_AUTH") {
		sessionPersister, err := models.NewMapSessionPersister()
		if err != nil {
			logrus.Fatal(err)
		}
		defer sessionPersister.ClosePersister()

		resultPersister, err := models.NewBitCaskResultsPersister(viper.GetString("USER_DATA_FOLDER"))
		if err != nil {
			logrus.Fatal(err)
		}
		defer resultPersister.CloseResultPersister()

		randID, _ := uuid.NewV4()
		cookieSessionStore = sessions.NewCookieStore(randID.Bytes())
		saasBaseURL := viper.GetString("SAAS_BASE_URL")
		// if saasBaseURL == "" {
		// 	logrus.Fatalf("SAAS_BASE_URL environment variable not set.")
		// }
		prov = &models.LocalProvider{
			SessionName: "meshery",
			SaaSBaseURL: saasBaseURL,
			// SessionStore: fileSessionStore,
			SessionStore:        cookieSessionStore,
			MapSessionPersister: sessionPersister,
			ResultPersister:     resultPersister,
		}
	} else {
		sessionPersister, err := models.NewBitCaskSessionPersister(viper.GetString("USER_DATA_FOLDER"))
		if err != nil {
			logrus.Fatal(err)
		}
		defer sessionPersister.ClosePersister()

		cookieSessionStore = sessions.NewCookieStore([]byte("Meshery"))
		saasBaseURL := viper.GetString("SAAS_BASE_URL")
		if saasBaseURL == "" {
			logrus.Fatalf("SAAS_BASE_URL environment variable not set.")
		}
		cp := &models.CloudProvider{
			SaaSBaseURL:   saasBaseURL,
			RefCookieName: "meshery_ref",
			SessionName:   "meshery",
			// SessionStore: fileSessionStore,
			SessionStore:            cookieSessionStore,
			SaaSTokenName:           "meshery_saas",
			LoginCookieDuration:     1 * time.Hour,
			BitCaskSessionPersister: sessionPersister,
		}
		cp.SyncPreferences()
		defer cp.StopSyncPreferences()
		prov = cp
	}

	h := handlers.NewHandlerInstance(&models.HandlerConfig{
		// SaaSBaseURL: saasBaseURL,

		// RefCookieName: "meshery_ref",

		// SessionName: "meshery",
		// // SessionStore: fileSessionStore,
		// SessionStore: cookieSessionStore,

		// SaaSTokenName: "meshery_saas",

		Provider: prov,

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

	// go func() {
	// 	err := mainQueue.Consumer().Start()
	// 	if err != nil {
	// 		logrus.Fatal(err)
	// 	}
	// }()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		logrus.Infof("Starting Server listening on :%d", port)
		if err := r.Run(); err != nil {
			logrus.Fatalf("ListenAndServe Error: %v", err)
		}
	}()
	<-c
	logrus.Info("Shutting down Meshery")
}
