package domaincheck

import (
	"fmt"
	"testing"
)

func Test_ParseDomain(t *testing.T) {
	t.Parallel()
	tld1 := "com"
	domain1 := "jaronrolfe.com"
	output, err := ParseDomain(domain1)
	if err != nil {
		panic(err)
	}
	if output != tld1 {
		t.Fatalf("%s tld is %s, but parsed to %s", domain1, tld1, output)
	}
}

func Test_GetRawWhois(t *testing.T) {
	t.Parallel()
	output, err := GetRawWhois("jaronrolfe.com", "whois.iana.org")
	if err != nil {
		t.Fatalf("a bad")
	}
	fmt.Println(output)
}

func Test_GetWhois(t *testing.T) {
	t.Skip("Not implemented")
}

func Test_ParseReferServer(t *testing.T) {
	t.Skip("Not implemented")
}

func Test_ValidateDomainParam(t *testing.T) {
	t.Skip("Not implemented")
}

func Test_ParseWhoisData(t *testing.T) {
	t.Skip("Not implemented")
}
