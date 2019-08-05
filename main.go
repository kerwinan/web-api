package main

import (
	"dana-tech.com/orm"
	"dana-tech.com/web-api/options"
	_ "dana-tech.com/web-api/routers"
	"github.com/astaxie/beego"
)

func main() {
	cfg := options.Init()
	orm.Init(cfg)

	// oss.Init(cfg)
	// client.Init()

	beego.Run()

	// count := Add(func(b int) int {
	// 	return b % 2
	// })
	// fmt.Println(count)
}
