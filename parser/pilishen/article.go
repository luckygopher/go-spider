package pilishen

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"go-spider/engine"
	"strings"
)

// 文章内容解析器
func ArticleParser(body []byte) engine.ParserResult {
	// 初始化返回的解析结果
	var result engine.ParserResult
	// 使用 xpath 获取数据
	doc, err := htmlquery.Parse(strings.NewReader(string(body)))
	if err != nil {
		fmt.Errorf("xpath parse err: %v", err)
	}
	// 使用 xpath 匹配需要的数据
	title := htmlquery.InnerText(htmlquery.FindOne(doc, `//*[@id="app"]/div/div/article/div[1]/div[2]/div[1]/h1`))
	content := htmlquery.OutputHTML(htmlquery.FindOne(doc, `//*[@id="app"]/div/div/article/div[1]/div[2]/div[2]/p[1]`), true)
	result.Requests = make([]engine.Request, 0)
	result.Data = append(result.Data, map[string]interface{}{
		"title":   title,
		"content": content,
	})
	return result
}
