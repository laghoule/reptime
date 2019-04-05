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

	// cmdline args
	verbosePtr := flag.Bool("verbose", true, "Enable verbose mode")
	configFilePtr := flag.String("config", "/etc/reptime/repcollect.conf", "configuration file")
	flag.Parse()

	// Call the target and get response time to stdout
	config := LoadConfig(*configFilePtr)
	for _, target := range config.Targets {
		metrics = corelibs.GetMetrics(config.Protocol+"://"+target, config.Count, config.Interval, *verbosePtr)
		corelibs.SendToQueue(metrics, config.QueueURL)
	}

}
