package model

const (
	MimeTypeHtml = "text/html"
)

type SectionType interface {
	GetSection() Section
	SetSection(section Section)
}

type Post struct {
	ID            string          `gorm:"column:id;primary_key;type:varchar(36);default:uuid_generate_v4()"`
	Sections      []Section       `gorm:"foreignKey:PostID;references:ID"`
	Description   string          `gorm:"column:description;type:varchar(255);not null"`
	Slug          string          `gorm:"column:slug;type:varchar(255);not null;index:idx_post_slug,unique"`
	Timestampable TimestampFields `gorm:"embedded"`
}

type HtmlSection struct {
	ID      string  `gorm:"column:id;primary_key;type:varchar(36)"`
	Content string  `gorm:"column:content;type:text"`
	Section Section `gorm:"-"`
}

type Section struct {
	ID            string          `gorm:"column:id;primary_key;type:varchar(36);default:uuid_generate_v4()"`
	PostID        string          `gorm:"column:post_id;type:varchar(36)"`
	Post          Post            `gorm:"references:ID"`
	Mimetype      string          `gorm:"column:mimetype;type:varchar(255);not null"`
	Sort          int             `gorm:"column:sort;type:integer;not null"`
	SectionType   SectionType     `gorm:"-"`
	Timestampable TimestampFields `gorm:"embedded"`
}

func (htmlSection *HtmlSection) GetSection() Section {
	return htmlSection.Section
}

func (htmlSection *HtmlSection) SetSection(section Section) {
	htmlSection.Section = section
}
