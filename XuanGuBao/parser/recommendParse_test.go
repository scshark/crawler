package XuanGuBao

import (
	"st-crawler/fetcher"
	"testing"
)

func TestRecommendParse(t *testing.T) {
	bytes, err := fetcher.Fetcher("https://baoer-api.xuangubao.cn/api/v6/message/recommend_messages_page_token/610876?token=&limit=10")
	if err != nil {
		panic("fetcher error")
	}
	recommendParse(bytes)
	// t.Errorf("%T",bytes)
}
