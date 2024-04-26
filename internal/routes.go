package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func RedirectDomain(w http.ResponseWriter, r *http.Request) {
	id, _ := mux.Vars(r)["id"]

	for _, val := range DOMAIN_LIST {
		if val.ID == id {
			http.Redirect(w, r, val.Domain, http.StatusMovedPermanently)
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("Bad request!")
}

func CreateShortener(w http.ResponseWriter, r *http.Request) {
	var NewURL Domain
	json.NewDecoder(r.Body).Decode(&NewURL)
	for {
		UUID := RandomString(10)
		for _, val := range DOMAIN_LIST {
			if val.ID == UUID {
				break
			}
		}
		NewURL.ID = UUID
		break
	}

	DOMAIN_LIST = append(DOMAIN_LIST, NewURL)
	fmt.Fprintf(w, "Shortend URL: %s/%s", os.Getenv("URL_SHORTENER_URL"), NewURL.ID)
}
