package main

import (
	api "opnsense-client-api/pkg/client"
	"os"
)

// main is the entry point of the application. It initializes the Opnsense configuration
// and calls the API client to interact with the OPNsense instance.
func main() {
	if len(os.Args) < 4 {
		panic("Usage: <program> <url> <username> <password>")
	}
	opnsense := api.Opnsense{
		Url:      os.Args[1],
		Username: os.Args[2],
		Password: os.Args[3],
	}

	api.Client(opnsense)
}
