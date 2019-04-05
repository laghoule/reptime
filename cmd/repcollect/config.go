package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"strings"
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
func (config *RepcollectConfig) validateConfig(configFile string) error {

	if len(config.AccessKey) == 0 {
		return &LoadConfigError{configFile, "aws", "access_key", "Must not be empty"}
	}

	if len(config.SecretKey) == 0 {
		return &LoadConfigError{configFile, "aws", "secret_key", "Must not be empty"}
	}

	if len(config.QueueURL) == 0 {
		return &LoadConfigError{configFile, "aws", "queue_url", "Must not be empty"}
	}

	if len(config.Targets) == 0 {
		return &LoadConfigError{configFile, "repcollect", "target", "Must not be empty"}
	}

	if config.Protocol != "http" && config.Protocol != "https" {
		return &LoadConfigError{configFile, "repcollect", "protocol", "Must be http or https"}
	}

	if config.Count < 1 || config.Count > 60 {
		return &LoadConfigError{configFile, "repcollect", "count", "Must be between 1 and 60"}
	}

	if config.Interval < 1 || config.Count > 60 {
		return &LoadConfigError{configFile, "repcollect", "interval", "Must be between 1 and 60"}
	}

	if config.Timeout < 1 || config.Count > 30 {
		return &LoadConfigError{configFile, "repcollect", "count", "Must be between 1 and 30"}
	}

	return nil
}

// LoadConfig from configuration file
func LoadConfig(configFile string) RepcollectConfig {
	var config RepcollectConfig

	cfgfile, err := ini.Load(configFile)
	if err != nil {
		fmt.Printf("Fail to read config file: %v", err)
		os.Exit(1)
	}

	config.AccessKey = cfgfile.Section("aws").Key("access_key").String()
	config.SecretKey = cfgfile.Section("aws").Key("secret_key").String()
	config.Region = cfgfile.Section("aws").Key("region").MustString("us-east-1")
	config.QueueURL = cfgfile.Section("aws").Key("queue_url").String()
	config.Targets = strings.Fields(cfgfile.Section("repcollect").Key("target").Value())
	config.Protocol = cfgfile.Section("repcollect").Key("protocol").MustString("https")
	config.Count = cfgfile.Section("repcollect").Key("count").MustInt(5)
	config.Interval = cfgfile.Section("repcollect").Key("count").MustInt(1)
	config.Timeout = cfgfile.Section("repcollect").Key("timeout").MustInt(5)

	if err := config.validateConfig(configFile); err != nil {
		log.Fatal(err)
	}

	return config
}
