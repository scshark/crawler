package engine

type Request struct {
	Url string
	ParseFunction func([]byte) ParseResult
}
type ParseResult struct {
	Request []Request
	Item []interface{}
}