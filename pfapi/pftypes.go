package pfapi

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
	Animals        []Animal
	PaginationData Pagination
}

type Animal struct {
	ID             string
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
	tags           []string
	contact        Contact
}

type Breeds struct {
	Primary   string
	Secondary string
	Mixed     string
	Unknown   string
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
	SpayedNeutered bool
	HouseTrainied  bool
	Declawed       bool
	SpecialNeeds   bool
	ShotsCurrent   bool
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
	CountPerPage int
	TotalCount   int
	CurrentPage  int
	TotalPages   int
}

type AnimalLinks struct {
	Self         Link
	Type         Link
	Organization Link
}
