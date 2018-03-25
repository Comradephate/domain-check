package domaincheck

import (
	"testing"
)

var fakewhoisdata = `% IANA WHOIS server
% for more information on IANA, visit http://www.iana.org
% This query returned 1 object

refer:        whois.verisign-grs.com

domain:       COM

organisation: VeriSign Global Registry Services
address:      12061 Bluemont Way
address:      Reston Virginia 20190
address:      United States

contact:      administrative
name:         Registry Customer Service
organisation: VeriSign Global Registry Services`

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
	_, err := GetRawWhois("jaronrolfe.com", "whois.iana.org")
	if err != nil {
		t.Fatalf("a bad")
	}

}

func Test_GetWhois(t *testing.T) {
	t.Parallel()
	domain := GetWhois("jaronrolfe.com")
	if domain.Creation != "2011-09-24t00:38:24z" {
		t.Fatalf("jaronrolfe.com creation date reported incorrectly: %s", domain.Creation)
	}
}

func Test_ParseReferServer(t *testing.T) {
	result := ParseReferServer(fakewhoisdata)
	if result != "whois.verisign-grs.com" {
		t.Fatalf("failed to extract whois referral")
	}
}

func Test_ValidateDomainParam(t *testing.T) {
	t.Skip("Not implemented")
}

func Test_ParseWhoisData(t *testing.T) {
	t.Skip("Not implemented")
}
