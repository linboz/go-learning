package main

import (
	"fmt"
	"sync"
	"time"
)

var c chan string
var w sync.WaitGroup
func reader(){
	msg := <-c
	fmt.Println("I am reader, ",msg)
	w.Done()
}

func main(){
	c = make(chan string)
	w.Add(1)
	go reader()
	fmt.Println("begin sleep")
	time.Sleep(time.Second * 3)
	c <- "hello"
	w.Wait()
	//time.Sleep(time.Second * 1)
}
