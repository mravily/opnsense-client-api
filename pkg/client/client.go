package api

import (
	"io"
	"fmt"
	"crypto/tls"
	"net/http"
	"opnsense-client-api/internal/interfaces/overview"
)

// Opnsense holds the connection information for the OPNsense API.
type Opnsense struct {
	Url      string
	Username string
	Password string
}

// Client initializes an HTTP client and makes a request to the OPNsense API using the provided Opnsense configuration.
func Client(o Opnsense) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}

	MakeRequest(client, o)
}

// MakeRequest creates and sends an HTTP request to the OPNsense API using the provided client and configuration.
func MakeRequest(c *http.Client, o Opnsense) {
	req, err := interfaces.GetInterfacesOverview(o.Url)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(o.Username, o.Password)
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return
	}

	fmt.Println(string(body))
}