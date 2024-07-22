package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"slices"

	"github.com/gorilla/mux"
)

type jsonMessage map[string]string

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
	resp, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(resp, &NewURL)
	if err != nil {
		log.Fatal(err)
	}
	if id, _ := GetStrByDomain(NewURL.Domain); id != "" {
		json.NewEncoder(w).Encode(jsonMessage{
			"status":  "406",
			"message": "Domain already shortend",
		})
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
		json.NewEncoder(w).Encode(jsonMessage{
			"status":  "200",
			"message": "http://" + os.Getenv("URL_SHORTENER_URL") + "/" + NewURL.ID,
		})
	}
}
