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
