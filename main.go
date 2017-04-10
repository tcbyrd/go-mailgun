package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/mailgun/mailgun-go.v1"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func mailer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}
	APIKey := "key-ee1103e4605570723f7b6287dd1391d8"
	publicAPIKey := "pubkey-9fab2901bde6e9be903a0618e533dfef"
	domain := "app015fce21413f416191c94899f21ea43c.mailgun.org"
	mg := mailgun.NewMailgun(domain, APIKey, publicAPIKey)
	email := r.FormValue("email")
	msgSubject := r.FormValue("msg_subject")
	msgText := r.FormValue("message")
	message := mailgun.NewMessage(
		email,
		msgSubject,
		msgText,
		"thomascbyrd@gmail.com")
	resp, id, err := mg.Send(message)
	if err != nil {
		log.Println(err)
		log.Printf("Message contents: \n")
		log.Printf(" - Email: %s\n", email)
		log.Printf(" - Subject: %s\n", msgSubject)
		log.Printf(" - Text: %s\n", msgText)
		fmt.Fprint(w, "Message was invalid")
		return
	}
	fmt.Fprintf(w, "ID: %s Resp: %s\n", id, resp)
}

// func tester(w http.ResponseWriter, r *http.Request) {
// 	dump, err := httputil.DumpRequest(r, true)
// 	if err != nil {
// 		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
// 		return
// 	}
// 	fmt.Fprintf(w, "%q", dump)
// 	name := r.FormValue("name")
// 	email := r.FormValue("email")
// 	msgSubject := r.FormValue("msg_subject")
// 	message := r.FormValue("message")
// 	fmt.Fprintf(w, "Name: %s", name)
// 	fmt.Fprintf(w, "Email: %s", email)
// 	fmt.Fprintf(w, "Subject: %s", msgSubject)
// 	fmt.Fprintf(w, "message: %s", message)
// }

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
