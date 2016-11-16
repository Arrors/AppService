package main

import (
	_ "AppService/router"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("appService", "./static/js/ios")
	beego.Run()
}
