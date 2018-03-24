package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Comradephate/domain-check/domaincheck"
	"github.com/gorilla/mux"
)

func main() {
	dc := domaincheck.DomainCheck{}
	r := mux.NewRouter()
	r.HandleFunc("/", dc.HomeHandler)
	r.HandleFunc("/whois", dc.WhoisHandler) //.Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger := log.New(os.Stdout, "bad: ", log.LstdFlags)
	logger.Println("Doing the thing")
	//log.SetOutput(os.Stdout)

	logger.Fatal(srv.ListenAndServe())
}
