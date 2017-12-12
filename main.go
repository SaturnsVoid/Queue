package main

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
)

var commandQueue = make(chan string, 1) //Single Channel for handling this is all we need.
var QueueSleepInterval = struct{ min, max int }{1, 3}



func test(){
	for {
		wsi := QueueSleepInterval
		s := rand.Intn(wsi.max-wsi.min) + wsi.min
		time.Sleep(time.Duration(s) * time.Second)

		data := <-commandQueue

		fmt.Println("[",time.Now().Format(time.RFC850),"]", data)
	}
}

func main() {
	var i = 0
	go test()
	fmt.Println("Start Test.")
	commandQueue <- "Command A"
	time.Sleep(time.Duration(5) * time.Second)
	commandQueue <- "Command B"
	time.Sleep(time.Duration(5) * time.Second)
	commandQueue <- "Command C"
	for{
		time.Sleep(time.Duration(10) * time.Second)
		commandQueue <- "New Command Added [" + strconv.Itoa(i) +"]"
		i++
	}
}