package router

import (
	"AppService/controller"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/appService/forceUpdate", &controller.AppForceUpdateService{})
	beego.Router("/appService/getPatchVersion", &controller.AppPatchService{})
}
