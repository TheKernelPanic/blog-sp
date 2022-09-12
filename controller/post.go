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
			sectionsModel = append(
				sectionsModel,
				&model.HtmlSection{
					Content: sectionType.Content,
					Section: model.Section{
						Mimetype: sectionType.Mimetype,
						Sort:     sectionType.Sort}})
		}
	}

	// Check existing slug
	result := database.DatabaseConnection.First(&model.Post{}, "slug = ?", postModel.Slug)
	if result.RowsAffected == 1 {
		err = context.SendStatus(409)
		if err != nil {
			panic(err)
		}
		return context.Send(nil)
	}

	transaction := database.DatabaseConnection.Begin()
	transaction.Create(&postModel)
	for _, sectionModel := range sectionsModel {

		var section model.Section

		section = sectionModel.GetSection()
		section.PostID = postModel.ID

		switch sectionType := sectionModel.(type) {
		case *model.HtmlSection:
			transaction.Create(&section)
			sectionType.ID = section.ID
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

func ListingGetController(context *fiber.Ctx) error {

	var posts []model.Post
	database.DatabaseConnection.Model(&model.Post{}).Find(&posts)

	postsDto := dto.MapList(posts)
	encoded, _ := json.Marshal(&postsDto)

	return context.Type("json", "utf-8").Send(encoded)
}

func SlugGetController(context *fiber.Ctx) error {

	var post model.Post
	result := database.DatabaseConnection.Model(&model.Post{}).Preload("Sections").First(&post, "slug = ?", context.Params("slug"))

	if result.RowsAffected == 0 {
		err := context.SendStatus(404)
		if err != nil {
			panic(err)
		}
		return context.Send(nil)
	}

	for index, section := range post.Sections {

		var sectionType model.SectionType

		switch section.Mimetype {
		case model.MimeTypeHtml:
			var htmlSectionType *model.HtmlSection
			result = database.DatabaseConnection.Model(&model.HtmlSection{}).First(&htmlSectionType, "id = ?", section.ID)
			if result.RowsAffected == 0 {
				err := context.SendStatus(404)
				if err != nil {
					panic(err)
				}
				return context.Send(nil)
			}
			sectionType = htmlSectionType
			break
		default:
			// TODO: Handle
		}
		sectionType.SetSection(section)
		post.Sections[index].SectionType = sectionType
	}
	postDto := dto.Map(post)

	encoded, _ := json.Marshal(&postDto)

	return context.Type("json", "utf-8").Send(encoded)
}
