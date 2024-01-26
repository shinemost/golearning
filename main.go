package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"

	clientV3 "go.etcd.io/etcd/client/v3"
)

var (
	addr      = flag.String("addr", "http://localhost:2379", "etcd address")
	localName = flag.String("name", "my-test-queue", "queue name")
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

	q := recipe.NewQueue(cli, *localName)

	consoleScanner := bufio.NewScanner(os.Stdin)
	for consoleScanner.Scan() {
		action := consoleScanner.Text()
		items := strings.Split(action, " ")
		switch items[0] {
		case "push":
			if len(items) != 2 {
				fmt.Println("must set value to push")
				continue
			}
			err := q.Enqueue(items[1])
			if err != nil {
				log.Fatal(err)
			}
		case "pop":
			v, err := q.Dequeue()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(v)
		case "quit", "exit":
			return
		default:
			fmt.Println("unknown action")
		}
	}
}
