package main

import (
	"crypto/tls"
	"fmt"
	"time"

	gomail "gopkg.in/mail.v2"
)

func main() {
	now := time.Now()
	f := make(chan string)
	go sendmail("tatvajoshi0@gmail.com", "tatvajoshi2000@gmail.com", "tatva972000", f)
	go sendmail("tatvajoshi0@gmail.com", "tatvajoshi0@gmail.com", "tatva972000", f)
	go sendmail("tatvajoshi0@gmail.com", "dhruvil.p@zymr.com", "tatva972000", f)

	a := <-f
	fmt.Println(a)
	fmt.Println("elapsed", time.Since(now))

}

func sendmail(fro, too, pass string, c chan string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", fro)

	// Set E-Mail receivers
	m.SetHeader("To", too)

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "This is Gomail test body")

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, fro, pass)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	fmt.Println("Authenticaation Started")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
	r := "Email Sent Successfully!!!"
	c <- r
}
