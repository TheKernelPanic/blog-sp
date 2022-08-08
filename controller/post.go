package controller

import (
	"blog-sp-kernelpanic/database"
	"blog-sp-kernelpanic/dto"
	"blog-sp-kernelpanic/model"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func CreatePostController(context *fiber.Ctx) error {

	var postDto dto.Post
	var postModel model.Post
	var sectionsModel []model.SectionType

	err := json.Unmarshal(context.Body(), &postDto)
	if err != nil {
		err = context.SendStatus(400)
		if err != nil {
			panic(err)
		}
		return context.Send(nil)
	}
	postModel = model.Post{Description: postDto.Description, Slug: postDto.Slug, Sections: []model.Section{}}
	for _, sectionDto := range postDto.Sections {
		switch sectionType := sectionDto.(type) {
		case *dto.HtmlSection:
			sectionsModel = append(sectionsModel, &model.HtmlSection{Content: sectionType.Content, Section: model.Section{Mimetype: sectionType.Mimetype, Sort: sectionType.Sort}})
		}
	}

	transaction := database.DatabaseConnection.Begin()
	transaction.Create(&postModel)
	for _, sectionModel := range sectionsModel {

		var section model.Section

		section = sectionModel.GetSection()
		section.PostID = postModel.ID
		sectionModel.SetSection(section)
		switch sectionType := sectionModel.(type) {
		case *model.HtmlSection:
			sectionType.SetSection(section)
			transaction.Create(&sectionType)
		default:
			transaction.Rollback()
			err = context.SendStatus(400)
			if err != nil {
				panic(err)
			}
			return context.Send(nil)
		}
	}
	transaction.Commit()

	err = context.SendStatus(201)
	if err != nil {
		panic(err)
	}
	return context.Send(nil)
}
