// Here will be a description of the libs

package corelibs

import (
	"fmt"
	"github.com/tcnksm/go-httpstat"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

// GetMetrics call getBodyResponse and call getMeanTimes for getting average
// response time
func GetMetrics(target string, count uint, verbose bool) {

	// Slice of the metrics, will have len of `count`
	var metrics []httpstat.Result

	for i := 0; i <= int(count); i++ {
		metrics = append(metrics, getBobyResponseTime(target, verbose))
	}

	// Work in progress
	// Will return mean time in httpstat.Result type
	getMeanTimes(metrics)
}

// getMeanTimes collect metrics and return average response times
// Remove the highest and lowest value of each dataset (for more acuracy?)
func getMeanTimes(metrics []httpstat.Result) {

	var (
		DNSLookup        time.Duration
		TCPConnection    time.Duration
		TLSHandshake     time.Duration
		ServerProcessing time.Duration
		//ContentTransfer time.Duration
		//Total time.Duration
	)

	for _, metric := range metrics {
		DNSLookup += metric.DNSLookup
		TCPConnection += metric.TCPConnection
		TLSHandshake += metric.TLSHandshake
		ServerProcessing += metric.ServerProcessing
		//ContentTransfer += metric.ContentTransfer(time.Now())
		//Total += metric.Total(time.Now())
	}

	fmt.Printf("Mean DNS: %d ", int(DNSLookup/time.Millisecond)/len(metrics))
	fmt.Printf("TCPConnection: %d ", int(TCPConnection/time.Millisecond)/len(metrics))
	fmt.Printf("TLSHandshake: %d ", int(TLSHandshake/time.Millisecond)/len(metrics))
	fmt.Printf("ServerProcessing: %d ", int(ServerProcessing/time.Millisecond)/len(metrics))
	//fmt.Printf("ContentTransfer: %d ", int(ContentTransfer(time.Now())/time.Millisecond)/len(metrics))
	//fmt.Printf("Total: %d", int(Total(time.Now())/time.Millisecond)/len(metrics))
}

// getBobyResponseTime connect to http/https target and give response time
// Based on https://medium.com/@deeeet/trancing-http-request-latency-in-golang-65b2463f548c
func getBobyResponseTime(target string, verbose bool) httpstat.Result {

	// Create a new HTTP request
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create a httpstat powered context
	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)
	req.Close = true

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
		fmt.Printf("Total time of the transaction: %d ms\n\n", int(result.Total(time.Now())/time.Millisecond))
	}

	return result
}
