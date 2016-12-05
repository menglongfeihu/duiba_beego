package utils

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type UserInfo struct {
	Status   int    `json:"status"`
	Uid      int64  `json:"uid"`
	NickName string `json:"nickname"`
	SmallImg string `json:"smallimg"`
	BigImg   string `json:bigimg`
}

// 根据passport 获取用户 uid
func GetAccountIdByPassport(passport string) (string, error) {

	requestParams := make(map[string]string)

	requestParams["coding"] = "utf8"
	requestParams["p"] = passport

	data, err := HttpDoGetMapParams(beego.AppConfig.String("passporturl"), requestParams)
	beego.BeeLogger.Info("response:" + data)

	var userInfo UserInfo
	err = json.Unmarshal([]byte(data), &userInfo)

	if err == nil {
		beego.BeeLogger.Info("userInfo:%s, %s, %s", userInfo.NickName, userInfo.BigImg, userInfo.SmallImg)
		if userInfo.Status == 1 {
			return strconv.FormatInt(userInfo.Uid, 10), nil
		} else {
			return "", fmt.Errorf("uid not in response : %s", string(data))
		}
	} else {
		return "", err
	}

}

// 定义解析结构
type Attachment struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
}
type CheckLoginJson struct {
	Status     int        `json:"status"`
	Message    string     `json:"message"`
	Attachment Attachment `json:"attachment"`
}

// 根据passport和token, 判断用户是否已经登录
func CheckLogin(passport string, token string) bool {
	beego.Info("check Login, passport =" + passport + ", token =" + token)
	login := false
	if passport != "" && token != "" {
		params := "passport=" + passport + "&token=" + token + "&poid=123&plat=17&partner=1&sysver=1&sver=1.0&api_key=f351515304020cad28c92f70f002261c"
		if data, err := HttpDoGetStrParams(beego.AppConfig.String("checkloginurl"), params); err == nil {
			beego.BeeLogger.Info("response:" + data)
			var checkLoginJson CheckLoginJson
			err = json.Unmarshal([]byte(data), &checkLoginJson)
			if err == nil {
				if checkLoginJson.Status == 200 && checkLoginJson.Attachment.Status == 0 {
					login = true
				}
			}
		}
	}
	return login
}
