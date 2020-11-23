package pilishen

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"go-spider/engine"
	"strings"
)

func Parser(body []byte) engine.ParserResult {
	// 初始化返回的解析结果
	var result engine.ParserResult
	// 使用 xpath 获取数据
	doc, err := htmlquery.Parse(strings.NewReader(string(body)))
	if err != nil {
		fmt.Errorf("xpath parse err: %v", err)
	}
	// 使用 xpath 匹配数据
	articleListData := htmlquery.Find(doc, `//*[@id="app"]/div/div/div[1]/div`)
	// 遍历匹配的html节点
	for _, item := range articleListData {
		// 使用 xpath 获取到 item 的html 输出查看是否获取正确
		//itemHtml := htmlquery.OutputHTML(item, true)

		// 对获取到 item 进行正则或者 xpath 匹配，获取需要的数据
		cardUrl := "http:" + htmlquery.InnerText(htmlquery.FindOne(item, `//a[@class="card"]/@href`))
		// 返回链接
		result.Requests = append(result.Requests, engine.Request{
			Url:        cardUrl,
			ParserFunc: engine.EmptyParser,
		})
	}
	// 返回解析结果
	return result
}
