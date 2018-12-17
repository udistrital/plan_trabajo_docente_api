package main

import (
	_ "plan_trabajo_docente_api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"

	"github.com/astaxie/beego/plugins/cors"

)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://tester:tester@127.0.0.1:5432/pruebaslocal?sslmode=disable&search_path=academica")
	orm.Debug = true
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter,cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT","GET","POST", "DELETE","PATCH"},
		AllowHeaders: []string{"Origin", "x-requested-with",
				"content-type",
				"accept",
				"origin",
				"authorization",
				"x-csrftoken"},
			ExposeHeaders:    []string{"Content-Length"},
AllowCredentials: true,
	}))
	beego.Run()
}


