package scheduler

import (
	"st-crawler/engine"
)

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (q *QueueScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}
func (q *QueueScheduler) WorkerReady (c chan engine.Request){
	q.workerChan <- c
}
func (q *QueueScheduler)Run(){
	q.workerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerChanQ []chan engine.Request

		for {

			var activeRequest engine.Request
			var activeWorkerChan chan engine.Request
			if len(requestQ) > 0 && len(workerChanQ) > 0 {
				activeRequest = requestQ[0]
				activeWorkerChan = workerChanQ[0]
			}
			select {
			case r:= <-q.requestChan:
				requestQ = append(requestQ,r)
			case w:= <-q.workerChan:
				workerChanQ = append(workerChanQ,w)
			case activeWorkerChan <- activeRequest:
				requestQ = requestQ[1:]
				workerChanQ = workerChanQ[1:]
			}
		}
	}()
}


