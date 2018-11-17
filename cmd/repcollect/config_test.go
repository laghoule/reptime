package main

import (
	"reflect"
	"testing"
)

// TestLoadConfig verify the functionality of the LoadConfig() func
func TestLoadConfig(t *testing.T) {

	configTemplate := Config{
		Targets: []string{"www.example.com"},
		Protocol: "https", 
		Count: 5, 
		Timeout: 5,
		AwsConfig: AwsConfig {
			AccessKey: "mykey", 
			SecretKey: "mysecretkey", 
			Region: "myregion", 
			QueueURL: "myqueueurl",
		},
	}

	// Need to remove hardcoded config file
	config := LoadConfig("/etc/reptime/repcollect.conf")

	// https://golang.org/pkg/reflect/#DeepEqual
	if reflect.DeepEqual(config, configTemplate) == false {
		t.Errorf("Configuration is incorrect, got: %v, want: %v.", config, configTemplate)
	}
}
