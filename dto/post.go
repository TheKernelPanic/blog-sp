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
}

type HtmlSection struct {
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
		switch section.Mimetype {
		case model.MimeTypeHtml:
			sectionType = &HtmlSection{}
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
		switch sectionType := sectionModel.SectionType.(type) {
		case *model.HtmlSection:
			section := HtmlSection{
				Content: sectionType.Content,
				Section: Section{
					Mimetype: sectionType.Section.Mimetype,
					Sort:     sectionType.Section.Sort}}

			post.Sections = append(post.Sections, section)
		default:
			// TODO: Handle
		}
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
