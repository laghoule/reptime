package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// LoadConfigError is a error handler struct
type LoadConfigError struct {
	filename string
	section  string
	key      string
	reason   string
}

// Error handler
func (err *LoadConfigError) Error() string {
	return fmt.Sprintf("Error in configuration file %s (section: %s | key: %s)\n%s", err.filename, err.section, err.key, err.reason)
}

// validateConfig check for the integrity of the parsed configuration
func validateConfig(config *RepcollectConfig, configFile string) error {

	if len(config.AwsConfig.AccessKey) == 0 {
		return &LoadConfigError{configFile, "aws", "access_key", "Cannot be empty"}
	}

	if len(config.AwsConfig.SecretKey) == 0 {
		return &LoadConfigError{configFile, "aws", "secret_key", "Cannot be empty"}
	}

	// Todo: Add check for URL
	if len(config.AwsConfig.QueueURL) == 0 {
		return &LoadConfigError{configFile, "aws", "queue_url", "Cannot not be empty"}
	}

	for index := range config.Targets {
		if len(config.Targets[index].URL) == 0 {
			return &LoadConfigError{configFile, "target", "url", "Must not be empty"}
		}

		if config.Targets[index].Count < 1 || config.Targets[index].Count > 60 {
			return &LoadConfigError{configFile, "target", "count", "Must be between 1 and 60"}
		}

		if config.Targets[index].Interval < 1 || config.Targets[index].Interval > 60 {
			return &LoadConfigError{configFile, "target", "interval", "Must be between 1 and 60"}
		}

		if config.Targets[index].Timeout < 1 || config.Targets[index].Timeout > 30 {
			return &LoadConfigError{configFile, "target", "count", "Must be between 1 and 30"}
		}
	}

	return nil
}

// LoadConfig from configuration file
func (c *RepcollectConfig) LoadConfig(configFile string) error {

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return err
	}

	err = validateConfig(c, configFile)
	if err != nil {
		return err
	}

	return nil
}
