package model

type SectionType interface {
	GetSection() Section
	SetSection(section Section)
}

type Post struct {
	ID          string    `gorm:"primary_key;type:varchar(36);default:uuid_generate_v4()" json:"id"`
	Sections    []Section `gorm:"foreignKey:PostID;references:ID" json:"sections"`
	Description string    `gorm:"type:varchar(255);not null" json:"description"`
	Slug        string    `gorm:"type:varchar(255);not null;index:idx_post_slug,unique" json:"slug"`
	*TimestampFields
}

type Section struct {
	ID          string `gorm:"primary_key;type:varchar(36);default:uuid_generate_v4()" json:"id"`
	SectionID   int
	SectionType string
	PostID      string `gorm:"type:varchar(36)" json:"post_id"`
	Mimetype    string `gorm:"type:varchar(255);not null" json:"mimetype"`
	Sort        int    `gorm:"type:integer;not null"`
}

type HtmlSection struct {
	ID      int
	Content string  `gorm:"type:text" json:"content"`
	Section Section `gorm:"polymorphic:Section;polymorphicValue:html"`
}

func (htmlSection *HtmlSection) GetSection() Section {
	return htmlSection.Section
}

func (htmlSection *HtmlSection) SetSection(section Section) {
	htmlSection.Section = section
}
