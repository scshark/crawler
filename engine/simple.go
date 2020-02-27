package engine

import (
	"log"
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
