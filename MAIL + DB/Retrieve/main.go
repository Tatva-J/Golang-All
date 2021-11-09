package main

import (
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB //Global db variable for accesing database
var e error     //error variable
type User struct {
	gorm.Model
	Name    string
	Emailid string
}

func main() { //connectig to the Postgresql DB
	now := time.Now()
	dsn := "host=localhost user=postgres password=tatva972000 dbname=CONN port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, e := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println("COnnection has been Established!!")
	if e != nil {
		fmt.Println(e) //if any error print error
	} else {
		fmt.Println("Connection Established")
	}
	var name []string
	var mailid []string
	var wg sync.WaitGroup
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		db.Raw("SELECT name FROM users").Scan(&name) //it takes 3.77ms alone
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		db.Raw("SELECT Emailid FROM users").Scan(&mailid) //it takes 3.64ms alone
	}(&wg)
	wg.Wait() //5.9ms to 6.5ms
	fmt.Println("Elapsed:", time.Since(now))
	//fmt.Println(mailid[4])
} //6.22ms

// }
// db.AutoMigrate(&User{}) //create tables
// db.Create(&User{Name: "TatvaZ", Emailid: "tatva.joshi@zymr.com"})
// db.Create(&User{Name: "Tatva2000", Emailid: "tatvajoshi2000@gmail.com"})
// db.Create(&User{Name: "Dhruvil", Emailid: "dhruvil.patel@zymr.com"})
// db.Create(&User{Name: "Neel", Emailid: "neel.kotadia@zymr.com"})
// db.Create(&User{Name: "Abdullah", Emailid: "abdullah.priyani@zymr.com"})
// for _, usr := range name {
// 	fmt.Println(usr)
// }
// for _, mail := range mailid {
// 	fmt.Println(mail)
// }
