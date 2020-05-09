package myq

import (
	"fmt"
	"strconv"
	"time"
)

var runChanQ bool

func InitChanQ(jobChan <-chan string) {
	runChanQ = true
	numWorkers := 2
	for i := 0; i < numWorkers; i++ {
		go chanWorker(i, "chan_worker_"+strconv.Itoa(i), jobChan)
	}
}

func chanWorker(index int, name string, jobChan <-chan string) {
	wid := index
	wname := name
	fmt.Println("Starting channel worker ... id - ", wid, ", name - ", wname)
	
	for runChanQ {
		message := <-jobChan
		fmt.Println("Chan Worker id - ",wid," name - ",wname," got message - ",message)
		time.Sleep(3 * 1000 * time.Millisecond)
	}
}

