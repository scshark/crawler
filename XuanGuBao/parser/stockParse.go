package XuanGuBao

import (
	"regexp"
	"st-crawler/XuanGuBao/config"
	"st-crawler/common"
	"st-crawler/engine"
	"st-crawler/model"
	"strings"
)

// datetime and source
var articleParseCompile = regexp.MustCompile(
	`<div class="meta_[a-zA-Z0-9]+"><time>([^>]+)</time>.*?[<span>]*?([^<!>]*?)[</span>]*?</div>`)
// content
var summaryParseCompile = regexp.MustCompile(
	`<pre.*class="pre_[a-zA-Z0-9]+">([^>]+)</pre>`)
// stock
var stockParseCompile = regexp.MustCompile(
	`<div class="title_[a-zA-Z0-9]+"><a href="/stock/.+".+><header>([^>]+)</header>.*<span>([0-9]+.[A-Z]*)</span>.*<div class="price_[a-zA-Z0-9]+"><span.+>([^<]+)</span>.+>([^<]+%)</span>`)
// plate
var plateParseCompile = regexp.MustCompile(
	`<div class="info_[0-9a-zA-Z]+"><header><a.+>([^<]+)</a></header>[^<]?<span[^>]+>([^<]+)</span>`)
var idParseCompile = regexp.MustCompile(
	`https://xuangubao.cn/article/([\d]+)`)
func StockParse(title string,articleUrl string, content []byte) engine.ParseResult {

	//
	parseResult := engine.ParseResult{}
	// article
	article := model.Article{}
	article.Title = title
	articleSub := extractSlice(content, articleParseCompile, 2)
	article.DateTime = articleSub[0]
	article.Source = strings.Replace(articleSub[1],"--","",-1)
	summarySub := extractSlice(content, summaryParseCompile, 1)
	article.Content = summarySub[0]
	// stock
	var stock []model.Stock
	stockSub := stockParseCompile.FindAllSubmatch(content, -1)
	if len(stockSub) > 0 {
		for _, s := range stockSub {
			mStock := model.Stock{}

			stockName := strings.Replace(string(s[1]), "\n", "", -1)
			stockName = strings.Replace(stockName, " ", "", -1)
			mStock.Name = stockName
			mStock.Code = string(s[2])
			mStock.Price = string(s[3])
			mStock.Float = string(s[4])
			stock = append(stock, mStock)
		}
	}
	article.Stock = stock

	// plate
	var plate []model.Plate
	plateSub := plateParseCompile.FindAllSubmatch(content, -1)
	if len(plateSub) > 0 {
		for _, p := range plateSub {
			mPlate := model.Plate{}
			plateName := common.RemoveAllLineSpace(string(p[1]))
			mPlate.Name = plateName
			mPlate.Float = string(p[2])
			plate = append(plate, mPlate)
		}
	}
	article.Plate = plate

	// get id by url
	articleId := extractSlice([]byte(articleUrl),idParseCompile,1)
	parseResult.Item = append(parseResult.Item, engine.Item{
		Url: articleUrl,
		Id: articleId[0],
		Index: config.EsConfig.ProIndex,
		Type: config.EsConfig.ProType,
		PayLoad:article,
	})
	parseResult.Request = append(parseResult.Request, engine.Request{
		Url: getRecommendUrl(articleId[0]),
		ParseFunction: recommendParse,
	})
	return parseResult
}
func extractSlice(content []byte, compile *regexp.Regexp, i int) []string {
	match := compile.FindSubmatch(content)
	result := make([]string, i)

	if len(match) > i {
		for k := range result {
			// 已经忽略了匹配到的第一个数据
			result[k] = string(match[k+1])
		}
		// result
	}

	return result
}
func stockParseFunction(name string) engine.ParseFunction  {
	return func(url string, bytes []byte) engine.ParseResult {
		return StockParse(name,url,bytes)
	}
}

