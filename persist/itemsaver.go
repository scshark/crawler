package persist

import (
	"context"
	"github.com/olivere/elastic"
	"log"
	"st-crawler/engine"
)

// 存储器 goroutine去接收item channel的数据 给es存储
func ItemServer() (chan engine.Item,error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil,err
	}
	itemChan := make(chan engine.Item)
	go func() {
		itemCount := 0
		for{
			item := <-itemChan
			// save
			err := Save(client,item)
			if err != nil {
				log.Printf("save item error : %v",err)
				continue
			}
			log.Printf("item %d ; %s",itemCount,item)
			itemCount++
		}

	}()
	return itemChan,nil
}
// es存储数据  ---es7 以后一个index只能有一个type
func Save(client *elastic.Client,item engine.Item) ( err error) {

	indexService := client.Index().
		Index(item.Index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
