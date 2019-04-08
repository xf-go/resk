package infra

import "fmt"

func init() {
	Register(&ConfStarter{})
}

type ConfStarter struct {
	BaseStarter
}

func (c *ConfStarter) Init(ctx StarterContext) {
	fmt.Println("配置初始化")
}

func (c *ConfStarter) Setup(ctx StarterContext) {
	fmt.Println("配置安装")
}

func (c *ConfStarter) Start(ctx StarterContext) {
	fmt.Println("配置启动")
}
