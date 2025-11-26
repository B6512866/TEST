package entity 

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
}