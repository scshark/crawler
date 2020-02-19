package main

import (
	"st-crawler/XuanGuBao/parser"
	"st-crawler/engine"
	"st-crawler/scheduler"
)

func main() {

	e := engine.ConcurrentEngine{
		WorkerCount: 30,
		Scheduler:   &scheduler.SimpleScheduler{},
	}
	e.Run(engine.Request{Url: "https://xuangubao.cn/live", ParseFunction: XuanGuBao.LiveParse})

}
