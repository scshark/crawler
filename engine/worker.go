package engine


import (
	"log"
	"st-crawler/fetcher"
)

func worker(r Request) (ParseResult,error){
	body, err := fetcher.Fetcher(r.Url)

	if err != nil {
		log.Printf("fetcher error %s err:%v",r.Url,err)
		return ParseResult{},err
	}

	return r.ParseFunction(r.Url,body),nil

}