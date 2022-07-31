package model

type Tag struct {
	ID          string `gorm:"primary_key;type:varchar(36)" json:"id"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	*TimestampFields
}
