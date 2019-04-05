package main

import (
	"reflect"
	"testing"
)

// TestLoadConfig verify the functionality of the LoadConfig() func
func TestLoadConfig(t *testing.T) {

	configTemplate := RepcollectConfig{
		Targets:  []string{"www.example.com", "www.examples.com"},
		Protocol: "https",
		Count:    1,
		Interval: 1,
		Timeout:  5,
		AwsConfig: AwsConfig{
			AccessKey: "mykey",
			SecretKey: "mysecretkey",
			Region:    "myregion",
			QueueURL:  "myqueueurl",
		},
	}

	// Load testdata config
	config := LoadConfig("testdata/repcollect.conf")

	// https://golang.org/pkg/reflect/#DeepEqual
	if reflect.DeepEqual(config, configTemplate) == false {
		t.Errorf("Configuration is incorrect, got: %v, want: %v.", config, configTemplate)
	}
}
