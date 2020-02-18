package engine

import (
	"log"
	"time"
)

type ConcurrentEngine struct {
	WorkerCount int
	Scheduler   Scheduler
}
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (c *ConcurrentEngine) Run(seed ...Request) {

	// make channel in and out
	in := make(chan Request)
	out := make(chan ParseResult)
	// make worker by workerCount
	for i := 0; i < c.WorkerCount; i++ {
		c.createWorker(in, out)
	}
	// configure worker channel for scheduler

	c.Scheduler.ConfigureMasterWorkerChan(in)
	// submit request to scheduler
	for _, s := range seed {
		c.Scheduler.Submit(s)
	}
	// engine start
	for {
		result := <- out
		for _,item := range result.Item{
			log.Printf("%s",item)
		}
		for _,request := range result.Request  {
			c.Scheduler.Submit(request)
		}
	}
	// get parser result of out channel
	// print of result items
	// send result request to channel in
}
func (ConcurrentEngine) createWorker(in chan Request, out chan ParseResult) {
	// make a goroutine
	go func() {
		for {
			requests := <- in
			result, err := worker(requests)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
	// get request of in channel
	// run fetcher and parserFun by worker fun
	// if result not error
	// send result to out channel
	// else continue

}
