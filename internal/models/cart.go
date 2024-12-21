package models

import "time"

type Cart struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    *uint     `gorm:"null"`                                           // Để `user_id` có thể nhận giá trị NULL
	User      *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Ràng buộc khóa ngoại
	Products  []Product `gorm:"many2many:cart_products;"`                       // Quan hệ nhiều-nhiều với Product
	CreatedAt time.Time
	UpdatedAt time.Time
}
