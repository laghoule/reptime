// Here will be a description of the libs

package corelibs

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"github.com/tcnksm/go-httpstat"
)

// GetBody connect to the http/https target, get the body response, and return
// the result as the return
func GetBody(target string) string {
	res, err := http.Get(target)
	if err != nil {
		log.Fatal(err)
	}

	// Read the body of the request, and close the connection
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

// GetBobyResponseTime connect to http/https target and give response time
// Based on https://medium.com/@deeeet/trancing-http-request-latency-in-golang-65b2463f548c
func GetBobyResponseTime(target string, verbose bool) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create a httpstat powered context
	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)

	// Send request by default HTTP client
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if verbose {
		// Show the results
		fmt.Printf("DNS lookup: %d ms\n", int(result.DNSLookup/time.Millisecond))
		fmt.Printf("TCP connection: %d ms\n", int(result.TCPConnection/time.Millisecond))
		fmt.Printf("TLS handshake: %d ms\n", int(result.TLSHandshake/time.Millisecond))
		fmt.Printf("Server processing: %d ms\n", int(result.ServerProcessing/time.Millisecond))
		fmt.Printf("Content transfer: %d ms\n", int(result.ContentTransfer(time.Now())/time.Millisecond))
		fmt.Printf("Total time of the transaction: %d ms\n", int(result.Total(time.Now())/time.Millisecond))
	}
}
