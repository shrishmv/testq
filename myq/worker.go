package myq

import (
	"fmt"
	"sync"
	"reflect"
	"container/list"
	"strconv"
	"time"
)

var once sync.Once
var localq Queue
var runQ bool

func getLocalQ() Queue {
	once.Do(func() {
		localq = NewQ()
		initQ()
	})
	return localq
}

func initQ() {
	runQ = true
	numWorkers := 2
	for i := 0; i < numWorkers; i++ {
		go queueWorker(i, "worker_"+strconv.Itoa(i))
	}
}

func queueWorker(index int, name string) {

	wid := index
	wname := name
	q := getLocalQ()

	fmt.Println("Starting worker ... id - ", wid, ", name - ", wname)

	for runQ {
		for q.Len() > 0 {
			e := q.Remove()
			if el, ok := e.(*list.Element); ok {
				if message, ok := el.Value.(string); ok {
					fmt.Println("Worker id - ",wid," name - ",wname," got message - ",message)
					time.Sleep(3 * 1000 * time.Millisecond)
				} else {
					fmt.Println("Error: Unwanted object type - ", reflect.TypeOf(el.Value).String())
				}
			} else {
				fmt.Println("Error: Not list.Element object type - ", reflect.TypeOf(e).String())
			}
		}
		//fmt.Println("No new messages to processes, waiting for 1 sec ... id - ", wid, ", name - ", wname)
		time.Sleep(1000 * time.Millisecond)
	}
}

// AddChunkToQueue - add chunk info to q
func AddChunkToQueue(message string) {
	q := getLocalQ()
	q.Add(message)
}

func GetQLen() int {
	q := getLocalQ()
	return q.Len()
}