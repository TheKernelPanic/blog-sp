package model

type Category struct {
	ID            string          `gorm:"column:id;primary_key;type:varchar(36)"`
	Description   string          `gorm:"column:description;type:varchar(255);not null"`
	Slug          string          `gorm:"column:slug;type:varchar(128);unique;not null"`
	Timestampable TimestampFields `gorm:"embedded"`
}
