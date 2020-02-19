package engine

import (
	"log"
)

type ConcurrentEngine struct {
	WorkerCount int
	Scheduler   Scheduler
}
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (c *ConcurrentEngine) Run(seed ...Request) {

	c.Scheduler.Run()
	// make channel in and out
	// in := make(chan Request)
	out := make(chan ParseResult)

	// make worker by workerCount
	for i := 0; i < c.WorkerCount; i++ {
		createWorker(out,c.Scheduler)
	}

	// configure worker channel for scheduler

	// c.Scheduler.ConfigureMasterWorkerChan(in)
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
func createWorker(out chan ParseResult,s Scheduler) {
	// make a goroutine
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
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
