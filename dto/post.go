package dto

import (
	"blog-sp-kernelpanic/model"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SectionType interface {
}

type Section struct {
	Mimetype string `json:"mimetype"`
	Sort     int    `json:"sort"`
	Type     string `json:"type"`
}

type FileSection struct {
	Section
	FileUploadedID string       `json:"file_uploaded_id,omitempty"`
	FileUploaded   FileUploaded `json:"file_uploaded"`
}

type TextSection struct {
	Section
	Content string `json:"content"`
}

type Post struct {
	Description string            `json:"description"`
	Slug        string            `json:"slug"`
	Sections    []SectionType     `json:"-"`
	RawSections []json.RawMessage `json:"sections,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
}

func (p *Post) UnmarshalJSON(b []byte) error {

	type postDto Post

	err := json.Unmarshal(b, (*postDto)(p))
	if err != nil {
		panic(err)
	}
	for _, raw := range p.RawSections {
		var section Section
		err = json.Unmarshal(raw, &section)
		if err != nil {
			panic(err)
		}

		var sectionType SectionType
		switch section.Type {
		case model.TextType:
			sectionType = &TextSection{}
			break
		case model.FileType:
			sectionType = &FileSection{}
			break
		default:
			return errors.New(fmt.Sprintf("Unsupported mimetype %s", section.Mimetype))
		}
		err = json.Unmarshal(raw, sectionType)
		if err != nil {
			panic(err)
		}
		p.Sections = append(p.Sections, sectionType)
	}
	return nil
}

func (p *Post) MarshalJSON() ([]byte, error) {

	type post Post

	if p.Sections != nil {
		for _, section := range p.Sections {
			encoded, err := json.Marshal(section)
			if err != nil {
				return nil, err
			}
			p.RawSections = append(p.RawSections, encoded)
		}
	}
	return json.Marshal((*post)(p))
}

func PostModelMapper(postModel model.Post) Post {

	var post Post
	post.Slug = postModel.Slug
	post.Description = postModel.Description
	post.CreatedAt = postModel.Timestampable.CreatedAt

	var sectionModel model.Section

	if len(postModel.Sections) == 0 {
		return post
	}
	for _, sectionModel = range postModel.Sections {
		var section SectionType
		switch sectionType := sectionModel.SectionType.(type) {
		case *model.TextSection:
			section = TextSection{
				Content: sectionType.Content,
				Section: Section{
					Mimetype: sectionType.Section.Mimetype,
					Type:     sectionType.Section.Type,
					Sort:     sectionType.Section.Sort}}

			break
		case *model.FileSection:
			section = FileSection{
				FileUploaded: FileUploaded{
					ID:       sectionType.FileUploaded.ID,
					Filename: sectionType.FileUploaded.Filename,
					Type:     sectionType.FileUploaded.Type},
				Section: Section{
					Mimetype: sectionType.Section.Mimetype,
					Type:     sectionType.Section.Type,
					Sort:     sectionType.Section.Sort}}
			break
		default:
			panic("Unsupported section type")
		}
		post.Sections = append(post.Sections, section)
	}
	return post
}

func PostModelMapperList(postModelList []model.Post) []Post {
	var postList []Post

	for _, postModel := range postModelList {
		postList = append(postList, PostModelMapper(postModel))
	}
	return postList
}
