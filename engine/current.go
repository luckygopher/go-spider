package engine

// 定义引擎
type CurrentEngine struct {
}

// 引擎构造函数
func NewCurrentEngine() *CurrentEngine {
	c := new(CurrentEngine)
	return c
}

// 引擎启动方法
func (c *CurrentEngine) Run(seeds ...Request) {
	// 声明一个请求种子队列
	var requests []Request
	// 遍历push任务到队列中
	for _, seed := range seeds {
		requests = append(requests, seed)
	}
	// for循环pop队列，直到队列为空
	for len(requests) > 0 {
		// 取出一个种子
		//seed := requests[0]
		// 重新赋值种子队列
		//requests = requests[1:]
		// TODO 调用任务处理
		// TODO 处理任务返回结果,将探索到的 seed 添加到队列中继续处理
	}
}

// 任务处理
func worker(r Request) (ParserResult, error) {
	// TODO 1、需要实现fetch爬取到种子的html内容
	// TODO 2、需要实现parser解析fetch的html内容
	// TODO 3、返回我们解析结果给引擎做后续处理
	return ParserResult{},nil
}
