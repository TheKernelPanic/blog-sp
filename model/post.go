package model

const (
	MimeTypeHtml = "text/html"
)

const (
	TextType = "TEXT"
	FileType = "FILE"
)

type TypeSection interface {
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

type TextSection struct {
	ID      string  `gorm:"column:id;primary_key;type:varchar(36)"`
	Content string  `gorm:"column:content;type:text"`
	Section Section `gorm:"-"`
}

type FileSection struct {
	ID             string       `gorm:"column:id;primary_key;type:varchar(36)"`
	FileUploadedID string       `gorm:"column:file_uploaded_id;type:varchar(36)"`
	FileUploaded   FileUploaded `gorm:"references:ID"`
	Section        Section      `gorm:"-"`
}

type Section struct {
	ID            string          `gorm:"column:id;primary_key;type:varchar(36);default:uuid_generate_v4()"`
	PostID        string          `gorm:"column:post_id;type:varchar(36)"`
	Post          Post            `gorm:"references:ID"`
	Type          string          `gorm:"column:type;type:varchar(16);not null"`
	Mimetype      string          `gorm:"column:mimetype;type:varchar(255);not null"`
	Sort          int             `gorm:"column:sort;type:integer;not null"`
	SectionType   TypeSection     `gorm:"-"`
	Timestampable TimestampFields `gorm:"embedded"`
}

func (textSection *TextSection) GetSection() Section {
	return textSection.Section
}

func (textSection *TextSection) SetSection(section Section) {
	textSection.Section = section
}

func (fileSection *FileSection) GetSection() Section {
	return fileSection.Section
}

func (fileSection *FileSection) SetSection(section Section) {
	fileSection.Section = section
}
