package corelibs

import (
	"fmt"
	"os"
	"gopkg.in/ini.v1"
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
	
	config.accessKey = cfgfile.Section("aws").Key("access_key").String()
	config.secretKey = cfgfile.Section("aws").Key("secret_key").String()
	config.region = cfgfile.Section("aws").Key("region").String()
	config.queueURL = cfgfile.Section("aws").Key("queue_url").String()
	//config.target = cfgfile.Section("repcollect").Key("target").String()
	//config.protocol = cfgfile.Section("repcollect").Key("protocol").Int()
	//config.count = cfgfile.Section("repcollect").Key("count").Int()
	//config.timeout = cfgfile.Section("repcollect").Key("timeout").Int()

	return config
}