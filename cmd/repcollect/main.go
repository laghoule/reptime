/*
Here will be a description of the tools
*/

package main

import (
	"flag"
	"github.com/laghoule/reptime/internal/pkg/corelibs"
)

func main() {
	var metrics []corelibs.HTTPMetric
	var config RepcollectConfig

	// cmdline args
	verbose := flag.Bool("verbose", true, "Enable verbose mode")
	configFile := flag.String("config", "/etc/reptime/repcollect.conf", "configuration file")
	flag.Parse()

	// Call the target and get response time to stdout
	config.LoadConfig(*configFile)
	for _, target := range config.Targets {
		metrics = corelibs.GetMetrics(target.URL, target.Count, target.Interval, *verbose)
		corelibs.SendToQueue(metrics, config.AwsConfig.QueueURL)
	}

}
