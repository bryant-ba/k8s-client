package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	client2 "k8s-client/client"
)

func main() {

	ctx := context.Background()
	client, err := client2.Init()
	if err != nil {
		log.Panic(err)
	}
	//client2.GetNameSpace(client, ctx)
	client2.GetNode(client, ctx)
	namespace := "default"
	client2.GetPods(client, ctx, string(namespace))
	//client2.GetPVC(client, ctx, namespace)
	//client2.GetDeploy(client, ctx, deployname, namespace)
}
