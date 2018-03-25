package domaincheck

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type DomainCheck struct {
	Configs struct {
		Postgresql struct {
			Host     string
			database string
			username string
			password string
		}
	}
}

type DomainRequest struct {
	Name string `json:"name"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("ok test")
	fmt.Fprintf(w, "Hello!")
}

func WhoisHandler(w http.ResponseWriter, r *http.Request) {
	var dr DomainRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 10240))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &dr); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	log.Println("ok test")
	domain := GetWhois(dr.Name)
	response, _ := json.Marshal(domain)

	w.Write(response)

}
