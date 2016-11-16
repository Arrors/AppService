package controller

import (
	"AppService/model"

	"github.com/astaxie/beego"
)

// AppForceUpdateService 强制更新服务
type AppForceUpdateService struct {
	beego.Controller
}

// Get 强制更新服务
func (service *AppForceUpdateService) Get() {

	info := &model.AppInfo{
		LatestVersion:     "2.0.0",
		AppStoreURL:       "https://www.amazon.cn",
		PromptTitle:       "提示",
		PromptMessage:     "使用浏览器打开亚马逊首页",
		PromptActionTitle: "确定",
	}

	service.Data["json"] = info

	service.ServeJSON()
}
