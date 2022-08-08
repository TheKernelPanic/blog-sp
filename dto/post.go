package dto

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SectionType interface {
}

type Section struct {
	Mimetype string `json:"mimetype"`
	Sort     int    `json:"sort"`
}

type HtmlSection struct {
	*Section
	Content string `json:"content"`
}

type Post struct {
	Description string            `json:"description"`
	Slug        string            `json:"slug"`
	Sections    []SectionType     `json:"_"`
	RawSections []json.RawMessage `json:"sections"`
}

func (post *Post) UnmarshalJSON(b []byte) error {
	type postDto Post

	err := json.Unmarshal(b, (*postDto)(post))
	if err != nil {
		panic(err)
	}
	for _, raw := range post.RawSections {
		var section Section
		err = json.Unmarshal(raw, &section)
		if err != nil {
			panic(err)
		}

		var sectionType SectionType
		switch section.Mimetype {
		case "text/html":
			sectionType = &HtmlSection{}
		default:
			return errors.New(fmt.Sprintf("Unsupported mimetype %s", section.Mimetype))
		}
		err = json.Unmarshal(raw, sectionType)
		if err != nil {
			panic(err)
		}
		post.Sections = append(post.Sections, sectionType)
	}
	return nil
}
