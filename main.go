package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Person struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	dsn := "user:password@tcp(localhost:3306)/db?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("mysql can not open:", err)
	}

	db.AutoMigrate(&Person{})

	//db.Create(&Person{
	//	Name: "jack",
	//	Age:  21,
	//})

	//db.Model(&Person{}).Where(1).Update("name", "Keny")

	db.Delete(&Person{}, 1)

	var person Person

	result := db.First(&person, 1)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println(person.Name)
}
