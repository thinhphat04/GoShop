package models

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:text"`
	Price       float64
	Stock       int
	Category    string `gorm:"type:varchar(100)"`
}
