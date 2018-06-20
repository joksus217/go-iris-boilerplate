package routes

import (
	"github.com/kataras/iris"
	"web/controllers"
)

func AuthorRoutes(app *iris.Application) {
	app.Get("/author", controllers.AuthorsGet)
	app.Get("/author/{id:int}", controllers.AuthorGetBy)
}
