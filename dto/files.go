package dto

import "blog-sp-kernelpanic/model"

type FileUploaded struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
	Type     string `json:"type"`
}

func FileUploadedModelMapper(fileUploadedModel *model.FileUploaded) FileUploaded {

	var fileUploadedDto FileUploaded

	fileUploadedDto.ID = fileUploadedModel.ID
	fileUploadedDto.Type = fileUploadedModel.Type
	fileUploadedDto.Filename = fileUploadedModel.Filename

	return fileUploadedDto
}
