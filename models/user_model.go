package models

import (
	"database/sql"
	"errors"
	"fmt"

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
	Uid      string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
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
			retBody[Uid] = &User{Uid, Username, Password}
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
			fmt.Println("id=", Uid)
			retBody[Uid] = &User{Uid, Username, Password}
		}
	}
	return retBody
}

func UserLogin(js *simplejson.Json) (retBody map[string]interface{}, err error) {
	body, ok := js.CheckGet("body")
	if !ok {
		beego.Error("CheckGet body err: ")
	}
	username, _ := body.Get("username").String()
	password, _ := body.Get("password").String()
	if username == "" || password == "" {
		return nil, errors.New("name or pwd is empty")
	}

	rows, err := orm.Engine.Query("select username, password from userinfo where username=?", username)

	if err == nil && rows != nil {
		for _, row := range rows {
			if username == string(row["username"]) &&
				password == string(row["password"]) {
				retBody["Status"] = "Login success"
				retBody["Code"] = "200 OK"
				return retBody, nil
			}
		}
	}
	return nil, errors.New("name or pwd error")
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

// checkUser true: exists, false: not exists
func checkUser(name string) bool {
	ok, _ := orm.Engine.Sql("select * from userinfo where username=?", name).Exist()
	return ok
}

func identifyInfo(name, pwd string) bool {
	ok, _ := orm.Engine.Sql("select * from userinfo where username=? and password=?", name, pwd).Exist()
	return ok
}
