package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"go.etcd.io/etcd/client/v3/concurrency"

	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"

	clientV3 "go.etcd.io/etcd/client/v3"
)

var (
	addr    = flag.String("addr", "http://localhost:2379", "etcd address")
	barrier = flag.String("name", "my-test-barrier", "barrier name")
	//action    = flag.String("rw", "w", "r means acquiring road lock,w means acquiring write lock")
)

func main() {
	flag.Parse()
	//rand.NewSource(time.Now().UnixNano())

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

	b := recipe.NewDoubleBarrier(sl, *barrier, 2)

	consoleScanner := bufio.NewScanner(os.Stdin)
	for consoleScanner.Scan() {
		action := consoleScanner.Text()
		items := strings.Split(action, " ")
		switch items[0] {
		case "enter":
			err := b.Enter()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("after enter")
		case "leave":
			err := b.Leave()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("after leave")
		case "quit", "exit":
			return
		default:
			fmt.Println("unknown action")
		}
	}
}
