package main

import (
	"fmt"
	"net/http"
	"os"
	"pf-api-sdk-go/pfapi"
)

var client *http.Client

func main() {
	clientID := os.Getenv("PF_CLIENT_ID")
	clientSecret := os.Getenv("PF_CLIENT_SECRET")

	pfclient, err := pfapi.NewClient(clientID, clientSecret)
	if err != nil {
		fmt.Println("Could not create client")
	}

	var types []pfapi.AnimalType
	types, _ = pfclient.GetAllTypes()
	for _, t := range types {
		fmt.Println("Name: ", t.Name)
		fmt.Println("Colors: ", t.Colors)
		fmt.Println("Self Link: ", t.Links.Self.Href)
	}
}
