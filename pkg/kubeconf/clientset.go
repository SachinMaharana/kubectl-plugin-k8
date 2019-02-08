package kubeconf

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

// NewClientSet creates a client based on the current k8 context
func NewClientSet() (*kubernetes.Clientset, error) {
	config, err := getKubeConfig()
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}

func getKubeConfig() (*rest.Config, error) {
	var kubeconfig string
	if os.Getenv("KUBECONFIG") != "" {
		kubeconfig = os.Getenv("KUBECONFIG")
	} else if home := homeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	} else {
		fmt.Println("Parsing kubeconfig failed, please set KUBECONFIG env var")
		os.Exit(1)
	}

	if _, err := os.Stat(kubeconfig); err != nil {
		// kubeconfig doesn't exist
		fmt.Printf("%s does not exist - please make sure you have a kubeconfig configured.\n", kubeconfig)
		panic(err.Error())
	}

	return clientcmd.BuildConfigFromFlags("", kubeconfig)
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
