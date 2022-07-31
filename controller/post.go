package controller

import (
	"blog-sp-kernelpanic/database"
	"blog-sp-kernelpanic/model"
	"github.com/gofiber/fiber/v2"
)

func CreatePostController(context *fiber.Ctx) error {

	var post model.Post
	post = model.Post{Description: "Sample Port", Slug: "sample-post", Sections: []model.Section{model.Section{}}}
	database.DatabaseConnection.Save(&post)

	err := context.SendStatus(201)
	if err != nil {
		panic(err)
	}
	return context.Send(nil)
}
