/*
Here will be a description of the tools
*/

package main

import (
	"flag"
	"fmt"
	"github.com/laghoule/reptime/internal/pkg/corelibs"
)

func main() {
	// cmdline args
	// countPtr := flag.Uint("count", 10, "Number of request to do")
	targetPtr := flag.String("target", "www.example.com", "Endpoint target")
	flag.Parse()

	fmt.Printf("Body: %s", corelibs.GetBody(*targetPtr))
}
