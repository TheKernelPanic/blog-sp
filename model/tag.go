package model

type Tag struct {
	ID            string          `gorm:"column:id;primary_key;type:varchar(36)"`
	Description   string          `gorm:"column:description;type:varchar(255)"`
	Timestampable TimestampFields `gorm:"embedded"`
}
