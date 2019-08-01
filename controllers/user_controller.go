package controllers

import (
	"github.com/bitly/go-simplejson"

	"dana-tech.com/web-api/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

func (u *UserController) Post() {

	httpBody := u.Ctx.Input.RequestBody
	js, err := simplejson.NewJson(httpBody)
	if err != nil {
		beego.Error("NewJson error: ", err.Error())
	}
	cmd, err := js.Get("cmd").String()
	if err != nil {
		beego.Error("Get cmd error: ", err.Error())
	}

	retResponse := make(map[string]interface{})
	switch cmd {
	case "GetUserInfo":
		retResponse = u.GetUserInfo(js)
	case "AddUserInfo":
		retResponse = u.AddUserInfo(js)
	case "GetAllUserInfo":
		retResponse = u.GetAllUserInfo()
	case "UserLogin":
		retResponse = u.UserLogin(js)
	default:
		retResponse["Code"] = "cmd error"
	}
	u.Data["json"] = map[string]interface{}{"retResponse": retResponse}
	u.ServeJSON()
}

func (u *UserController) GetAllUserInfo() (retBody map[string]interface{}) {
	retBody = models.GetAllUserInfo()
	return retBody
}

func (u *UserController) Get() {
	u.Ctx.WriteString("")
	u.ServeJSON()

}

func (u *UserController) UserLogin(js *simplejson.Json) (retBody map[string]interface{}) {
	retBody = make(map[string]interface{})
	err := models.UserLogin(js)
	if err == nil {
		retBody["StatusCode"] = "200 ok"
		return
	}
	retBody["error"] = err.Error()
	return retBody
}

func (u *UserController) GetUserInfo(js *simplejson.Json) (retBody map[string]interface{}) {
	retBody = make(map[string]interface{})
	retBody, err := models.GetUserInfo(js)
	if err != nil {
		beego.Error("GetUserInfo cmd err: ", err.Error())
	}
	return retBody
}

func (u *UserController) AddUserInfo(js *simplejson.Json) (retBody map[string]interface{}) {
	retBody = make(map[string]interface{})
	retBody, err := models.AddUserInfo(js)
	if err != nil {
		beego.Error(err)
	}
	return retBody
}
