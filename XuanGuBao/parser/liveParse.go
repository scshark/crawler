package XuanGuBao

import (
	"regexp"
	"st-crawler/XuanGuBao/config"
	"st-crawler/engine"
	"strings"
)

var parseCompile = regexp.MustCompile(`<header.+class="title_[0-9a-zA-Z]+"><a.+href="([^"]+)".+>([^<]+)</a>.+</header>`)

func LiveParse(content []byte) engine.ParseResult {
	header := parseCompile.FindAllSubmatch(content, -1)
	parseResult := engine.ParseResult{}
	for _, h := range header {
		title := strings.Replace(strings.Replace(string(h[2])," ","",-1),"\n","",-1)
		resUrl := string(h[1])
		recommendUrl := strings.Split(resUrl,"/")

		if strings.Index(resUrl, config.Web.Domain) == -1 {
			resUrl = config.Web.Host + resUrl
		}
		parseResult.Item = append(parseResult.Item, title)
		parseResult.Request = append(parseResult.Request, engine.Request{
			Url: resUrl,
			ParseFunction: func(bytes []byte) engine.ParseResult {
				return StockParse(title, bytes)
			},
		})
		if len(recommendUrl) > 2 {
			recUrl := "https://baoer-api.xuangubao.cn/api/v6/message/recommend_messages_page_token/" + string(recommendUrl[2])  + "?token=&limit="+config.Web.RecommendNum
			parseResult.Request = append(parseResult.Request, engine.Request{
				Url: recUrl,
				ParseFunction: recommendParse,
			})
		}
	}
	return parseResult
}
