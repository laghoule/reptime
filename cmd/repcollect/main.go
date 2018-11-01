/*
Here will be a description of the tools
*/

package main

import (
	"flag"
	"github.com/laghoule/reptime/internal/pkg/corelibs"
)

func main() {
	// cmdline args
	// countPtr := flag.Uint("count", 10, "Number of request to do")
	targetPtr := flag.String("target", "http://www.example.com", "Endpoint target")
	verbosePtr := flag.Bool("verbose", true, "Enable verbose mode")
	flag.Parse()

	// Call the target and get reponse time to stdout
	corelibs.GetBobyResponseTime(*targetPtr, *verbosePtr)
}
