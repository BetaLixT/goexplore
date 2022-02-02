package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"
	"time"
)

func main() {
	
	var wg sync.WaitGroup
	
	go func() {
		defer wg.Done()
		serve()
	}()
	time.Sleep(100000)
	echo("test")
	wg.Wait()
}

func serve() {
	msg := new(Message)

	rpc.Register(msg)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Error listening...", e)
	}
	http.Serve(l, nil)
}

func echo (msg string) {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("Error dialing...", err)
	}
	response := ""
	err = client.Call("Message.Echo", &msg, &response)
	if err != nil {
		log.Fatal("Error envoking", err)
	}
	fmt.Println(response)
}
