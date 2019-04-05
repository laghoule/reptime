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

// GetMetrics call getBodyResponse and return un array of the HTTPMetric type
// **DONT DO SLICE HERE, AND NO COUNT**
func GetMetrics(target string, count int, interval int, verbose bool) []HTTPMetric {

	var metric []HTTPMetric

	for i := 0; i < int(count); i++ {
		metric = append(metric, getBobyResponseTime(target))
		if verbose {
			printMetric(metric[i])
		}
		if count > 1 {
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}

	return metric
}

// getBobyResponseTime connect to http/https target and give response time
// Based on https://medium.com/@deeeet/trancing-http-request-latency-in-golang-65b2463f548c
func getBobyResponseTime(target string) HTTPMetric {

	// Create a new HTTP request
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create a httpstat powered context
	var result httpstat.Result
	var metric HTTPMetric
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)
	req.Close = true

	// Send request by default HTTP client
	metric.timestamp = time.Now()
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	reqEnd := time.Now()
	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		log.Fatal(err)
	}

	metric.target = target
	metric.nsLookup = result.DNSLookup
	metric.tcpConnection = result.TCPConnection
	metric.tlsHandshake = result.TLSHandshake
	metric.serverProcessing = result.ServerProcessing
	metric.contentTransfer = time.Since(reqEnd)
	metric.total = time.Since(metric.timestamp)

	res.Body.Close()
	result.End(time.Now())

	return metric
}

// printMetric output metrics to STDOUT
func printMetric(metric HTTPMetric) {
	fmt.Printf("DNS lookup: %d ms\n", int(metric.nsLookup/time.Millisecond))
	fmt.Printf("TCP connection: %d ms\n", int(metric.tcpConnection/time.Millisecond))
	fmt.Printf("TLS handshake: %d ms\n", int(metric.tlsHandshake/time.Millisecond))
	fmt.Printf("Server processing: %d ms\n", int(metric.serverProcessing/time.Millisecond))
	fmt.Printf("Content transfer: %d ms\n", int(metric.contentTransfer/time.Millisecond))
	fmt.Printf("Total: %d ms\n", int(metric.total/time.Millisecond))
}