package model

type Post struct {
	ID          string    `gorm:"primary_key;type:varchar(36);default:uuid_generate_v4()" json:"id"`
	Sections    []Section `gorm:"foreignKey:PostID;references:ID"`
	Description string    `gorm:"type:varchar(255);not null" json:"description"`
	Slug        string    `gorm:"type:varchar(255);not null" json:"slug"`
	*TimestampFields
}

type Section struct {
	ID           string `gorm:"primary_key;type:varchar(36);default:uuid_generate_v4()" json:"id"`
	PostID       string `gorm:"type:varchar(36)" json:"post_id"`
	Description  string `gorm:"type:varchar(255);not null" json:"description"`
	MimeTypeID   int    `json:"mimetype_id"`
	MimeTypeType string `json:"mimetype"`
}

type HtmlSection struct {
	ID      int
	Content string  `gorm:"type:text" json:"content"`
	Section Section `gorm:"polymorphic:MimeType;polymorphicValue:html"`
}
