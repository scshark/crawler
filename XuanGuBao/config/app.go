package config

type WebConfig struct {
	Host string
	Domain string
	ArticleModel string
	RecommendNum string
}

var Web = WebConfig{
	Host:"https://xuangubao.cn",
	Domain:"xuangubao.cn",
	ArticleModel: "/article/",
	RecommendNum: "50",
}