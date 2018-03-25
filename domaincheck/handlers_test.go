package domaincheck

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_WhoisHandler(t *testing.T) {
	t.Skip("Not implemented")
}

func Test_HomeHandler(t *testing.T) {
	var dr Domain
	jsonStr := []byte(`{"name":"jaronrolfe.com"}`)
	req, err := http.NewRequest("POST", "/api/v1/whois", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	WhoisHandler(resp, req)
	if err != nil {
		t.Fatalf("ah shit, %s", err)
	}
	if resp.Code != http.StatusOK {
		t.Fatalf("unexpected status code: %v: %v", "200", resp.Code)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &dr)
	if dr.Creation != "2011-09-24t00:38:24z" {
		t.Fatalf("returned domain creation time was %s", dr.Creation)
	}
}
