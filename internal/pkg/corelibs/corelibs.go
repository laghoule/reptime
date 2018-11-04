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

	for i := 0; i < int(count); i++ {
		metrics = append(metrics, getBobyResponseTime(target, verbose))
	}

	// Work in progress
	// Will return mean time in httpstat.Result type
	getMeanTimes(metrics)
}

// getMeanTimes collect metrics and return average response times
// Remove the highest and lowest value of each dataset (for more acuracy?)
func getMeanTimes(metrics []httpstat.Result) {

	var meanMetrics	httpstat.Result

	for _, metric := range metrics {
		meanMetrics.DNSLookup += metric.DNSLookup
		meanMetrics.TCPConnection += metric.TCPConnection
		meanMetrics.TLSHandshake += metric.TLSHandshake
		meanMetrics.ServerProcessing += metric.ServerProcessing
		//meanMetrics.Total += (metric.DNSLookup + metric.TCPConnection + metric.TLSHandshake + metric.ServerProcessing)
	}
	
	meanMetrics.DNSLookup = time.Duration(int(meanMetrics.DNSLookup)/len(metrics))
	meanMetrics.TCPConnection = time.Duration(int(meanMetrics.TCPConnection)/len(metrics))
	meanMetrics.TLSHandshake = time.Duration(int(meanMetrics.TLSHandshake)/len(metrics))
	meanMetrics.ServerProcessing = time.Duration(int(meanMetrics.ServerProcessing)/len(metrics))
	
	fmt.Println("Mean time:")
	printMetrics(meanMetrics)
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
	// TODO: change result for metric
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

	// If verbose flag, we output to STDOUT
	if verbose {
		printMetrics(result)
	}

	return result
}

// printMetrics output metrics to STDOUT
func printMetrics(metric httpstat.Result) {
	total := metric.DNSLookup + metric.TCPConnection + metric.TLSHandshake + metric.ServerProcessing
	fmt.Printf("DNS lookup: %d ms\n", int(metric.DNSLookup/time.Millisecond))
	fmt.Printf("TCP connection: %d ms\n", int(metric.TCPConnection/time.Millisecond))
	fmt.Printf("TLS handshake: %d ms\n", int(metric.TLSHandshake/time.Millisecond))
	fmt.Printf("Server processing: %d ms\n", int(metric.ServerProcessing/time.Millisecond))
	fmt.Printf("Total: %d ms\n\n", total/time.Millisecond)
}
