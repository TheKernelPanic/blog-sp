package controller

import (
	"blog-sp-kernelpanic/database"
	"blog-sp-kernelpanic/model"
	"blog-sp-kernelpanic/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type mimetypesList []string

var UploadsDirectory string
var imageMimetypeAllowed mimetypesList = []string{"image/jpeg", "image/png", "image/gif"}

const (
	ImageFolder = "images"
)

func UploadImagePostController(context *fiber.Ctx) error {

	form, err := context.MultipartForm()
	if err != nil {
		err := context.SendStatus(400)
		if err != nil {
			panic(err)
		}
		return context.Send(nil)
	}

	files := form.File["image"]

	if len(files) != 1 {
		err = context.SendStatus(400)
		if err != nil {
			panic(err)
		}
		return context.Send(nil)
	}
	mimetype := files[0].Header["Content-Type"][0]
	if !imageMimetypeAllowed.checkMimetypeIsValid(mimetype) {
		err = context.SendStatus(400)
		if err != nil {
			panic(err)
		}
		return context.Send(nil)
	}

	filenameGenerated := utils.FilenameGenerator(mimetype)

	fileUploaded := model.FileUploaded{Filename: filenameGenerated, Type: "image"}
	transaction := database.DatabaseConnection.Begin()
	transaction.Create(&fileUploaded)

	path := fmt.Sprintf("%s/%s/%s", UploadsDirectory, ImageFolder, filenameGenerated)
	if err := context.SaveFile(files[0], path); err != nil {
		transaction.Rollback()
		return err
	}
	transaction.Commit()
	err = context.SendStatus(201)
	if err != nil {
		panic(err)
	}
	return context.Send(nil)
}

func (mimetypes mimetypesList) checkMimetypeIsValid(mimetype string) bool {

	for _, b := range mimetypes {
		if b == mimetype {
			return true
		}
	}
	return false
}
