package pfapi

import "fmt"

type AnimalType struct {
	Name    string
	Coats   []string
	Genders []string
	Colors  []string
	Links   TypeLinks `mapstructure:"_links"`
}

type TypeLinks struct {
	Self   Link
	Breeds Link
}

type Link struct {
	Href string
}

type AnimalResponse struct {
	Animals    []Animal
	Pagination Pagination
}

type Animal struct {
	ID             int
	OrganizationID string
	URL            string
	Type           string
	Species        string
	Breeds         Breeds
	Colors         Colors
	Age            string
	Gender         string
	Size           string
	Coat           string
	Name           string
	Description    string
	Photos         []Photo
	Status         string
	Attributes     Attribute
	Environment    Environment
	Tags           []string
	Contact        Contact
}

type Breeds struct {
	Primary   string
	Secondary string
	Mixed     bool
	Unknown   bool
}

type Colors struct {
	Primary   string
	Secondary string
	Tertiary  string
}

type Photo struct {
	Small  string
	Medium string
	Large  string
	Full   string
}

type Attribute struct {
	SpayedNeutered bool `mapstructure:"spayed_neutered"`
	HouseTrainied  bool `mapstructure:"house_trained"`
	Declawed       bool
	SpecialNeeds   bool `mapstructure:"special_needs"`
	ShotsCurrent   bool `mapstructure:"shots_current"`
}

type Environment struct {
	Children bool
	Dogs     bool
	cats     bool
}
type Contact struct {
	Email   string
	Phone   string
	Address Address
}

type Address struct {
	Address1 string
	Address2 string
	City     string
	State    string
	PostCode string
	Country  string
}

type Pagination struct {
	CountPerPage int `mapstructure:"count_per_page"`
	TotalCount   int `mapstructure:"total_count"`
	CurrentPage  int `mapstructure:"current_page"`
	TotalPages   int `mapstructure:"total_pages"`
}

type AnimalLinks struct {
	Self         Link
	Type         Link
	Organization Link
}

type PetSearchParams map[string]string

func (p PetSearchParams) CreateQueryString() string {
	paramString := "?"
	for paramKey, paramValue := range p {
		paramString += fmt.Sprintf("%s=%s&", paramKey, paramValue)
	}
	return paramString
}

func (p PetSearchParams) AddParam(key string, value string) {
	p[key] = p[value]
}

func NewPetSearchParams() PetSearchParams {
	return PetSearchParams{}
}
