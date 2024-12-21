package models

import "time"

// Coupon represents discount details
type Coupon struct {
	ID          uint      `gorm:"primaryKey"`
	Code        string    `gorm:"type:varchar(50);unique;not null"`               // Mã giảm giá
	Description string    `gorm:"type:text"`                                      // Mô tả
	Discount    float64   `gorm:"not null"`                                       // Giảm giá (số tiền hoặc %)
	ExpiryDate  time.Time `gorm:"not null"`                                       // Ngày hết hạn
	UserID      *uint     `gorm:"null"`                                           // Người tạo mã (nullable)
	User        *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Khóa ngoại
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
