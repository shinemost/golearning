package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"

	"go.etcd.io/etcd/client/v3/concurrency"

	clientV3 "go.etcd.io/etcd/client/v3"
)

var (
	addr = flag.String("addr", "http://localhost:2379", "etcd address")
	//barrier = flag.String("name", "my-test-barrier", "barrier name")
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
	totalAccount := 5
	for i := 0; i < totalAccount; i++ {
		k := fmt.Sprintf("accts/%d", i)
		if _, err = cli.Put(context.TODO(), k, "100"); err != nil {
			log.Fatal(err)
		}
	}

	exchange := func(stm concurrency.STM) error {
		from, to := rand.Intn(totalAccount), rand.Intn(totalAccount)
		if from == to {
			return nil
		}
		fromK, toK := fmt.Sprintf("accts/%d", from), fmt.Sprintf("accts/%d", to)
		fromV, toV := stm.Get(fromK), stm.Get(toK)
		fromInt, toInt := 0, 0
		//将账户金额存入fromInt,toInt
		fmt.Sscanf(fromV, "%d", &fromInt)
		fmt.Sscanf(toV, "%d", &toInt)

		fer := fromInt / 2
		fromInt, toInt = fromInt-fer, toInt+fer

		stm.Put(fromK, fmt.Sprintf("%d", fromInt))
		stm.Put(toK, fmt.Sprintf("%d", toInt))
		return nil
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				if _, seer := concurrency.NewSTM(cli, exchange); err != nil {
					log.Fatal(seer)
				}
			}
		}()
	}
	wg.Wait()

	sum := 0
	//查出所有以accts/为前缀的key，可能是多个
	accts, err := cli.Get(context.TODO(), "accts/", clientV3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}

	for _, kv := range accts.Kvs {
		v := 0
		fmt.Sscanf(string(kv.Value), "%d", &v)
		sum += v
		log.Printf("account %s:%d", kv.Key, v)
	}

	log.Println("account sum is", sum)

}
