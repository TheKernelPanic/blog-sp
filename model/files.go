package model

type FileUploaded struct {
	ID            string          `gorm:"primary_key;type:varchar(36);default:uuid_generate_v4()"`
	Filename      string          `gorm:"type:varchar(255);not null;unique"`
	Type          string          `gorm:"type:varchar(16);not null"`
	Timestampable TimestampFields `gorm:"embedded"`
}
