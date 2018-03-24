package domaincheck

import (
	"fmt"
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

func (dc *DomainCheck) HomeHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("ok test")
	fmt.Fprintf(w, "Hello!")
}

func (dc *DomainCheck) WhoisHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ok test")
	domain := GetWhois("bofh.wtf")

	fmt.Fprintf(w, domain.NameServers[0])
}
