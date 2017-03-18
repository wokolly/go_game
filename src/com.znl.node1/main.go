package main

import (
	"time"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/AsynkronIT/protoactor-go/examples/remoteactivate/messages"
	"fmt"
	"github.com/AsynkronIT/goconsole"
)

func main() {
	timeout := 5 * time.Second
	remote.Start("127.0.0.1:8081")
	pid, _ := remote.SpawnNamed("127.0.0.1:8080", "remote", "hello", timeout)
	res, _ := pid.RequestFuture(&messages.HelloRequest{}, timeout).Result()
	response := res.(*messages.HelloResponse)
	fmt.Printf("Response from remote %v", response.Message)

	pid2, _ := remote.SpawnNamed("127.0.0.1:8080", "remote", "hello2", timeout)
	res2, _ := pid2.RequestFuture(&messages.HelloRequest{}, timeout).Result()
	response2 := res2.(*messages.HelloResponse)
	fmt.Printf("Response from remote2 %v", response2.Message)

	console.ReadLine()
}
