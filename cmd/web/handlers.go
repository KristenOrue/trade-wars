package main

import (
	"html/template"
	"log"
	"net/http"
)

func players(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	// if GET method, display "enter your callsign"
	if r.Method == http.MethodGet {
		ts, err := template.ParseFiles("./ui/html/index.html")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
	}   // else display the user's callsign
	else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}

}

func navigation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	ts, err := template.ParseFiles("./ui/html/navigation.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
        http.Redirect(w, r, "/players", 303)
}
