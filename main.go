package main

import (
	XuanGuBao "st-crawler/XuanGuBao/parser"
	"st-crawler/engine"
	"st-crawler/persist"
	"st-crawler/scheduler"
)

func main() {

	saveItemsChan, err := persist.ItemServer()

	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		WorkerCount: 30,
		Scheduler:   &scheduler.SimpleScheduler{},
		ItemChan:saveItemsChan,
	}
	e.Run(engine.Request{Url: "https://xuangubao.cn/live", ParseFunction: XuanGuBao.LiveParse})

}
