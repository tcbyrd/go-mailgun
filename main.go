package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/mailgun/mailgun-go.v1"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from http server!")
}

func mailer(w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprintf(w, "ID: %s Resp: %s\n", id, resp)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Printf("Listening on port %s\n", port)
	http.HandleFunc("/mailer", mailer)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
