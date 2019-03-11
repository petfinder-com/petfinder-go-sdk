package main

import (
	"fmt"
	"net/http"
	"os"
	"pf-api-sdk-go/pfapi"
)

var client *http.Client

//Main() is used as an example for accessing petfinder api
func main() {

	//Pull Client ID key and Client Secret Key from environment variables
	clientID := os.Getenv("PF_CLIENT_ID")
	clientSecret := os.Getenv("PF_CLIENT_SECRET")

	//Create pfclient Object
	pfclient, err := pfapi.NewClient(clientID, clientSecret)
	if err != nil {
		fmt.Println("Could not create client")
	}

	//Create variable based on AnimalType struct
	var types []pfapi.AnimalType

	//Retreive all animal types, put into struct
	types, _ = pfclient.GetAllTypes()

	//Iterate through animal types using struct data
	for _, t := range types {
		fmt.Println("Name: ", t.Name)
		fmt.Println("Colors: ", t.Colors)
		fmt.Println("Self Link: ", t.Links.Self.Href)
	}

	//Get a specific type
	myType, err := pfclient.GetType("dog")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myType.Name)

	//Get a particular animal by id
	myAnimal, err := pfclient.GetAnimalById("39140238")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myAnimal.ID, myAnimal.Species, myAnimal.Breeds)

	myParams := pfapi.NewPetSearchParams()
	myParams.AddParam("type", "Dog")
	myParams.AddParam("coat", "Medium")

	myAnimals, err := pfclient.GetAnimals(myParams)
	if err != nil {
		fmt.Println(err)
	}
	for _, a := range myAnimals.Animals {
		fmt.Println(a.Name)
		fmt.Println(a.Photos)
	}
	fmt.Println(myAnimals.Pagination.TotalCount)

	//Orgs
	/*myOrgs, err := pfclient.GetOrganizations()
	if err != nil {
		fmt.Println(err)
	}
	for _, a := range myOrgs.Organizations {
		fmt.Println(a.Name)
		fmt.Println(a.Photos)
	}
	fmt.Println(myOrgs.Pagination.TotalCount)*/
}
