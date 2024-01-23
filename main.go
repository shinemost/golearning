package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	clientV3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

var (
	nodeID    = flag.Int("id", 0, "node ID")
	addr      = flag.String("addr", "http://localhost:2379", "etcd address")
	electName = flag.String("name", "my-test-elect", "election name")
)

func main() {
	flag.Parse()

	endpoints := strings.Split(*addr, ",")
	cli, err := clientV3.New(clientV3.Config{Endpoints: endpoints})

	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	session, err := concurrency.NewSession(cli)
	defer session.Close()

	el := concurrency.NewElection(session, *electName)

	consoleScanner := bufio.NewScanner(os.Stdin)

	for consoleScanner.Scan() {
		action := consoleScanner.Text()
		switch action {
		case "elect":
			go elect(el, *electName)
		case "proclaim":
			proclaim(el, *electName)
		case "resign":
			resign(el)
		case "watch":
			go watch(el, *electName)
		case "query":
			query(el, *electName)
		case "rev":
			rev(el, *electName)
		default:
			fmt.Println("unknown action")
		}
	}

}

var count int

func elect(el *concurrency.Election, electName string) {
	log.Println("campaigning for ID:", *nodeID)
	if err := el.Campaign(context.Background(), fmt.Sprintf("value-%d-%d", *nodeID, count)); err != nil {
		log.Println(err)
	}
	log.Println("campaigned for ID:", *nodeID)
	count++
}

func proclaim(el *concurrency.Election, electName string) {
	log.Println("proclaiming for ID:", *nodeID)
	if err := el.Proclaim(context.Background(), fmt.Sprintf("value-%d-%d", *nodeID, count)); err != nil {
		log.Println(err)
	}
	log.Println("proclaimed for ID:", *nodeID)
	count++
}

func resign(el *concurrency.Election) {
	log.Println("resign for ID:", *nodeID)
	if err := el.Resign(context.TODO()); err != nil {
		log.Println(err)
	}
	log.Println("resigned for ID:", *nodeID)
}

func query(el *concurrency.Election, electName string) {
	resp, err := el.Leader(context.Background())
	if err != nil {
		log.Printf("failed to get the current leader:%v", err)
	}
	log.Println("current leader:", string(resp.Kvs[0].Key), string(resp.Kvs[0].Value))
}

func rev(el *concurrency.Election, electName string) {
	rev := el.Rev()
	log.Println("current rev:", rev)
}

func watch(el *concurrency.Election, electName string) {
	ch := el.Observe(context.TODO())
	log.Println("start to watch for ID:", *nodeID)
	for i := 0; i < 10; i++ {
		resp := <-ch
		log.Println("leader changed to", string(resp.Kvs[0].Key), string(resp.Kvs[0].Value))
	}
}
