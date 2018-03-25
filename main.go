package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Comradephate/domain-check/domaincheck"
)

func main() {
	dc := domaincheck.DomainCheck{}
	r := dc.NewRouter

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger := log.New(os.Stdout, "api: ", log.LstdFlags)

	logger.Fatal(srv.ListenAndServe())
}
