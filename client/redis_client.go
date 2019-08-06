package client

import (
	"dana-tech.com/web-api/dao/redis"
	"github.com/astaxie/beego"
)

func SET(key string, value interface{}) {
	_, err := redis.LoginRedis.Set(key, value, 0).Result()
	if err != nil {
		beego.Error("Redis SET error: ", err.Error())
	}
}

func GET(key string) (value string) {
	value, err := redis.LoginRedis.Get(key).Result()
	if err != nil {
		beego.Error("Redis GET error: ", err.Error())
	}
	return
}
