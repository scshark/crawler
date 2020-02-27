package XuanGuBao

import (
	"regexp"
	"st-crawler/XuanGuBao/config"
	"st-crawler/common"
	"st-crawler/engine"
	"strings"
)

var parseCompile = regexp.MustCompile(`<header.+class="title_[0-9a-zA-Z]+"><a.+href="([^"]+)".+>([^<]+)</a>.+</header>`)

func LiveParse(_ string,content []byte) engine.ParseResult {
	header := parseCompile.FindAllSubmatch(content, -1)
	parseResult := engine.ParseResult{}
	for _, h := range header {

		title := common.RemoveAllLineSpace(string(h[2]))
		resUrl := string(h[1])
		recommendUrl := strings.Split(resUrl,"/")

		parseResult.Request = append(parseResult.Request, engine.Request{
			Url: config.Web.Host + resUrl,
			ParseFunction: stockParseFunction(title),
		})
		if len(recommendUrl) > 2 {
			// 同时获取推荐文章
			parseResult.Request = append(parseResult.Request, engine.Request{
				Url: getRecommendUrl(recommendUrl[2]),
				ParseFunction: recommendParse,
			})
		}
	}
	return parseResult
}
