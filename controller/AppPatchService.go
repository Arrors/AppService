package controller

import (
	"github.com/astaxie/beego"
)

// AppPatchService 热修复补丁版本服务
type AppPatchService struct {
	beego.Controller
}

// Get 获取热修复补丁版本
func (server *AppPatchService) Get() {
	server.Ctx.WriteString("")
}
