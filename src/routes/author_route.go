package routes

import (
	"github.com/kataras/iris"
	"controllers"
)

func AuthorRoutes(app *iris.Application) {
	app.Get("/author", controllers.AuthorsGet)
	app.Get("/author/{id:int}", controllers.AuthorGetBy)
	app.Post("/author", controllers.AuthorInsert)
}
