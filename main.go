package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

var json = `
{
	"apiVersion": "apps/v1",
	"kind": "Deployment",
	"metadata": {
		"name": "node-web-deploy",
		"labels": {
			"app": "devevopment"
		}
	},
	"spec": {
		"selector": {
			"matchLabels": {
				"run": "node-web"
			}
		},
		"replicas": 2,
		"minReadySeconds": 10,
		"strategy": {
			"rollingUpdate": {
				"maxSurge": 1,
				"maxUnavailable": 0
			},
			"type": "RollingUpdate"
		},
		"template": {
			"metadata": {
				"labels": {
					"run": "node-web"
				}
			},
			"spec": {
				"containers": [
					{
						"image": "sachinnicky/node-web:v0.1.0",
						"name": "node-web-container",
						"ports": [
							{
								"containerPort": 8090
							}
						]
					}
				]
			}
		}
	}
}`

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	decode := scheme.Codecs.UniversalDeserializer().Decode

	obj, _, err := decode([]byte(json), nil, nil)
	if err != nil {
		fmt.Printf("%#v", err)
	}

	deployment := obj.(*v1.Deployment)

	// Create Deployment
	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(deployment)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())

}
