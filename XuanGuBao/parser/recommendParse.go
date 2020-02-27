package XuanGuBao

import (
	json2 "encoding/json"
	"github.com/bitly/go-simplejson"
	"log"
	"st-crawler/XuanGuBao/config"
	"st-crawler/engine"
)

func recommendParse(_ string,body []byte) engine.ParseResult{

	notResult := engine.ParseResult{}
	json, err := simplejson.NewJson(body)
	if err != nil {
		log.Printf("json decode error : %v",err)
		return notResult
	}
	itemData ,err := json.Get("data").Get("items").Array()
	if err != nil {
		log.Printf("recommend data items not found")
		return notResult
	}
	if len(itemData) == 0 {
		return notResult
	}

	result := engine.ParseResult{}
	for _,items := range itemData{
		if  item,ok := items.(map[string]interface{}); ok{
			var articleId string
			var articleTitle string
			for k,i := range item{
				switch k {
				case "id":
					articleId = i.(json2.Number).String()
				case "title":
					articleTitle = i.(string)
				}
			}
			articleUrl := config.Web.Host + config.Web.ArticleModel +  articleId
			result.Request = append(result.Request,engine.Request{
				Url:articleUrl,
				ParseFunction: stockParseFunction(articleTitle),
			})
		}

	}
	return result

}
func getRecommendUrl(articleId string) string  {
	return "https://baoer-api.xuangubao.cn/api/v6/message/recommend_messages_page_token/" + articleId + "?token=&limit="+config.Web.RecommendNum
}
