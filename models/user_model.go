package models

import (
	"database/sql"
	"errors"

	"dana-tech.com/orm"

	"github.com/bitly/go-simplejson"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var (
	UserList map[string]*User
	db       *sql.DB
)

type User struct {
	Uid      string
	Username string
	Password string
}

func init() {
	// UserList = make(map[string]*User)
	// updateDB()
}

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "Sher:123456@/myself?charset=utf8")
	if err != nil {
		beego.Error("Open db err: ", err.Error())
	}
	return db
}

func closeDB() {
	err := db.Close()
	if err != nil {
		beego.Error("Close db err: ", err.Error())
	}
}

func AddUserInfo(js *simplejson.Json) (retBody map[string]interface{}, err error) {
	body, ok := js.CheckGet("body")
	if !ok {
		beego.Error("CheckGet body err: ")
	}
	retBody = make(map[string]interface{})
	username, err := body.Get("Username").String()
	if username == "" || err != nil {
		beego.Error("Username is empty...")
		return nil, err
	}
	if ok := checkUser(username); ok {
		return nil, errors.New("Username exist...")
	}

	retBody["Username"] = username

	password, err := body.Get("Password").String()
	if password == "" || err != nil {
		beego.Error("Password is empty...")
		return nil, err
	}
	_, err = orm.Engine.Query("insert into userinfo(username, password) values(?, ?)", username, password)
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
	username, _ := body.Get("Username").String()
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
	rows, err := orm.Engine.Query("select * from userinfo")
	if err == nil && rows != nil {
		for _, row := range rows {
			Uid := string(row["uid"])
			Username := string(row["username"])
			Password := string(row["password"])
			retBody[Uid] = &User{Uid, Username, Password}
		}
	}
	return retBody
}

func UserLogin(js *simplejson.Json) error {
	body, ok := js.CheckGet("body")
	if !ok {
		beego.Error("CheckGet body err: ")
	}
	username, _ := body.Get("Username").String()
	password, _ := body.Get("Password").String()
	if username == "" || password == "" {
		return errors.New("name or pwd is empty")
	}

	rows, err := orm.Engine.Query("select username, password from userinfo where username=?", username)

	if err == nil && rows != nil {
		for _, row := range rows {
			if username == string(row["username"]) &&
				password == string(row["password"]) {
				return nil
			}
		}
	}
	return errors.New("name or pwd error")
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}

// checkUser true: exists, false: not exists
func checkUser(username string) bool {

	rows, err := orm.Engine.Query("select * from userinfo where username=?", username)
	if err == nil && rows != nil {
		return true
	}
	return false
}
