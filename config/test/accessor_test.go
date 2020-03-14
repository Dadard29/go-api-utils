package test

import (
	"fmt"
	"github.com/Dadard29/go-api-utils/config"
	"testing"
)

var configFile = "config/test/test.json"
var accessor = config.NewAccessor(configFile, true)

func TestNewAccessor(t *testing.T) {
	config.NewAccessor(configFile, true)
}

func TestGetValueFromFile(t *testing.T) {

	_, err := accessor.GetValueFromFile("salut", "ca", "va")
	if err != nil {
		t.Errorf("Failed to get config")
	}
}

func TestGetSubcategoryFromFile(t *testing.T) {

	value, err := accessor.GetSubcategoryFromFile("salut", "ca")
	if err != nil {
		t.Errorf("Failed to get config")
	}

	fmt.Printf("%v\n", value)
}

func TestGetCategoryFromFile(t *testing.T) {

	value, err := accessor.GetCategoryFromFile("salut")
	if err != nil {
		t.Errorf("Failed to get config")
	}

	fmt.Printf("%v\n", value)
}

func TestGetEnv(t *testing.T) {
	value := accessor.GetEnv("TEST_CONFIG")
	if value != "VALUE" {
		t.Errorf("Failed to get config")
	}
}
