package main

import (
	"io/ioutil"
	"log"
	"github.com/ericchiang/k8s"

	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	"context"
	"fmt"
	"encoding/json"
)

func main() {

	// read keys
	data, err := ioutil.ReadFile("k8s/keys/test.json")
	if err != nil {
		log.Fatalf("Read Key Fail, %s\n", err)
	}

	// josn to config use yaml don't unmarshal unit8
	// https://github.com/ericchiang/k8s/issues/81
	var config k8s.Config
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalf("Load Config JSON Fail, %s\n" ,err)
	}

	// register client
	client, err := k8s.NewClient(&config)
	if err != nil {
		log.Fatalf("Register Client Fail, %s\n" ,err)
	}

	// Call client.list return to nodes
	var nodes corev1.NodeList
	if err := client.List(context.Background(), "", &nodes); err != nil {
		log.Fatalf("Call K8s client.list Fail, %s\n", err)
	}

	// action nodes
	for _, node := range nodes.Items {
		fmt.Printf("node=%q schedulable=%t\n", *node.Metadata.Name, !*node.Spec.Unschedulable)
	}

	// action pods
	var pods corev1.PodList
	if err := client.List(context.Background(), k8s.AllNamespaces, &pods); err != nil {
		log.Fatal(err)
	}
	for _, pod := range pods.Items {
		fmt.Printf("namespace[%q]\tpod[%q]\n", *pod.Metadata.Namespace, *pod.Metadata.Name)
	}

}
