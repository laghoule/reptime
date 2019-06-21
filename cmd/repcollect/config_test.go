package main

import (
	"reflect"
	"testing"
)

// TestLoadConfig verify the functionality of the LoadConfig() func
func TestLoadConfig(t *testing.T) {
	var config RepcollectConfig

	configTemplate := RepcollectConfig{
		AwsConfig: AwsConfig{
			AccessKey: "myKey",
			SecretKey: "mySecretKey",
			QueueURL:  "https://myQueueUrl",
			Region:    "myRegion",
		},
		Targets: []Targets{
			{
				URL:      "https://www.example.com",
				Count:    1,
				Interval: 1,
				Timeout:  5,
			},
			{
				URL:      "https://www.examples.com",
				Count:    1,
				Interval: 1,
				Timeout:  1,
			},
		},
	}

	// Load testdata config
	err := config.LoadConfig("testdata/repcollect.yaml")
	if err != nil {
		t.Errorf("Error loading config file: %v", err)
	}

	// https://golang.org/pkg/reflect/#DeepEqual
	if reflect.DeepEqual(config, configTemplate) == false {
		t.Errorf("Configuration is incorrect, got: %v, want: %v.", config, configTemplate)
	}
}
