package engine

type ParseFunction  func(url string,contents []byte) ParseResult
type Request struct {
	Url string
	ParseFunction ParseFunction
}
type ParseResult struct {
	Request []Request
	Item []Item
}