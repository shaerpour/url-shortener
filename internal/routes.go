package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"

	"github.com/gorilla/mux"
)

func RedirectDomain(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	url, err := GetDomainByStr(id)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func CreateShortener(w http.ResponseWriter, r *http.Request) {
	var NewURL Domain
	NewURL.Domain = r.URL.Query().Get("domain")
	if id, _ := GetStrByDomain(NewURL.Domain); id != "" {
		fmt.Fprintf(w, "The domain has been created! Send another domain")
	} else {
		for {
			id = RandomString(10)
			if slices.Contains(GetAllDomains(), id) {
				continue
			} else {
				NewURL.ID = id
				break
			}
		}
		err := AddDomain(NewURL.Domain, NewURL.ID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "Shortend URL: http://%s/%s\n", os.Getenv("URL_SHORTENER_URL"), NewURL.ID)
	}
}
