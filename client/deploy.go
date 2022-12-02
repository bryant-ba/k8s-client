package client

import (
	"context"
	"fmt"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

func GetNameSpace(client *kubernetes.Clientset, ctx context.Context) {
	ns, err := client.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for name := range ns.Items {
		fmt.Println(name)
	}
}

func GetNode(client *kubernetes.Clientset, ctx context.Context) {
	nodes, err := client.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, node := range nodes.Items {
		fmt.Println("k8s node ===>", node.Name)
	}
}

func GetDeploy(client *kubernetes.Clientset, ctx context.Context, deployName, namespace string) {
	// get deploy
	deployment, err := client.AppsV1().Deployments(namespace).Get(ctx, deployName, metaV1.GetOptions{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("deployment name ===> ", deployment.Name)
	fmt.Println("deployment json ===> ", deployment)
}

func GetPods(client *kubernetes.Clientset, ctx context.Context, namespace string) {
	// get pod
	pods, err := client.CoreV1().Pods(namespace).List(ctx, metaV1.ListOptions{})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("pod Name ===> ", pods.Items[0].Status.ContainerStatuses[0].Name)
	fmt.Println("pod Image ===> ", pods.Items[0].Status.ContainerStatuses[0].Image)
	fmt.Println("pod State ===> ", pods.Items[0].Status.ContainerStatuses[0].State.Running)
}

//func GetPVC(client *kubernetes.Clientset, ctx context.Context, namespace string) {
//	pvc, err := client.CoreV1().PersistentVolumeClaims(namespace).List(ctx, metaV1.ListOptions{})
//	if err != nil {
//		log.Panic(err)
//	}
//	fmt.Println("pvc name ===> ", pvc.Unmarshal())
//}

//func UpdateDeployImage(client *kubernetes.Clientset, ctx context.Context, deployName, namespace, image string) {
//	deployment, err := client.AppsV1().Deployments(namespace).Get(ctx, deployName, metaV1.GetOptions{})
//	if err != nil {
//		log.Println(err)
//	}
//	deployment.Spec.Template.Spec.Containers[0].Image = image
//	deployment, err = client.AppsV1().Deployments(namespace).Update(ctx, deployment, metaV1.UpdateOptions{})
//}
