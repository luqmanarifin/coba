package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID        int64
	Name      string
	EmailID   *int64
	CreatedAt mysql.NullTime
	UpdatedAt mysql.NullTime
}

type Email struct {
	ID        int64
	Email     string
	Owner     *User `gorm:"foreignkey:EmailID"`
	CreatedAt mysql.NullTime
	UpdatedAt mysql.NullTime
}

func getDbConnection() (db *gorm.DB, err error) {
	d, err := gorm.Open("mysql", "luqman:luqman@tcp(127.0.0.1:3306)/test?charset=utf8")
	return d, err
}

func GormAssociation() {
	db, _ := getDbConnection()
	db.AutoMigrate(&User{}, &Email{})

	email := Email{
		Email: "l.arifin.siswanto@gmail.com",
	}
	db.Create(&email)
	num := int64(1)
	user := User{
		Name:    "Luqman Arifin",
		EmailID: &num,
	}
	db.Create(&user)

	lala := Email{}
	db.Where("ID = ?", 1).First(&lala)
	if lala.Owner == (*User)(nil) {
		log.Printf("user ga ada")
	}
	// log.Printf("Email %v", lala.Owner.Name)
}
