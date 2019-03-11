package pfapi

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/mapstructure"
)

var typesJson = []byte(`
{
    "types": [
        {
            "name": "Rabbit",
            "coats": [
                "Short",
                "Long"
            ],
            "colors": [
                "Agouti",
                "Black",
                "Blue / Gray",
                "Brown / Chocolate",
                "Cream",
                "Lilac",
                "Orange / Red",
                "Sable",
                "Silver Marten",
                "Tan",
                "Tortoiseshell",
                "White"
            ],
            "genders": [
                "Male",
                "Female"
            ],
            "_links": {
                "self": {
                    "href": "/v2/types/rabbit"
                },
                "breeds": {
                    "href": "/v2/types/rabbit/breeds"
                }
            }
        },
        {
            "name": "Bird",
            "coats": [],
            "colors": [
                "Black",
                "Blue",
                "Brown",
                "Buff",
                "Gray",
                "Green",
                "Olive",
                "Orange",
                "Pink",
                "Purple / Violet",
                "Red",
                "Rust / Rufous",
                "Tan",
                "White",
                "Yellow"
            ],
            "genders": [
                "Male",
                "Female",
                "Unknown"
            ],
            "_links": {
                "self": {
                    "href": "/v2/types/bird"
                },
                "breeds": {
                    "href": "/v2/types/bird/breeds"
                }
            }
        }
    ]
}
`)

var singleTypeJson = []byte(`
{
    "type": {
        "name": "Dog",
        "coats": [
            "Hairless",
            "Short",
            "Medium",
            "Long",
            "Wire",
            "Curly"
        ],
        "colors": [
            "Apricot / Beige",
            "Bicolor",
            "Black",
            "Brindle",
            "Brown / Chocolate",
            "Golden",
            "Gray / Blue / Silver",
            "Harlequin",
            "Merle (Blue)",
            "Merle (Red)",
            "Red / Chestnut / Orange",
            "Sable",
            "Tricolor (Brown, Black, & White)",
            "White / Cream",
            "Yellow / Tan / Blond / Fawn"
        ],
        "genders": [
            "Male",
            "Female"
        ],
        "_links": {
            "self": {
                "href": "/v2/types/dog"
            },
            "breeds": {
                "href": "/v2/types/dog/breeds"
            }
        }
    }
}
`)

func TestAnimalTypesDecode(t *testing.T) {
	var animalTypes []AnimalType
	var message interface{}
	err := json.Unmarshal(typesJson, &message)
	if err != nil {
		t.Errorf("Could not unmarshal JSON: %v", err)
	}
	messageMap := message.(map[string]interface{})
	typesMap := messageMap["types"].([]interface{})

	err = mapstructure.Decode(typesMap, &animalTypes)
	if err != nil {
		t.Errorf("Could not mapstructure Decode: %v", err)
	}

	if len(animalTypes) != 2 {
		t.Errorf("Should have 2 animal types, got %v", len(animalTypes))
	}

	rabbit := animalTypes[0]
	if rabbit.Name != "Rabbit" {
		t.Errorf("First animal type should be rabbit, got %v", rabbit.Name)
	}

	if rabbit.Coats[0] != "Short" {
		t.Errorf("First animal type should be rabbit, got %v", rabbit.Name)
	}

	if rabbit.Coats[1] != "Long" {
		t.Errorf("First animal type should be rabbit, got %v", rabbit.Coats[1])
	}

	if rabbit.Colors[0] != "Agouti" {
		t.Errorf("Rabbit should have Agouti color first, got %v", rabbit.Colors[0])
	}

	if rabbit.Genders[0] != "Male" {
		t.Errorf("Rabbit should have Male gender first, got %v", rabbit.Genders[0])
	}

	if rabbit.Links.Self.Href != "/v2/types/rabbit" {
		t.Errorf("Rabbit self link should be /v2/types/rabbit, got: %v", rabbit.Links.Self.Href)
	}

	if rabbit.Links.Breeds.Href != "/v2/types/rabbit/breeds" {
		t.Errorf("Rabbit self link should be /v2/types/rabbit/breeds, got: %v", rabbit.Links.Self.Href)
	}

	bird := animalTypes[1]
	if bird.Name != "Bird" {
		t.Errorf("Second animal type should be Bird, got %v", bird.Name)
	}
}

func TestAnimalTypeDecode(t *testing.T) {
	var animalType AnimalType
	var message interface{}
	err := json.Unmarshal(singleTypeJson, &message)
	if err != nil {
		t.Errorf("Could not unmarshal JSON: %v", err)
	}
	messageMap := message.(map[string]interface{})
	typesMap := messageMap["type"].(interface{})

	err = mapstructure.Decode(typesMap, &animalType)

	if animalType.Name != "Dog" {
		t.Errorf("Animal type should have been Dog, got %v", animalType.Name)
	}

	if len(animalType.Colors) != 15 {
		t.Errorf("Animal type should have had 15 colors, got %v", len(animalType.Colors))
	}

	if len(animalType.Coats) != 6 {
		t.Errorf("Animal type should have had 6 coats, got %v", len(animalType.Coats))
	}

	if len(animalType.Genders) != 2 {
		t.Errorf("Animal type should have had 2 genders, got %v", len(animalType.Genders))
	}

	if animalType.Links.Self.Href != "/v2/types/dog" {
		t.Errorf("Animal type should have a self link of /v2/types/dog, got %v", animalType.Links.Self.Href)
	}
}

func TestQueryStringBuilder(t *testing.T) {
	myParams := NewPetSearchParams()
	myParams.AddParam("type", "Dog")

	expectedQueryString := "?type=Dog&"
	queryString := myParams.CreateQueryString()
	if expectedQueryString != queryString {
		t.Errorf("Expected query string %v, got %v", expectedQueryString, queryString)
	}
}
