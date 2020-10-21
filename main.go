package main

import (
	"go-spider/engine"
	"go-spider/parser/pilishen"
)

func main() {
	// 初始化引擎
	e := engine.NewCurrentEngine()
	// 启动
	e.Run(engine.Request{
		Url:        "https://www.pilishen.com/posts",
		ParserFunc: pilishen.Parser,
	})
}
