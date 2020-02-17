package main

import (
	XuanGuBao "st-crawler/XuanGuBao/parser"
	"st-crawler/engine"
)

func main() {

	engine.Run(engine.Request{Url:"https://xuangubao.cn/live",ParseFunction:XuanGuBao.LiveParse})

}

