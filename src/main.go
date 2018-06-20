package main

import (
	"github.com/kataras/iris"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"common"
	"routes"
)

func main() {
	common.Init()
	defer common.GetDB().Close()

	app := iris.New()
	app.Logger().SetLevel("debug")
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(map[string]interface{}{"status": iris.StatusNotFound, "message": "404 not found"})
	})

	//Author Routes
	routes.InitRoutes(app)

	app.Run(
		iris.Addr(":8080"),
		iris.WithOptimizations,
	)
}
