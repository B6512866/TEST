package config

import (
	"fmt"
	"comshop.com/ComShop/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
func DB() *gorm.DB{
	return db
}

func ConnectionDB(){
	database, err:= gorm.Open(sqlite.Open("sa.db?cache=shared"),&gorm.Config{})
	if err != nil {
		panic("failed to connection database")
	}

	fmt.Println("connected database")
	db = database
}

func SetuoDatabase(){
	db.AutoMigrate(
		&entity.Users{},
		&entity.Admin{},
	)


	hashedPassword, _ := HashPassword("123456")
	Admin := &entity.Admin{
		Email: "admin@gmail.com",
		Password:  hashedPassword,
	}
	db.FirstOrCreate(Admin,&entity.Admin{
		Email: "admin@gmail.com",
	})
}