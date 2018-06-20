package routes

import (
	"github.com/kataras/iris"
)

func InitRoutes(app *iris.Application) {
	AuthorRoutes(app)
}
