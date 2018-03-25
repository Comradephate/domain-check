package domaincheck

import (
	"github.com/gorilla/mux"
)

func (dc *DomainCheck) MakeRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/api/v1/whois", WhoisHandler).Methods("POST")

	return r
}
