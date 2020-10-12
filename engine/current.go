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

}
