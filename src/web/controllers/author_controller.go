package controllers

import (
	"repositories"
	"github.com/kataras/iris"
)

func AuthorsGet(ctx iris.Context) {
	result, err := repositories.GetAll()
	ctx.JSON(map[string]interface{}{
		"result": result.Value,
		"error":  err,
	})
}

func AuthorGetBy(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("id")
	result, err := repositories.GetByID(id)
	ctx.JSON(map[string]interface{}{
		"result": result.Value,
		"error":  err,
	})
}
