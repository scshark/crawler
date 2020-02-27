package engine

type ConcurrentEngine struct {
	WorkerCount int
	Scheduler   Scheduler
	ItemChan    chan Item
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
type Item struct {
	Id string
	Url string
	Index string
	Type string
	PayLoad interface{}
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
		if isDuplicate(s.Url) {
			continue
		}
		c.Scheduler.Submit(s)
	}
	// engine start
	for {
		result := <- out
		for _,item := range result.Item{
			// log.Printf("%s",item)
			go func() {
				c.ItemChan <- item
			}()
		}
		for _,request := range result.Request  {
			if isDuplicate(request.Url) {
				continue
			}
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

var DuplicateUrl = make(map[string]bool)

func isDuplicate(Url string) bool {
	if _, ok := DuplicateUrl[Url]; ok {
		return true
	}
	DuplicateUrl[Url] = true
	return false
}
