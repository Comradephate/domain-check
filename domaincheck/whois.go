package domaincheck

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"regexp"
	"strings"
	"time"
)

//Domain struct with domain stuff in it
type Domain struct {
	Expiry      string   `json:"expiry"`
	Updated     string   `json:"updated"`
	Creation    string   `json:"creation"`
	NameServers []string `json:"nameservers"`
	Status      string   `json:"status"`
}

//GetWhois is the main way to interface with this library, takes in a domain name and returns a Domain struct
func GetWhois(domain string) (d Domain) {

	whois, _ := getRawWhois(domain, "whois.iana.org")

	whois = parseReferServer(whois)

	rawdata, _ := getRawWhois(domain, whois)

	d = parseWhoisData(rawdata)

	return
}

//ParseDomain checks if the domain is of a valid structure and return the tld
func ParseDomain(domain string) (tld string, err error) {
	var split []string

	split = strings.Split(domain, ".")
	if len(split) < 2 {
		err = fmt.Errorf("%s is likely not a valid domain name", domain)
		return
	}
	tld = split[len(split)-1]

	return
}

func getRawWhois(domain string, server string) (result string, err error) {
	var connection net.Conn
	var timeout time.Duration
	var buffer []byte
	// tld, err := ParseDomain(domain) we will need the tld once we start caching whois servers
	timeout = time.Second * 5

	connection, err = net.DialTimeout("tcp", net.JoinHostPort(server, "43"), timeout)

	if err != nil {
		return
	}

	defer connection.Close()

	//connection.Write([]byte(domain + "\r\n"))
	fmt.Fprintf(connection, "%s\r\n", domain)

	buffer, err = ioutil.ReadAll(connection)

	if err != nil {
		return
	}

	result = string(buffer)

	return

}

// takes raw whois data, returns the whois server if one is found. Assumes exactly one result.
func parseReferServer(whois string) string {
	return parser(regexp.MustCompile(`(?i)refer:\s+(.*?)(\s|$)`), 1, whois)[0]
}

//for parameters that we always expect exactly one response
func validateDomainParam(data []string, param string, domain string) error {
	switch {
	case len(data) == 0:
		return errors.New(domain + " has no data for " + param)
	case len(data) > 1:
		return errors.New(domain + " has too many matches for " + param)
	default:
		return nil
	}
}

var (
	nameRE        = regexp.MustCompile(`(?i)Domain Name:\s+(.*?)(\s|$)`)
	expiryRE      = regexp.MustCompile(`(?i)Expiry Date:\s+(.*?)(\s|$)`)
	creationRE    = regexp.MustCompile(`(?i)Creation Date:\s+(.*?)(\s|$)`)
	updatedRE     = regexp.MustCompile(`(?i)(Updated Date|Last updated):\s+(.*?)(\s|$)`)
	statusRE      = regexp.MustCompile(`(?i)(Domain|Registration )?Status:\s+(.*?)(\s|$)`)
	nameserversRE = regexp.MustCompile(`(?i)Name Server[s]?:\s+(.*?)(\s|$)`)
)

//Parse whois data and return a bunch of crap
func parseWhoisData(whois string) (d Domain) {

	name := parser(nameRE, 1, whois)[0]

	expiry := parser(expiryRE, 1, whois)
	if validateDomainParam(expiry, "expiry", name) == nil {
		d.Expiry = expiry[0]
	}
	creation := parser(creationRE, 1, whois)
	if validateDomainParam(creation, "creation", name) == nil {
		d.Creation = creation[0]
	} else {
		log.Printf("failed to assign creation time, value was %s\n", creation)
		log.Println(validateDomainParam(creation, "creation", name))
	}
	updated := parser(updatedRE, 2, whois)
	if validateDomainParam(updated, "updated", name) == nil {
		d.Updated = updated[0]
	}
	status := parser(statusRE, 2, whois)
	if validateDomainParam(status, "status", name) == nil {
		d.Status = status[0]
	}
	nameservers := parser(nameserversRE, 1, whois)
	if len(nameservers) != 0 {
		d.NameServers = nameservers
	}
	return

}

func parser(re *regexp.Regexp, group int, data string) (result []string) {

	found := re.FindAllStringSubmatch(data, -1)
	//log.Println(found)
	//log.Print("\r\n\r\n")

	if len(found) > 0 {
		for _, one := range found {
			if len(one) >= 2 && len(one[group]) > 0 {

				result = appendIfMissing(result, one[group])

			}
		}
	}

	return
}

//currently lowercases the data before returning it, which may or may not be desirable
func appendIfMissing(slice []string, i string) []string {

	i = strings.ToLower(i)

	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}

	return append(slice, i)

}
