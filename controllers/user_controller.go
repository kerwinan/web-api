package controllers

import (
	"github.com/bitly/go-simplejson"

	"dana-tech.com/web-api/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
	Guid        string
	AccessToken string
	Username    string
	UserID      string
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
	case "UpdateUserInfo":
		retResponse = u.UpdateUserInfo(js)
	case "DeleteUser":
		retResponse = u.DeleteUser(js)
	case "UpdateUserToken":
		retResponse = u.UpdateUserToken(js)
	case "UserLogout":
		retResponse = u.UserLogout(js)

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
	retBody, err := models.UserLogin(js)
	if err != nil {
		beego.Error("login failed...", err.Error())
		return nil
	}
	return retBody
}

func (u *UserController) GetUserInfo(js *simplejson.Json) map[string]interface{} {
	// retBody = make(map[string]interface{})
	retBody, err := models.GetUserInfo(js)
	if err != nil {
		beego.Error("get user info error: ", err.Error())
		return nil
	}
	return retBody
}

func (u *UserController) AddUserInfo(js *simplejson.Json) map[string]interface{} {
	// retBody = make(map[string]interface{})
	retBody, err := models.AddUserInfo(js)
	if err != nil {
		beego.Error("insert user info error: ", err.Error())
		return nil
	}
	return retBody
}

func (u *UserController) UpdateUserInfo(js *simplejson.Json) map[string]interface{} {
	// retBody = make(map[string]interface{})
	retBody, err := models.UpdateUserInfo(js)
	if err != nil {
		// retBody["error"] = err.Error()
		beego.Error("update user info error: ", err.Error())
		return nil
	}
	return retBody
}

func (u *UserController) DeleteUser(js *simplejson.Json) map[string]interface{} {
	retBody, err := models.DeleteUser(js)
	if err != nil {
		beego.Error("delete user info error: ", err.Error())
		return retBody
	}
	return retBody
}

func (u *UserController) UpdateUserToken(js *simplejson.Json) map[string]interface{} {
	_, retBody, err := models.UpdateUserToken(js)
	if err != nil {
		retBody["Error"] = err.Error()
		return retBody
	}
	return retBody
}

func (u *UserController) UserLogout(js *simplejson.Json) map[string]interface{} {
	body, ok := js.CheckGet("body")
	if !ok {
		return nil
	}
	username, _ := body.Get("username").String()
	_, err := models.UserLogout(username)
	if err != nil {
		return nil
	}
	return nil
}
