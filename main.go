package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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

	q := recipe.NewPriorityQueue(cli, *localName)

	consoleScanner := bufio.NewScanner(os.Stdin)
	for consoleScanner.Scan() {
		action := consoleScanner.Text()
		items := strings.Split(action, " ")
		switch items[0] {
		case "push":
			if len(items) != 3 {
				fmt.Println("must set value and priority to push")
				continue
			}
			atom, err := strconv.Atoi(items[2])
			if err != nil {
				fmt.Println("must set uint16 to priority")
			}
			err = q.Enqueue(items[1], uint16(atom))
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
