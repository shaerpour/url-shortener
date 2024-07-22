package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	. "github.com/shaerpour/url-shortener/internal"
)

func init() {
	godotenv.Load()
	InitDB()
}

func main() {
	r := mux.NewRouter()
	ListenAddr := os.Getenv("URL_SHORTENER_LISTENADDR")
	ListenURL := os.Getenv("URL_SHORTENER_URL")

	r.HandleFunc("/add", CreateShortener).
		Host(ListenURL).
		Methods("POST")

	r.HandleFunc("/{id}", RedirectDomain).
		Host(ListenURL)

	fmt.Println("Running server ...")
	http.ListenAndServe(ListenAddr, r)
}
