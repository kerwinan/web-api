package redis

import (
	"dana-tech.com/web-api/global"
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)

var LoginRedis *redis.Client

func NewRedisDao() {
	host, _ := global.Cfg.GetString("redis", "host")
	port, _ := global.Cfg.GetString("redis", "port")
	LoginRedis = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: "",
	})
	_, err := LoginRedis.Ping().Result()
	if err != nil {
		beego.Error("New reids client error: ", err.Error())
	}
}
