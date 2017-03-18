package main

import "fmt"

//import "utils"
import (
	//"inter"
	"errors"
	//"socket"
	//"runtime"
	//"time"
	//"gopkg.in/redis.v5"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/goconsole"
)

func isPositiveEvenNumber(number int) (result bool) {

	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Recoverd a panic. [index=]", v)
		}
	}()

	if number < 0 {
		panic(errors.New(" error!!!"))
	}
	if number%2 == 0 {
		return true
	}

	return

}

type Person struct {
	Name    string
	Age     uint8
	Address Addr
}

type Addr struct {
	city     string
	district string
}


type Hello struct{ Who string}
type HelloActor struct {}

func (state *HelloActor) Receive(context actor.Context){
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("Hello %v\n", msg.Who)
	}
}

func main()  {
	props := actor.FromInstance(&HelloActor{})
	pid := actor.Spawn(props)
	pid.Tell(Hello{Who : "Roger"})
	console.ReadLine()
}

//func main() {
//	fmt.Print("HHHHHHHHH\n")
//	//var sum = utils.Add(1, 2)
//	//fmt.Println(sum)
//
//	isPositiveEvenNumber(-1)
//
//	//_, ok := interface{}(inter.SortableStrings{}).(inter.Interface)
//	//fmt.Println(ok)
//
//	//socket.Init()
//
//	//for i := 0; i < 10 ; i++  {
//	//	go func() {
//	//		time.Sleep(5 * time.Second)
//	//	}()
//	//	time.Sleep(time.Second)
//	//}
//
//	go println("GO! goroutine!")
//	time.Sleep(time.Second)
//
//	//name := "Eric"
//	//go func() {
//	//	fmt.Printf("Hello, %s.\n", name)
//	//}()
//	//
//	//runtime.Gosched()
//	//name = "Harry"
//
//	names := []string{"Eric", "Harry", "WOk1ww", "woko2"}
//	for _, name := range names {
//		go func(who string) {
//			fmt.Printf("Hello, %s.\n", who)
//		}(name)
//	}
//	runtime.Gosched()
//
//	var personChan = make(chan Person, 1)
//	p1 := Person{"Harry", 32, Addr{"Beijing", "HaiDian"}}
//	fmt.Printf("p1 (1): %v\n", p1)
//	personChan <- p1
//
//	p1.Address.district = "Shijings"
//	fmt.Printf("p1 (2): %v\n", p1)
//
//	p1_copy := <-personChan
//	fmt.Printf("p1_copy: %v\n", p1_copy)
//
//	ch := make(chan int, 5)
//	sign := make(chan byte, 2)
//	go func() {
//		for i := 0; i < 5; i++ {
//			ch <- i
//			time.Sleep(1 * time.Second)
//		}
//		close(ch)
//		fmt.Println("The channel is closed.")
//		sign <- 0
//	}()
//
//	go func() {
//		for {
//			e, ok := <-ch
//			fmt.Printf("%d (%v)\n", e, ok)
//			if !ok {
//				break
//			}
//			time.Sleep(2 * time.Second)
//		}
//		fmt.Println("Done.")
//		sign <- 1
//	}()
//
//	<-sign
//	<-sign
//
//	client := redis.NewClient(&redis.Options{
//		Addr: "localhost:6379",
//		Password: "",
//		DB: 0,
//	})
//	pong, err := client.Ping().Result()
//	fmt.Println(pong, err)
//
//	//client.Subscribe()
//
//}
