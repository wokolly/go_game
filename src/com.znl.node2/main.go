package main

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/examples/remoteactivate/messages"
	"github.com/AsynkronIT/protoactor-go/remote"
	"runtime"
	"github.com/AsynkronIT/goconsole"
)

type helloActor struct {

}

func (*helloActor) Receive(ctx actor.Context){
	switch ctx.Message().(type) {
	case *messages.HelloRequest:
		ctx.Respond(&messages.HelloResponse{
			Message : "Hello from remote node",
		})
	}
}

func newHelloActor() actor.Actor {
	return &helloActor{}
}

func init()  {
	remote.Register("hello", actor.FromProducer(newHelloActor))
	remote.Register("hello2", actor.FromProducer(newHelloActor))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	remote.Start("127.0.0.1:8080")

	console.ReadLine()
}
