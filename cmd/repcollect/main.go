/*
Here will be a description of the tools
*/

package main

import (
	"fmt"
	"flag"
	"github.com/laghoule/reptime/internal/pkg/corelibs"
)

func main() {
	// cmdline args
	countPtr := flag.Uint("count", 5, "Number of request to do")
	targetPtr := flag.String("target", "https://www.example.com", "Endpoint target")
	verbosePtr := flag.Bool("verbose", false, "Enable verbose mode")
	configFilePtr := flag.String("config", "/etc/reptime/repcollect.conf", "configuration file")
	flag.Parse()

	// Call the target and get response time to stdout
	fmt.Println(corelibs.LoadConfig(*configFilePtr))
	corelibs.GetMetrics(*targetPtr, *countPtr, *verbosePtr)
}
