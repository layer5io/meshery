package istio

import (
	"flag"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	kube "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type IstioClient struct {
	k8s                *kube.Clientset
	istioConfigApi     *rest.RESTClient
	istioNetworkingApi *rest.RESTClient
}

func ConfigClient() (*rest.Config, error) {
	kubeconfig := flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()
	if *kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", *kubeconfig)
	} else {
		return rest.InClusterConfig()
	}
}

func NewClient() (*IstioClient, error) {
	client := IstioClient{}
	config, err := ConfigClient()

	if err != nil {
		return nil, err
	}
	config.QPS = 100
	config.Burst = 200

	k8s, err := kube.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	client.k8s = k8s

	types := runtime.NewScheme()
	schemeBuilder := runtime.NewSchemeBuilder(
		func(scheme *runtime.Scheme) error {
			for _, kind := range istioKnownTypes {
				scheme.AddKnownTypes(*kind.groupVersion, kind.object, kind.collection)
			}
			meta_v1.AddToGroupVersion(scheme, istioConfigGroupVersion)
			meta_v1.AddToGroupVersion(scheme, istioNetworkingGroupVersion)
			return nil
		})

	err = schemeBuilder.AddToScheme(types)
	if err != nil {
		return nil, err
	}

	istioConfig := rest.Config{
		Host:    config.Host,
		APIPath: "/apis",
		ContentConfig: rest.ContentConfig{
			GroupVersion:         &istioConfigGroupVersion,
			NegotiatedSerializer: serializer.DirectCodecFactory{CodecFactory: serializer.NewCodecFactory(types)},
			ContentType:          runtime.ContentTypeJSON,
		},
		BearerToken:     config.BearerToken,
		TLSClientConfig: config.TLSClientConfig,
		QPS:             config.QPS,
		Burst:           config.Burst,
	}

	istioConfigApi, err := rest.RESTClientFor(&istioConfig)
	client.istioConfigApi = istioConfigApi
	if err != nil {
		return nil, err
	}

	istioNetworking := rest.Config{
		Host:    config.Host,
		APIPath: "/apis",
		ContentConfig: rest.ContentConfig{
			GroupVersion:         &istioNetworkingGroupVersion,
			NegotiatedSerializer: serializer.DirectCodecFactory{CodecFactory: serializer.NewCodecFactory(types)},
			ContentType:          runtime.ContentTypeJSON,
		},
		BearerToken:     config.BearerToken,
		TLSClientConfig: config.TLSClientConfig,
		QPS:             config.QPS,
		Burst:           config.Burst,
	}

	istioNetworkingApi, err := rest.RESTClientFor(&istioNetworking)
	client.istioNetworkingApi = istioNetworkingApi
	if err != nil {
		return nil, err
	}

	return &client, nil
}
