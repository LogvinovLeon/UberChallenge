package main

import (
	"fmt"
	"github.com/mailgun/mailgun-go"
	"os"
)

func main() {
	domain := "uberchallenge.email"
	mg := mailgun.NewMailgun(domain, os.Getenv("MAILGUN_API_KEY"), os.Getenv("MAILGUN_PUBLIC_KEY"))
	message := mailgun.NewMessage("admin@uberchallenge.email", "Fancy subject!", "Hello from Mailgun Go!", "logvinov.leon@gmail.com")
	s1, s2, e := mg.Send(message)
	fmt.Println(s1, s2, e)
}
