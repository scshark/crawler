package XuanGuBao

import (
	"io/ioutil"
	"st-crawler/model"
	"testing"
)

func TestStockParse(t *testing.T) {

	bytes, err := ioutil.ReadFile("stockParse.txt")
	if err != nil {
		panic("read file err")
	}
	result := StockParse("*ST河化：子公司为药物磷酸氯喹关键中间体主要生产厂家", bytes)

	if len(result.Item) != 3 {
		t.Errorf("item should contain 3 element. but was %v", result.Item)
	}

	// article test
	articleItem := result.Item[0].(model.Article)
	expectedArticleItem := model.Article{
		Title:    "*ST河化：子公司为药物磷酸氯喹关键中间体主要生产厂家",
		Content:  "*ST河化公告，全资子公司南松凯博为药物磷酸氯喹、羟氯喹的关键中间体氯喹侧链和羟基氯喹侧链的主要生产厂家，南松凯博获批复工复产后，人员、设备、原材料均已准备就续，但为其提供蒸汽的企业因员工出现新冠肺炎确诊病例而被要求停产控疫，南松凯博正积极协调争取解决。目前，磷酸氯喹的临床试验结果经官方公布“对新冠肺炎有一定的诊疗效果“，羟氯喹药物的临床试验最终结果尚未公布。",
		DateTime: "2020/02/16 17:17",
		Source:   "文章来源 巨潮资讯-深交所公告",
	}
	if articleItem != expectedArticleItem {
		t.Errorf("article expected %v ;\n but was %v", expectedArticleItem, articleItem)
	}
	// stock test
	stockItems := result.Item[1].([]model.Stock)
	stockItem := stockItems[0]
	expectedStockItem := model.Stock{
		Name:    "*ST河化",
		Code:  "000953.SZ",
		Price: "4.42",
		Float:   "-3.70%",
	}
	if stockItem != expectedStockItem {
		t.Errorf("stock expected %v ;\n but was %v", expectedStockItem, stockItem)
	}

	plateItems := result.Item[2].([]model.Plate)
	plateItem := plateItems[0]
	expectedPlateItem := model.Plate{
		Name:    "新型病毒防治",
		Float: "-0.88%",
	}
	if plateItem != expectedPlateItem {
		t.Errorf("plate expected %v ;but was %v", expectedPlateItem, plateItem)
	}
	// fmt.Printf("%s",parse.Item)
}
