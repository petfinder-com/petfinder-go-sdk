package pfapi

type AnimalType struct {
	Name    string
	Coats   []string
	Genders []string
	Colors  []string
	Links   Links `mapstructure:"_links"`
}

type Links struct {
	Self   Link
	Breeds Link
}

type Link struct {
	Href string
}
