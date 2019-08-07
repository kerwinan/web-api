package client

import (
	"fmt"
	"time"

	"dana-tech.com/web-api/dao/redis"
	"dana-tech.com/web-api/lib"
	"github.com/astaxie/beego"
)

func HSET(key string, field string, value interface{}) {

	if err := lib.CheckToken(value.(string)); err == nil && HEXISTS(key, field) {
		beego.Warn("user already login...")
		return
	}

	err := redis.LoginRedis.HSet(key, field, value).Err()
	if err != nil {

	}
}

// SET key:username   value:token
// calc token expiration time, and set key TTL
func SET(key string, value interface{}) {

	if GET(key) != "" {
		beego.Warn("user already login...")
		return
	}
	t := lib.CalcTokenTime(value.(string)) * int64(time.Second)
	err := redis.LoginRedis.Set(key, value, time.Duration(t)).Err()
	if err != nil {
		beego.Error("Redis SET key expiration error: ", err.Error())
	}
}

func HGET(key string, field string) string {
	str, err := redis.LoginRedis.HGet(key, field).Result()
	if err != nil {

	}
	return str
}

func GET(key string) string {
	return redis.LoginRedis.Get(key).Val()
}

func HDEL(key string, field string) bool {
	if !HEXISTS(key, field) {
		beego.Error("user not login...")
		return false
	}
	err := redis.LoginRedis.HDel(key, field).Err()
	if err != nil {

	}
	return true
}

func DEL(key string) bool {
	// if ret := redis.LoginRedis.Del(key).Val(); ret != 0 {
	// return false
	// }
	ret, _ := redis.LoginRedis.Del(key).Result()
	fmt.Println(ret)
	return true
}

func HEXISTS(key string, field string) bool {
	return redis.LoginRedis.HExists(key, field).Val()
}

func HGETALL(key string) map[string]string {
	maps, err := redis.LoginRedis.HGetAll(key).Result()
	if err != nil {

	}
	return maps
}

func HLEN(key string) int64 {
	return redis.LoginRedis.HLen(key).Val()
}

func EXPIRE(key string, token string) {
	redis.LoginRedis.Expire(key, time.Duration(lib.CalcTokenTime(token)))
}
