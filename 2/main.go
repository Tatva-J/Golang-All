package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB //Global db variable for accesing database
var e error     //error variable
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() { //connectig to the Postgresql DB
	now := time.Now()
	dsn := "host=localhost user=postgres password=tatva972000 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, e := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println("COnnection has been Established!!")
	if e != nil {
		fmt.Println(e) //if any error print error
	} else {
		fmt.Println("Connection Established")
	}
	var product1 Product
	var product2 Product
	go db.First(&product1, "code = ?", "1")
	go db.First(&product2, "code = ?", "2")

	fmt.Println(product1)
	fmt.Println(product2)
	fmt.Println("Elapsed:", time.Since(now))
	// db.AutoMigrate(&Product{}) //create tables
	// db.Create(&Product{Code: "0", Price: 100})
	// db.Create(&Product{Code: "1", Price: 200})
	// db.Create(&Product{Code: "2", Price: 300})
	// db.Create(&Product{Code: "3", Price: 400})
	// db.Create(&Product{Code: "4", Price: 100})

}
