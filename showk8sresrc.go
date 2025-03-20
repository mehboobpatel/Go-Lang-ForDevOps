package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Load kubeconfig file (for local Minikube)
	kubeconfig := filepath.Join("/home/maheboob/.kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal("Error loading kubeconfig:", err)
	}

	// Create Kubernetes client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("Error creating Kubernetes client:", err)
	}

	// List all Pods in the default namespace
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("Error listing pods:", err)
	}

	deployments, err := clientset.AppsV1().Deployments("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("Error listing deployments:", err)
	}

	service, err := clientset.CoreV1().Services("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("Error listing service:", err)
	}

	configmap, err := clientset.CoreV1().ConfigMaps("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("Error listing confgimaps:", err)
	}

	secrets, err := clientset.CoreV1().Secrets("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("Error listing secrets:", err)
	}

	Namespace, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("Error listing Namespace:", err)
	}

	Nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("Error listing Nodes:", err)
	}

	// Print Pod names
	for _, pod := range pods.Items {
		fmt.Println("Pod Name:", pod.Name)
	}

	for _, deployments := range deployments.Items {
		fmt.Println("Deployments Name:", deployments.Name)
	}

	for _, service := range service.Items {
		fmt.Println("Service Name:", service.Name)
	}

	for _, configmap := range configmap.Items {
		fmt.Println("Configmaps Name:", configmap.Name)
	}

	for _, secrets := range secrets.Items {
		fmt.Println("Secrets Name:", secrets.Name)
	}

	for _, Namespace := range Namespace.Items {
		fmt.Println("Namespaces Name:", Namespace.Name)
	}

	for _, Nodes := range Nodes.Items {
		fmt.Println("Nodes Name:", Nodes.Name)
	}
}
