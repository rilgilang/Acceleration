package main

import (
	"crypto/tls"
	"fmt"
	gomail "gopkg.in/mail.v2"
)

func main() {

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "pendikia20@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", "mikhail.nery@fixedfor.com", "jurgen.alon@fixedfor.com")

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "Yeaa boiii")

	// Settings for SMTP server
	d := gomail.NewDialer("in-v3.mailjet.com", 587, "58b935df64fbdab2a0159ffdfac759e9", "61ecbbbfba8e511e291d92b5eef19157")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return

}
