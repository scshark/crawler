package XuanGuBao

import (
	"io/ioutil"
	"testing"
)

func TestLiveParse(t *testing.T) {
	body, err := ioutil.ReadFile("liveParse.txt")
	if err != nil{
		t.Errorf("cant read fiel err:%v",err)
	}
	result := LiveParse("",body)

	requests := result.Request[0]
	if requests.Url == ""{
		t.Errorf("url not found")
	}

}