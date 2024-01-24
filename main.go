package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"strings"
	"time"

	clientV3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

var (
	addr      = flag.String("addr", "http://localhost:2379", "etcd address")
	localName = flag.String("name", "my-test-lock", "election name")
)

func main() {
	flag.Parse()

	endpoints := strings.Split(*addr, ",")
	cli, err := clientV3.New(clientV3.Config{Endpoints: endpoints})

	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	useMutex(cli)

}

func useLock(cli *clientV3.Client) {
	sl, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer sl.Close()
	locker := concurrency.NewLocker(sl, *localName)
	log.Println("acquiring lock")
	locker.Lock()
	log.Println("acquired lock")

	time.Sleep(time.Duration(rand.Intn(30)) * time.Second)
	locker.Unlock()
	log.Println("released lock")
}

func useMutex(cli *clientV3.Client) {
	sl, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer sl.Close()
	ml := concurrency.NewMutex(sl, *localName)
	log.Printf("before acquiring lock.key:%s", ml.Key())
	if err := ml.Lock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	log.Printf("acquired lock.key:%s", ml.Key())

	time.Sleep(time.Duration(rand.Intn(30)) * time.Second)
	if err := ml.Unlock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	log.Println("released lock")
}
