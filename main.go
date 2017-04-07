package main

import (
	"fmt"
	"log"

	"gopkg.in/mailgun/mailgun-go.v1"
)

func main() {
	APIKey := "key-ee1103e4605570723f7b6287dd1391d8"
	publicAPIKey := "pubkey-9fab2901bde6e9be903a0618e533dfef"
	domain := "app015fce21413f416191c94899f21ea43c.mailgun.org"
	mg := mailgun.NewMailgun(domain, APIKey, publicAPIKey)

	message := mailgun.NewMessage(
		"sender@example.com",
		"Fancy subject!",
		"Hello from Mailgun Go!",
		"thomascbyrd@gmail.com")
	resp, id, err := mg.Send(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
