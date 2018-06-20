package main

import (
	"github.com/kataras/iris"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"common"
	"web/routes"
)

func main() {
	common.Init()

	app := iris.New()
	app.Logger().SetLevel("debug")
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(map[string]interface{}{"state": false, "msg": "404 not found"})
	})

	//Author Routes
	routes.InitRoutes(app)

	app.Run(
		iris.Addr(":8080"),
		iris.WithOptimizations,
	)
}
