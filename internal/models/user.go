package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint    `gorm:"primaryKey"`
	Username  string  `gorm:"type:varchar(100);not null;unique"`
	Email     string  `gorm:"type:varchar(100);not null;unique"`
	Password  string  `gorm:"type:varchar(255);not null"`
	Role      string  `gorm:"type:varchar(50);default:'user'"`
	Orders    []Order `gorm:"foreignKey:UserID"` // Quan hệ One-to-Many
	CreatedAt time.Time
	UpdatedAt time.Time
}

// HashPassword hashes the password before saving
func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword checks if the provided password is correct
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
