package corelibs

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

	// Classic read of values, default section can be represented as empty string
	//fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
	//fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())

	config.AccessKey = cfgfile.Section("aws").Key("access_key").String()
	config.SecretKey = cfgfile.Section("aws").Key("secret_key").String()
	config.Region = cfgfile.Section("aws").Key("region").String()
	config.QueueURL = cfgfile.Section("aws").Key("queue_url").String()
	config.Targets = strings.Fields(cfgfile.Section("repcollect").Key("target").Value())
	config.Protocol = cfgfile.Section("repcollect").Key("protocol").String()
	config.Count, _ = cfgfile.Section("repcollect").Key("count").Int()
	config.Timeout, _ = cfgfile.Section("repcollect").Key("timeout").Int()

	return config
}
