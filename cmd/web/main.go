package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", redirect)
	mux.HandleFunc("/players", playersHandler)
	mux.HandleFunc("/map", mapHandler)

	// Learn GO: Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Learn GO: Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	godotenv.Load()
	port := os.Getenv("PORT")
	log.Println("Starting server on port :" + port)
	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
