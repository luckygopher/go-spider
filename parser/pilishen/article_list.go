package pilishen

import (
	"fmt"
	"go-spider/engine"
)

func Parser(body []byte) engine.ParserResult {
	fmt.Printf("body:%s",body)
	return engine.ParserResult{}
}