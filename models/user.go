package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"unique" json:"email"`
	Password     string `json:"password"`
	FavoriteColor string `json:"favorite_color"` // Favorite color for password recovery
}
