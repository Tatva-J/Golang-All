package main

import (
	"fmt"
	"net/smtp"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	now1 := time.Now() //start calculating the time.

	var wg sync.WaitGroup    //means to wait for all goroutines to finish(just creating a waitgroup)
	wg.Add(2)                //adding the number of goroutines to be finished in a waitgroup
	t := make(chan []string) //creating channel slice to talk or send data between goroutines.
	//below two goroutines will work in parallel and that is why the time is reduced
	go func(c chan []string, wg *sync.WaitGroup) {
		defer wg.Done()  //closing the waitgroup and using defer to make sure that the waitgroup is closed after function's every line is executed.
		_, e := DBconn() //creating database connection
		fmt.Println(e)   //printing fetched data from db
		c <- e           //giving data e over to channel c
	}(t, &wg) //6.1ms
	go func(c chan []string, wg *sync.WaitGroup) {
		defer wg.Done() //closing the waitgroup
		e := <-t        //picking up the data from channel c in e variable
		mail(e)         //calling mail function by sending data of email ids of users data fetch from database.
	}(t, &wg) //4.85s
	wg.Wait() //waiting for all goroutines to finish.

	fmt.Println("Total Elapsed Time For all the Operations:", time.Since(now1)) //4.85s for all the operations
	//4.75 total time with wiatgroups and goroutins
}

func sendmail(fro, too, pas string, c chan string) {
	from := fro
	password := pas
	// Receiver email address.
	to := []string{
		too,
		//"dhruvil.p@zymr.com",
		//"tatva.joshi2000@mail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	q := "Email Sent Successfully!"
	c <- q
}
func mail(e []string) {
	now := time.Now()
	t := make(chan string)
	// go sendmail("tatvajoshi0@gmail.com", e[0], "tatva972000", t) //4.9 seconds alone
	// go sendmail("tatvajoshi0@gmail.com", e[1], "tatva972000", t)
	// go sendmail("tatvajoshi0@gmail.com", e[2], "tatva972000", t)
	// go sendmail("tatvajoshi0@gmail.com", e[3], "tatva972000", t)
	// go sendmail("tatvajoshi0@gmail.com", e[4], "tatva972000", t)
	for _, v := range e {
		go sendmail("tatvajoshi0@gmail.com", v, "tatva972000", t) //4.9 seconds alone
	}
	//all taking 3.97,4.42,4.39,5.07,5.24,5.48,5.83s (means 3.9-6 seconds' range) to be sent Sucessufully
	e1 := <-t
	e2 := <-t
	e3 := <-t
	e4 := <-t
	e5 := <-t

	fmt.Println("elapsed Time Foe Sending Mails to reciepients", time.Since(now))
	fmt.Println(e1, e2, e3, e4, e5)
}
func DBconn() ([]string, []string) {
	now := time.Now() //calculating the time
	dsn := "host=localhost user=postgres password=tatva972000 dbname=CONN port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, e := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println("Database Connection has been Established!!")
	if e != nil {
		fmt.Println(e) //if any error print error
	} else {
		fmt.Println("Connection Established withour Error")
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
	return name, mailid
}

//Final timings
//DBconn:=6.59ms
//mail sending:=4.72s
//Total timings:==4.72s
