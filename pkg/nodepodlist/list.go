package nodepodlist

import (
	"fmt"
	"os"

	"github.com/SachinMaharana/kubectl-plugin-k8/pkg/kubeconf"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// List prints the collected pods and nodes on the cluster
func List(args []string, showNodePods bool) {
	podList, nodeList := getPodsAndNodes()

	for _, node := range nodeList.Items {
		fmt.Println("Node Name:", node.Name)
	}

	for _, pod := range podList.Items {
		fmt.Println("Pod Name:", pod.Name)
	}

}
// getPodsAndNodes interacts with the api to get the pods and nodes
func getPodsAndNodes() (*corev1.PodList, *corev1.NodeList) {
	clientset, err := kubeconf.NewClientSet()
	if err != nil {
		fmt.Printf("Error connecting to Kubernetes: %v\n", err)
		os.Exit(1)
	}

	nodeList, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error listing Nodes: %v\n", err)
		os.Exit(2)
	}

	podList, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error listing Pods: %v\n", err)
		os.Exit(3)
	}

	return podList, nodeList
}
