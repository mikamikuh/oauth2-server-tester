package main

import (
	"html/template"
	"log"
	"net/http"
)

type LoginPage struct {
	URL string
}

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		p := &LoginPage{URL: "https://www.facebook.com/dialog/oauth"}
		t, err := template.ParseFiles("html/login.html")

		if err != nil {
			log.Print(err)
			return
		}

		t.Execute(w, p)
	})

	log.Print("Client started...")
	http.ListenAndServe(":14001", nil)
}
