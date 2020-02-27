package config

type webConfig struct {
	Host string
	Domain string
	ArticleModel string
	RecommendNum string
}
type esConfig struct {
	ProIndex string
	ProType string
	TestIndex string
	TestType string
}
var Web = webConfig{
	Host:"https://xuangubao.cn",
	Domain:"xuangubao.cn",
	ArticleModel: "/article/",
	RecommendNum: "50",
}
var EsConfig = esConfig{
	ProIndex:"xuangubao",
	ProType:"stock",
	TestIndex:"xuangubao_test",
	TestType:"stock",
}