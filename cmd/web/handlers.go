package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func playersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	// GET (display text), POST (display callsign)
	if r.Method == http.MethodGet {
		ts, err := template.ParseFiles("./ui/html/index.html")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
		// store user's callsign in variable
		callsign := r.Form.Get("callsign")
		cookie := http.Cookie{
			Name:    "callsign",
			Value:   callsign,
			Expires: time.Now().AddDate(0, 0, 1),
			Path:    "/",
		}
		http.SetCookie(w, &cookie)

		// redirect to navigation after submitting callsign
		http.Redirect(w, r, "/map", 303)
	}
}

func mapHandler(w http.ResponseWriter, r *http.Request) {
	var cookie, err = r.Cookie("callsign")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error: Could not obtain callsign from cookie!", 500)
		return
	}
	callsign := cookie.Value
	log.Println("Welcome: " + callsign) // display callsign
	// w.Write([]byte(callsign))

	ts, err := template.ParseFiles("./ui/html/navigation.html")
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/players", 303)
	}
}
