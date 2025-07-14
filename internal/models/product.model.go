package models

import "time"

// Product represents a product entity in the catalog.
// @Description Product data structure
type Product struct {
	ID          uint      `json:"id" example:"1" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" example:"iPhone 14" gorm:"not null"              validate:"required,min=2,max=100"`
	Description string    `json:"description" example:"Latest Apple smartphone..."      gorm:"type:text;not null"           validate:"required"`
	Price       float64   `json:"price" example:"999.99"                                gorm:"not null"                     validate:"required,gt=0"`
	CoverImage  string    `json:"cover_image" example:"https://example.com/iphone14.jpg" gorm:"type:text;not null"          validate:"required,url"`
	CategoryID  uint      `json:"category_id" example:"2" gorm:"not null"               validate:"required"`
	Category    Category  `json:"category" gorm:"foreignKey:CategoryID"`
	BrandID     uint      `json:"brand_id" example:"1" gorm:"not null"                  validate:"required"`
	Brand       Brand     `json:"brand" gorm:"foreignKey:BrandID"`
	CreatedAt   time.Time `json:"created_at" example:"2025-07-09T15:04:05Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2025-07-09T15:04:05Z"`
}

// Brand represents a product brand or manufacturer.
// @Description Brand data structure for catalog items
type Brand struct {
	ID         uint      `json:"id" example:"1" gorm:"primaryKey;autoIncrement"`
	Name       string    `json:"name" example:"Apple"               gorm:"type:varchar(100);not null" validate:"required,min=2,max=100"`
	CoverImage string    `json:"cover_image" example:"https://example.com/apple.png" gorm:"type:text;not null" validate:"required,url"`
	CreatedAt  time.Time `json:"created_at" example:"2025-07-09T15:04:05Z"`
	UpdatedAt  time.Time `json:"updated_at" example:"2025-07-09T15:04:05Z"`
}

// Category represents a grouping for products in the catalog.
// @Description Category data structure for organizing products
type Category struct {
	ID         uint      `json:"id" example:"1" gorm:"primaryKey;autoIncrement"`
	Title      string    `json:"title" example:"Smartphones"        gorm:"type:varchar(100);not null" validate:"required,min=2,max=100"`
	CoverImage string    `json:"cover_image" example:"https://example.com/smartphones.jpg" gorm:"type:text;not null" validate:"required,url"`
	CreatedAt  time.Time `json:"created_at" example:"2025-07-09T15:04:05Z"`
	UpdatedAt  time.Time `json:"updated_at" example:"2025-07-09T15:04:05Z"`
}
