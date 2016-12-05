package routers

import (
	"duiba_activity/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/duiba/autologin/*", &controllers.DuibaController{}, "Get:AutoLoginUrl")
}
