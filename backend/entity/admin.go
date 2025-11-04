package entity 

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
}