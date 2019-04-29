package main

import (
	"github.com/astaxie/beego"
	_ "goServers/routers"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	
)

func main() {
	//开启调试模式，打印到控制台
	orm.Debug = true
	
	host := beego.AppConfig.String("db::mysqlurls")
    port := beego.AppConfig.String("db::mysqlport")
    dbname := beego.AppConfig.String("db::mysqldb")
    user := beego.AppConfig.String("db::mysqluser")
    pwd := beego.AppConfig.String("db::mysqlpass")

    dbcon := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	orm.RegisterDataBase("default","mysql",dbcon)
	// if beego.BConfig.RunMode == "dev" {
	// 	beego.BConfig.WebConfig.DirectoryIndex = true
	// 	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	// }
	beego.Run()
}
