package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"strings"
)

// LoadConfig from configuration file
func LoadConfig(configFile string) Config {
	var config Config

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
	config.Timeout = cfgfile.Section("repcollect").Key("timeout").MustInt(5)

	return config
}
