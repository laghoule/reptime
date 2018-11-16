package main

import (
	"reflect"
	"testing"
)

// TestLoadConfig verify the functionality of the LoadConfig() func
func TestLoadConfig(t *testing.T) {

	var configTemplate Config

	// Need to simplify this via struct literal
	configTemplate.AccessKey = "mykey"
	configTemplate.SecretKey = "mysecretkey"
	configTemplate.Count = 5
	configTemplate.Timeout = 5
	configTemplate.Targets = append(configTemplate.Targets, "www.example.com")
	configTemplate.Region = "myregion"

	config := LoadConfig("/etc/reptime/repcollect.conf")

	// https://golang.org/pkg/reflect/#DeepEqual
	if reflect.DeepEqual(config, configTemplate) {
		t.Errorf("Configuration is incorrect, got: %v, want: %v.", config, configTemplate)
	}
}
