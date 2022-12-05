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

	client2.GetNode(client, ctx)
	namespace := "iot"
	client2.GetPods(client, ctx, string(namespace))
}
