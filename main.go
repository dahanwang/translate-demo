package main

import (
	"github.com/astaxie/beego"
	_ "translate-demo/routers"
	_ "translate-demo/redis"
)

func main() {
	// beego
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
