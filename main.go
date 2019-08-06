package main

import (
	"dana-tech.com/orm"
	"dana-tech.com/web-api/dao/redis"
	"dana-tech.com/web-api/global"
	"dana-tech.com/web-api/options"
	_ "dana-tech.com/web-api/routers"
	"github.com/astaxie/beego"
)

func main() {
	options.Init()
	orm.Init(global.Cfg)
	global.Init(global.Cfg)
	redis.NewRedisDao()
	// oss.Init(cfg)
	// client.Init()
	// token := lib.GenToken()
	// fmt.Println(token, len(token))
	// // time.Sleep(time.Second * 3)
	// for {
	// 	ok := lib.CheckToken(token)
	// 	fmt.Println(ok)
	// 	time.Sleep(time.Second)
	// }

	// combinedToken, _ := cfg.GetString("token", "token")
	// userNode, AccessToken, StoreWay := lib.IdentifToken(combinedToken)
	// fmt.Println(userNode, AccessToken, StoreWay)

	beego.Run()

	// count := Add(func(b int) int {
	// 	return b % 2
	// })
	// fmt.Println(count)
}
