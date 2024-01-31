package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	clientV3 "go.etcd.io/etcd/client/v3"

	"go.etcd.io/etcd/client/v3/naming/endpoints"

	"github.com/gin-gonic/gin"
)

const (
	MyService = "my-service"
	MyEtcdURL = "http://localhost:2379"
)

func main() {
	// 接收命令行指定的 grpc 服务端口
	var port int
	flag.IntVar(&port, "port", 8080, "port")
	flag.Parse()
	addr := fmt.Sprintf("http://localhost:%d", port)

	mux := gin.Default()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 注册 grpc 服务节点到 etcd 中
	go registerEndPointToEtcd(ctx, addr)

	mux.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	// 注册服务后启动HTTP服务器

	if err := mux.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

func registerEndPointToEtcd(ctx context.Context, addr string) {
	// 创建 etcd 客户端
	etcdClient, _ := clientV3.NewFromURL(MyEtcdURL)
	etcdManager, _ := endpoints.NewManager(etcdClient, MyService)

	// 创建一个租约，每隔 10s 需要向 etcd 汇报一次心跳，证明当前节点仍然存活
	var ttl int64 = 100
	lease, _ := etcdClient.Grant(ctx, ttl)

	// 添加注册节点到 etcd 中，并且携带上租约 id
	_ = etcdManager.AddEndpoint(ctx, fmt.Sprintf("%s/%s", MyService, addr), endpoints.Endpoint{Addr: addr}, clientV3.WithLease(lease.ID))

	// 每隔 5 s进行一次延续租约的动作
	for {
		select {
		case <-time.After(5 * time.Second):
			// 续约操作
			resp, _ := etcdClient.KeepAliveOnce(ctx, lease.ID)
			fmt.Printf("keep alive resp: %+v", resp)
		case <-ctx.Done():
			return
		}
	}
}
