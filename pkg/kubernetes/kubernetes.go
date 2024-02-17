package kubernetes

import (
	"fmt"
	"os"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetRESTConfig(kubeconfig string, kubeconfigContext string) (*rest.Config, string, error) {
	if os.Getenv("KUBERNETES_SERVICE_HOST") == "" {
		if kubeconfig == "" {
			kubeconfig = clientcmd.NewDefaultClientConfigLoadingRules().GetDefaultFilename()
		}

		configLoader := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			&clientcmd.ClientConfigLoadingRules{
				ExplicitPath: kubeconfig,
			},
			&clientcmd.ConfigOverrides{
				CurrentContext: kubeconfigContext,
			},
		)

		config, err := configLoader.ClientConfig()
		if err != nil {
			return nil, "", fmt.Errorf("error loading kubeconfig: %v", err)
		}

		namespace, _, err := configLoader.Namespace()
		if err != nil {
			return nil, "", fmt.Errorf("error getting namespace from kubeconfig: %v", err)
		}

		return config, namespace, nil
	} else {
		config, err := rest.InClusterConfig()
		if err != nil {
			return nil, "", fmt.Errorf("error loading in-cluster kubeconfig: %v", err)
		}

		return config, "", nil
	}
}
