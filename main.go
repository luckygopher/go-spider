package main

import (
	"fmt"
	"go-spider/engine"
	"go-spider/parser/pilishen"
)

func main() {
	// 初始化引擎
	e := engine.NewCurrentEngine()
	var seeds []engine.Request
	for i := 1; i < 9; i++ {
		seeds = append(seeds, engine.Request{
			Url:        fmt.Sprintf("https://www.pilishen.com/posts?posts=%d", i),
			ParserFunc: pilishen.ListParser,
		})
	}
	// 启动
	e.Run(seeds...)
}
