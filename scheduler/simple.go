package scheduler

import "st-crawler/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.WorkerChan = c
}

// Submit 
func (s *SimpleScheduler) Submit( request engine.Request)  {
	go func() {
		s.WorkerChan <- request
	}()
}