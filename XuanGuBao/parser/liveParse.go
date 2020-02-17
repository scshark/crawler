package XuanGuBao

import (
	"regexp"
	"st-crawler/engine"
	"strings"
)

// <header class="title_3EhHc"><a href="/article/608146" target="_blank" class="link_xsH6g">
// 央行：金融市场短期冲击之后会回到长期的基本面上来
// </a> <!----></header>
var parseCompile = regexp.MustCompile(`<header.+class="title_[0-9a-zA-Z]+"><a.+href="([^"]+)".+>([^<]+)</a>.+</header>`)

func LiveParse(content []byte) engine.ParseResult {
	header := parseCompile.FindAllSubmatch(content, -1)
	parseResult := engine.ParseResult{}
	for _, h := range header {
		title := strings.Replace(strings.Replace(string(h[2])," ","",-1),"\n","",-1)
		resUrl := string(h[1])
		if strings.Index(resUrl, "xuangubao.cn") == -1 {
			resUrl = "https://xuangubao.cn" + resUrl
		}
		parseResult.Item = append(parseResult.Item, title)
		parseResult.Request = append(parseResult.Request, engine.Request{
			Url: resUrl,
			ParseFunction: func(bytes []byte) engine.ParseResult {
				return StockParse(title, bytes)
			},
		})
	}
	return parseResult
}
