package dto

import (
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
		case "text/html":
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
