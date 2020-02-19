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
	MakeRequestChan() chan Request
	Run()
	ReadyNotifier
}
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (c *ConcurrentEngine) Run(seed ...Request) {

	c.Scheduler.Run()
	// make channel in and out
	out := make(chan ParseResult)

	// make worker by workerCount
	for i := 0; i < c.WorkerCount; i++ {
		createWorker(c.Scheduler.MakeRequestChan(),out,c.Scheduler)
	}


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
func createWorker(in chan Request,out chan ParseResult,r ReadyNotifier) {
	// make a goroutine
	// in := make(chan Request)
	go func() {
		for {
			r.WorkerReady(in)
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
