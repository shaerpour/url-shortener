package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func RedirectDomain(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for _, val := range DOMAIN_LIST {
		if val.ID == id {
			http.Redirect(w, r, val.Domain, http.StatusMovedPermanently)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("Bad request!")
}

func CreateShortener(w http.ResponseWriter, r *http.Request) {
	var NewURL Domain
	NewURL.Domain = r.URL.Query().Get("domain")
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
	fmt.Fprintf(w, "Shortend URL: http://%s/%s\n", os.Getenv("URL_SHORTENER_URL"), NewURL.ID)
}
