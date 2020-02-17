package XuanGuBao

import (
	"io/ioutil"
	"testing"
)

func TestLiveParse(t *testing.T) {
	body, err := ioutil.ReadFile("liveParse.html")
	if err != nil{
		t.Errorf("cant read fiel err:%v",err)
	}
	result := LiveParse(body)

	requests := result.Request[0]
	if requests.Url == ""{
		t.Errorf("url not found")
	}
	item := result.Item[0].(string)

	expectedItem := "中储粮湖北分公司投放国家一次性储备玉米7万余吨保市场供应"
	if expectedItem != item{
		t.Errorf("item expected %v ,but was %v",expectedItem,item)
	}
}