package controllers

import (
	"repositories"
	"github.com/kataras/iris"
	"models"
	"strconv"
)

func AuthorsGet(ctx iris.Context) {
	result := repositories.GetAll()
	if err := result.Error; err != nil && result.RecordNotFound() {
		ctx.NotFound()
	} else {
		ctx.JSON(map[string]interface{}{
			"result":  result.Value,
			"status:": ctx.GetStatusCode(),
			"error":   result.Error,
		})
	}
}

func AuthorGetBy(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("id")
	result := repositories.GetByID(id)
	if err := result.Error; err != nil && result.RecordNotFound() {
		ctx.NotFound()
	} else {
		ctx.JSON(map[string]interface{}{
			"result":  result.Value,
			"status:": ctx.GetStatusCode(),
			"error":   result.Error,
		})
	}
}

func AuthorInsert(ctx iris.Context) {
	var author models.Author
	id, errId := strconv.ParseInt(ctx.FormValue("id"), 10, 64)
	age, errAge := strconv.Atoi(ctx.FormValue("age"))
	if errId != nil && errAge != nil {
		author = models.Author{id, ctx.FormValue("name"), ctx.FormValue("city"), age}
	}
	result := repositories.Insert(author)
	if err := result.Error; err != nil {
		panic(err)
	} else {
		ctx.JSON(map[string]interface{}{
			"status:": ctx.GetStatusCode(),
			"error":   err,
		})
	}
}
