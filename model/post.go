package model

const (
	MimeTypeHtml = "text/html"
)

type SectionType interface {
	GetSection() Section
	SetSection(section Section)
}

type Post struct {
	ID            string          `gorm:"primary_key;type:varchar(36);default:uuid_generate_v4()" json:"id"`
	Sections      []Section       `gorm:"foreignKey:PostID;references:ID" json:"sections"`
	Description   string          `gorm:"type:varchar(255);not null" json:"description"`
	Slug          string          `gorm:"type:varchar(255);not null;index:idx_post_slug,unique" json:"slug"`
	Timestampable TimestampFields `gorm:"embedded"`
}

type HtmlSection struct {
	ID      string  `gorm:"primary_key;type:varchar(36)" json:"id"`
	Content string  `gorm:"type:text" json:"content"`
	Section Section `gorm:"-"`
}

type Section struct {
	ID            string          `gorm:"primary_key;type:varchar(36);default:uuid_generate_v4()" json:"id"`
	PostID        string          `gorm:"type:varchar(36)" json:"post_id"`
	Post          Post            `gorm:"references:ID"`
	Mimetype      string          `gorm:"type:varchar(255);not null" json:"mimetype"`
	Sort          int             `gorm:"type:integer;not null"`
	SectionType   SectionType     `gorm:"-"`
	Timestampable TimestampFields `gorm:"embedded"`
}

func (htmlSection *HtmlSection) GetSection() Section {
	return htmlSection.Section
}

func (htmlSection *HtmlSection) SetSection(section Section) {
	htmlSection.Section = section
}
