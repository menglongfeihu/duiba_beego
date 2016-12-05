package controllers

import (
	"duiba_activity/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type DuibaController struct {
	beego.Controller
}

func (this *DuibaController) AutoLoginUrl() {
	redirect := this.GetString("dbredirect")

	var signparams = make(map[string]string)
	if redirect != "" {
		signparams["redirect"] = redirect
	}

	// 获取用户uid
	passport := this.GetString("passport")
	token := this.GetString("token")

	login := utils.CheckLogin(passport, token)

	signparams["uid"] = "not_login"
	if login {
		uid, err := utils.GetAccountIdByPassport(passport)
		if err == nil {
			signparams["uid"] = uid
		}
	} else {
		beego.BeeLogger.Error("user not login")
	}
	beego.BeeLogger.Info("uid = " + signparams["uid"])
	signparams["credits"] = "0"
	signparams["timestamp"] = strconv.FormatInt((time.Now().UnixNano() / 1e6), 10)

	autoLoginUrl := utils.BuildUrlWithSign(beego.AppConfig.String("autologinurl"), beego.AppConfig.String("appkey"), beego.AppConfig.String("appsecret"), signparams)

	beego.BeeLogger.Info("autologinurl =%s", autoLoginUrl)
	this.Redirect(autoLoginUrl, http.StatusFound)
}
