package client

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func Init() (*kubernetes.Clientset, error) {
	var kubeConfig *string
	// 从当前系统环境中读取家目录，然后拼接config 路径
	// 或者直接给一个kube config的绝对路径字符串也可
	if home := HomeDir(); home != "" {
		kubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config-balin"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeConfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// uses the current context get restConfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		log.Panic(err)
	}

	// 创建clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientSet, nil
}

func HomeDir() string {
	// linux
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	// windows
	return os.Getenv("USERPROFILE")
}
