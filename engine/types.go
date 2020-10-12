package engine

// 定义请求
type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}

// 定义解析器返回结果
type ParserResult struct {
	Requests []Request
	Data     []interface{}
}
