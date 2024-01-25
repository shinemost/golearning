package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"

	clientV3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

var (
	addr      = flag.String("addr", "http://localhost:2379", "etcd address")
	localName = flag.String("name", "my-test-lock", "election name")
	//action    = flag.String("rw", "w", "r means acquiring road lock,w means acquiring write lock")
)

func main() {
	flag.Parse()
	rand.NewSource(time.Now().UnixNano())

	endpoints := strings.Split(*addr, ",")
	cli, err := clientV3.New(clientV3.Config{Endpoints: endpoints})

	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	sl, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer sl.Close()

	ml := recipe.NewRWMutex(sl, *localName)

	consoleScanner := bufio.NewScanner(os.Stdin)
	for consoleScanner.Scan() {
		action := consoleScanner.Text()
		switch action {
		case "w":
			testWriteLocker(ml)
		case "r":
			testReadLocker(ml)
		default:
			fmt.Println("unknown action")
		}

	}

}

func testWriteLocker(mutex *recipe.RWMutex) {

	log.Println("acquiring write lock")
	err := mutex.Lock()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("acquired write lock")

	time.Sleep(time.Duration(rand.Intn(30)) * time.Second)
	err = mutex.Unlock()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("released write lock")
}

func testReadLocker(mutex *recipe.RWMutex) {

	log.Println("acquiring read lock")
	err := mutex.RLock()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("acquired read lock")

	time.Sleep(time.Duration(rand.Intn(30)) * time.Second)
	err = mutex.RUnlock()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("released read lock")
}
