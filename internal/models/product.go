package models

import "time"

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"type:varchar(100);not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"not null"`
	Stock       int     `gorm:"not null"`
	Orders      []Order `gorm:"many2many:order_products;"` // Quan hệ Many-to-Many với Order
	Carts       []Cart  `gorm:"many2many:cart_products;"`  // Quan hệ Many-to-Many với Cart
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
