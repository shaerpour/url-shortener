package main

import (
	"fmt"
	"net/http"
    "os"

    "github.com/joho/godotenv"
	"github.com/gorilla/mux"
	. "github.com/shaerpour/url-shortener/internal"
)

func main() {
    r := mux.NewRouter()
    godotenv.Load()
	ListenAddr := os.Getenv("URL_SHORTENER_LISTENADDR")
	ListenURL := os.Getenv("URL_SHORTENER_URL")

	r.HandleFunc("/add", CreateShortener).
		Host(ListenURL).
		Queries("domain", "").
		Methods("GET")

	r.HandleFunc("/{id}", RedirectDomain).
		Host(ListenURL)

	fmt.Println("Running server ...")
	http.ListenAndServe(ListenAddr, r)
}
