package models

import "time"

// Order represents an order
type Order struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null"` // Liên kết với User
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Products    []Product `gorm:"many2many:order_products;"` // Sản phẩm trong đơn hàng
	Coupons     []Coupon  `gorm:"many2many:order_coupons;"`  // Mã giảm giá được áp dụng
	TotalAmount float64   `gorm:"not null"`                  // Tổng giá trị đơn hàng
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
