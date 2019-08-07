package models

import (
	"database/sql"
	"errors"
	"fmt"

	"dana-tech.com/web-api/client"

	"dana-tech.com/web-api/lib"

	"dana-tech.com/orm"

	"github.com/bitly/go-simplejson"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

// ...
var (
	UserList map[string]*User
	db       *sql.DB
)

// User ...
type User struct {
	Uid         string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	TokenStatus bool   `json:"whether out of date"`
}

func init() {
	// UserList = make(map[string]*User)
	// updateDB()
}

// AddUserInfo ...
func AddUserInfo(js *simplejson.Json) (retBody map[string]interface{}, err error) {
	body, ok := js.CheckGet("body")
	if !ok {
		beego.Error("CheckGet body err: ")
	}
	retBody = make(map[string]interface{})
	username, err := body.Get("username").String()
	if username == "" || err != nil {
		beego.Error("Username is empty...")
		return nil, err
	}
	if ok := checkUser(username); ok {
		return nil, errors.New("Username exist...")
	}

	password, err := body.Get("password").String()
	if password == "" || err != nil {
		beego.Error("Password is empty...")
		return nil, err
	}
	pwdMd5 := lib.NewMD5(password)
	_, err = orm.Engine.Query("insert into userinfo(username, password) values(?, ?)", username, pwdMd5)
	if err != nil {
		beego.Error("insert db error: ", err.Error())
		return nil, err
	}
	retBody["StatusCode"] = "register success"
	retBody["code"] = "200 ok"

	return retBody, nil
}

func GetUserInfo(js *simplejson.Json) (retBody map[string]interface{}, err error) {
	body, ok := js.CheckGet("body")
	if !ok {
		beego.Error("CheckGet body err: ")
	}
	retBody = make(map[string]interface{})
	username, _ := body.Get("username").String()
	if username == "" {
		return nil, errors.New("username is tmpty")
	}
	rows, err := orm.Engine.Query("select * from userinfo where username=?", username)
	if err == nil && rows != nil {
		for _, row := range rows {
			Uid := string(row["uid"])
			Username := string(row["username"])
			Password := string(row["password"])
			token := string(row["token"])
			var tokenStatus bool
			err := lib.CheckToken(token)
			if err == nil {
				tokenStatus = true
			} else {
				tokenStatus = false
			}
			retBody[Uid] = &User{Uid, Username, Password, tokenStatus}
		}
		return retBody, nil
	}
	retBody["StatusCode"] = "this user not exists..."
	return retBody, err
}

func GetAllUserInfo() (retBody map[string]interface{}) {

	retBody = make(map[string]interface{})
	rows, err := orm.Engine.Query("select * from userinfo order by uid")
	if err == nil && rows != nil {
		for _, row := range rows {
			Uid := string(row["uid"])
			Username := string(row["username"])
			Password := string(row["password"])
			token := string(row["token"])
			var tokenStatus bool
			err := lib.CheckToken(token)
			if err == nil {
				tokenStatus = true
			} else {
				tokenStatus = false
			}
			retBody[Uid] = &User{Uid, Username, Password, tokenStatus}
		}
	}
	return retBody
}

// UserLogin 检测用户token过期方法
// 1. 使用redis hash对象(user)，将用户名(feild)与token(value)存入，当信息存在且token未过期，则用户重复登录。
// 2. 用户成功登录后，计算用户的token过期时间（单位：秒），将用户名（key）和过期时间（value）使用SET存入，并使用expire设置该键的生存时间（value），到期自动删除
func UserLogin(js *simplejson.Json) (retBody map[string]interface{}, err error) {
	body, ok := js.CheckGet("body")
	if !ok {
		beego.Error("CheckGet body err: ")
	}
	retBody = make(map[string]interface{})
	username, _ := body.Get("username").String()
	password, _ := body.Get("password").String()
	if username == "" || password == "" {
		return nil, errors.New("name or pwd is empty")
	}
	rows, err := orm.Engine.Query("select username, password, token from userinfo where username=?", username)

	pwdMD5 := lib.NewMD5(password)
	var token string
	if err == nil && rows != nil {
		for _, row := range rows {
			if username == string(row["username"]) &&
				pwdMD5 == string(row["password"]) {
				token = string(row["token"])
				if err := lib.CheckToken(token); err == nil {
					if client.GET(username) != "" {
						retBody["Status"] = "user already login..."
					} else {
						client.SET(username, token)
						retBody["Status"] = "Login success"
					}
				} else if err.Error() == "Token is expired" {
					token, _, _ = UpdateUserToken(js)
					client.SET(username, token)
					retBody["Status"] = "Refresh Token, Login success"
				}
				retBody["Code"] = "200 OK"
				return retBody, nil
			}
		}
	}
	return nil, errors.New("name or pwd error")
}

func UserLogout(username string) (retBody map[string]interface{}, err error) {
	retBody = make(map[string]interface{})
	if !client.DEL(username) {
		return retBody, nil
	}
	beego.Info("user logout...")
	return retBody, nil
}

// DeleteUser ...
func DeleteUser(js *simplejson.Json) (retBody map[string]interface{}, err error) {
	body, ok := js.CheckGet("body")
	if !ok {
		beego.Error("CheckGet body err: ")
	}
	retBody = make(map[string]interface{})
	username, _ := body.Get("username").String()
	password, _ := body.Get("password").String()
	if username == "" || password == "" {
		return nil, errors.New("name or password is empty...")
	}
	if ok := identifyInfo(username, lib.NewMD5(password)); !ok {
		retBody["error"] = "name or password error..."
		return retBody, errors.New("name or password error...")
	}
	_, err = orm.Engine.Query("delete from userinfo where username=?", username)
	if err != nil {
		beego.Error("delete sql err: ", err.Error())
	}
	retBody["Status"] = "delete success..."
	return retBody, nil
}

// UpdateUserInfo ...
func UpdateUserInfo(js *simplejson.Json) (retBody map[string]interface{}, err error) {
	body, ok := js.CheckGet("body")
	if !ok {
		beego.Error("CheckGet body err: ")
	}
	retBody = make(map[string]interface{})
	username, err := body.Get("username").String()
	if username == "" || err != nil {
		beego.Error("Username is empty...")
		return nil, err
	}
	fmt.Println("user=", username)
	if ok := checkUser(username); !ok {
		return nil, errors.New("Username not exist...")
	}

	password, err := body.Get("password").String()
	if password == "" || err != nil {
		beego.Error("Password is empty...")
		return nil, err
	}
	pwdMD5 := lib.NewMD5(password)
	result, _ := orm.Engine.Exec("update userinfo set password=? where username=?", pwdMD5, username)
	if row, _ := result.RowsAffected(); row != 1 {
		return nil, errors.New("update user info error...")
	}
	retBody["StatusCode"] = "update success..."
	return retBody, nil
}

func UpdateUserToken(js *simplejson.Json) (token string, retBody map[string]interface{}, err error) {
	body, ok := js.CheckGet("body")
	if !ok {
		beego.Error("CheckGet body error")
	}
	retBody = make(map[string]interface{})
	username, _ := body.Get("username").String()
	password, _ := body.Get("password").String()
	if username == "" || password == "" {
		return "", nil, errors.New("name or pwd is empty")
	}

	rows, err := orm.Engine.Query("select username, password from userinfo where username=?", username)
	pwdMD5 := lib.NewMD5(password)
	if err == nil && rows != nil {
		for _, row := range rows {
			if username == string(row["username"]) &&
				pwdMD5 == string(row["password"]) {
				token = lib.GenToken()
				_, err := orm.Engine.Query("update userinfo set token=? where username=? and password=?", token, username, pwdMD5)
				if err != nil {
					return token, nil, err
				}
				retBody["token"] = token
			}
		}
	}
	retBody["Status"] = "refresh token success"
	return token, retBody, nil
}

// checkUser true: exists, false: not exists
func checkUser(name string) bool {
	ok, _ := orm.Engine.Sql("select * from userinfo where username=?", name).Exist()
	return ok
}

func identifyInfo(name, pwd string) bool {
	ok, _ := orm.Engine.Sql("select * from userinfo where username=? and password=?", name, pwd).Exist()
	return ok
}
