package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"st-crawler/engine"
	"st-crawler/model"
	"testing"
)

func TestSave(t *testing.T) {


	expected := engine.Item{
		Url:"https://xuangubao.cn/article/12333",
		Id:"12333",
		Index:"xuangubao",
		Type:"stock",
		PayLoad:model.Article{
			Title:    "*ST河化：子公司为药物磷酸氯喹关键中间体主要生产厂家",
			Content:  "*ST河化公告，全资子公司南松凯博为药物磷酸氯喹、羟氯喹的关键中间体氯喹侧链和羟基氯喹侧链的主要生产厂家，南松凯博获批复工复产后，人员、设备、原材料均已准备就续，但为其提供蒸汽的企业因员工出现新冠肺炎确诊病例而被要求停产控疫，南松凯博正积极协调争取解决。目前，磷酸氯喹的临床试验结果经官方公布“对新冠肺炎有一定的诊疗效果“，羟氯喹药物的临床试验最终结果尚未公布。",
			DateTime: "2020/02/16 17:17",
			Source:   "文章来源 巨潮资讯-深交所公告",
			Stock:  []model.Stock{
				{
					Name:    "*ST河化",
					Code:  "000953.SZ",
					Price: "4.42",
					Float:   "-3.70%",
				},
			},
			Plate:[]model.Plate{
				{
					Name:    "新型病毒防治",
					Float: "-0.88%",
				},{
					Name:    "ST股",
					Float: "-0.37%",
				},
			},
		},
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		t.Errorf("create elastic client error %v",err)
	}
	err = Save(client,expected)
	if err != nil {
		t.Errorf("save error %v",err)
	}


	response, err := client.Get().
		Index(expected.Index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil{
		t.Errorf("es client get error %v",err)
	}

	result := engine.Item{}
	err = json.Unmarshal(*response.Source,&result)
	if err != nil {
		t.Errorf("json unmarsha error %v",err)
	}

	result.PayLoad,err = model.FromJson(result.PayLoad)
	if err != nil {
		t.Errorf("payload FromJson error %v",err)
	}



	// article test
	payload := result.PayLoad.(model.Article)
	result.PayLoad = nil
	expectedPayload := expected.PayLoad.(model.Article)
	expected.PayLoad = nil
	if result != expected {
		t.Errorf("result item expected %v ;\n but was %v", expected, result)
	}


	if payload.Title != expectedPayload.Title{
		t.Errorf("payload title expected %v ;\n but was %v", expectedPayload.Title, payload.Title)
	}
	if payload.Content != expectedPayload.Content {
		t.Errorf("payload content expected %v ;\n but was %v", expectedPayload.Content, payload.Content)
	}
	if payload.DateTime != expectedPayload.DateTime {
		t.Errorf("payload DateTime expected %v ;\n but was %v", expectedPayload.DateTime, payload.DateTime)
	}
	if payload.Content != expectedPayload.Content {
		t.Errorf("payload content expected %v ;\n but was %v", expectedPayload.Content, payload.Content)
	}
	if payload.Source != expectedPayload.Source {
		t.Errorf("payload Source expected %v ;\n but was %v", expectedPayload.Source, payload.Source)
	}

	for i,s := range payload.Stock{

		if s != expectedPayload.Stock[i] {
			t.Errorf("payload stock expected %v ;\n but was %v", expectedPayload.Stock[i], s)
		}
	}
	for i,p := range payload.Plate{

		if p != expectedPayload.Plate[i] {
			t.Errorf("payload plate expected %v ;\n but was %v", expectedPayload.Plate[i], p)
		}
	}


}