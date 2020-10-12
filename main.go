package main

import "go-spider/engine"

func main() {
	// 初始化引擎
	e := engine.NewCurrentEngine()
	// 启动
	e.Run(engine.Request{
		Url:        "https://www.pilishen.com/posts",
		ParserFunc: nil,
	})
}
