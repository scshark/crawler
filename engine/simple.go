package engine

import (
	"log"
	"st-crawler/fetcher"
)

type SimpleEngine struct {}

func (SimpleEngine) Run(seeds ...Request)  {
	var requests []Request
	for _,s := range seeds{
		requests = append(requests,s)
	}

	for len(requests) > 0{
		r := requests[0]
		requests = requests[1:]

		result, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests,result.Request...)

		for _,item := range result.Item{
			log.Printf("%s",item)
		}
	}
	
}
func worker(r Request) (ParseResult,error){
	// log.Printf("URL : %s",r.Url)
	body, err := fetcher.Fetcher(r.Url)

	if err != nil {
		log.Printf("fetcher error %s err:%v",r.Url,err)
		return ParseResult{},err
	}

	return r.ParseFunction(body),nil

}