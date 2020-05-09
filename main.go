package main

import (
	"fmt"
	"testq/myq"
	"strconv"
)

func main() {
	fmt.Println("Hello!")

	custom := false
	jobChan := make(chan string, 5)

	if custom {
		for i:=0; i < 10; i++ {
			message := "hey_" + strconv.Itoa(i)
			myq.AddChunkToQueue(message)
		}
	} else {
		myq.InitChanQ(jobChan)

		for i:=0; i < 10; i++ {
			message := "hey_" + strconv.Itoa(i)
			jobChan <- message
		}
	}

	fmt.Println("Done adding all msg to Q/Chan")

	if custom {
		for myq.GetQLen() > 0 { }
	} else {
		for len(jobChan) > 0 { }
		close(jobChan)
	}

	fmt.Println("Exiting")
}
