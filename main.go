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

	r.HandleFunc("/add", CreateShortener).
		Methods("POST")

	r.HandleFunc("/{id}", RedirectDomain)

	fmt.Println("Running server ...")
	http.ListenAndServe(ListenAddr, r)
}
