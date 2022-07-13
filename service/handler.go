package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pari-27/tinyURLProject/utils"
	"io/ioutil"
	"log"
	"net/http"
)

func CraeteTinyURL(deps Dependencies) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		log.Println(string(body))
		var urlM map[string]string
		err = json.Unmarshal(body, &urlM)
		if err != nil {
			fmt.Println("failed to parse body")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("failed to parse request"))
		}

		urlEntry := utils.GetTinyUrl(urlM["url"])
		err = deps.TinyUrlStore.Create(urlEntry)
		if err != nil {
			fmt.Println("failed to create url")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to create urlt"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	}
}
func GetTinyURL(deps Dependencies) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tinyUrl := mux.Vars(r)["url"]
		log.Println(string(tinyUrl))

		urlMap, err := deps.TinyUrlStore.Get(tinyUrl)
		if err != nil {
			fmt.Println("failed to create url")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to create urlt"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	}
}
