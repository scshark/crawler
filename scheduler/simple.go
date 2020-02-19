package scheduler

import "st-crawler/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) MakeRequestChan() chan engine.Request {
	return s.WorkerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
	// panic("implement me")
}

func (s *SimpleScheduler) Run() {

	s.WorkerChan = make(chan engine.Request)
}

// Submit 
func (s *SimpleScheduler) Submit( request engine.Request)  {
	go func() {
		s.WorkerChan <- request
	}()
}