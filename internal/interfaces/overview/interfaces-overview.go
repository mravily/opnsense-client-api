package interfaces

import (
	"net/http"
)

// GetInterfacesOverview creates a new HTTP GET request for the OPNsense interfaces overview endpoint.
//
// url: The base URL of the OPNsense instance.
// Returns: The constructed *http.Request and any error encountered.
func GetInterfacesOverview(url string) (*http.Request, error) {
	endpoint := url + "/api/interfaces/overview/export"
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	return req, err
}