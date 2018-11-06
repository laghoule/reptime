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

// httpstatConvert convert httpstat to HTTPMetric
func httpstatConvert(result httpstat.Result) HTTPMetric {
	var metric HTTPMetric

	metric.nsLookup = result.DNSLookup
	metric.tcpConnection = result.TCPConnection
	metric.tlsHandshake = result.TLSHandshake
	metric.serverProcessing = result.ServerProcessing
	metric.contentTransfer = result.ContentTransfer(time.Now())
	metric.totalTime = result.Total(time.Now())

	return metric
}

// GetMetrics call getBodyResponse and call getMeanTimes for getting average
// response time
func GetMetrics(target string, count uint, verbose bool) HTTPMetric {

	// metric is a function closure
	metric := getMeanTimes()
	var meanMetric HTTPMetric

	for i := 0; i < int(count); i++ {
		meanMetric = metric(httpstatConvert(getBobyResponseTime(target, verbose)))
	}

	if verbose {
		fmt.Println("Mean time:")
		printMetric(meanMetric)
	}

	return meanMetric
}

// getMeanTimes collect metric and return average response times (function closure)
func getMeanTimes() func(HTTPMetric) HTTPMetric {

	var countMetric HTTPMetric
	var meanMetric HTTPMetric
	var itemCount int

	return func(metric HTTPMetric) HTTPMetric {
		itemCount++

		countMetric.nsLookup += metric.nsLookup
		countMetric.tcpConnection += metric.tcpConnection
		countMetric.tlsHandshake += metric.tlsHandshake
		countMetric.serverProcessing += metric.serverProcessing
		countMetric.contentTransfer += metric.contentTransfer
		countMetric.totalTime += metric.totalTime

		meanMetric.nsLookup = time.Duration(int(countMetric.nsLookup) / itemCount)
		meanMetric.tcpConnection = time.Duration(int(countMetric.tcpConnection) / itemCount)
		meanMetric.tlsHandshake = time.Duration(int(countMetric.tlsHandshake) / itemCount)
		meanMetric.serverProcessing = time.Duration(int(countMetric.serverProcessing) / itemCount)
		meanMetric.contentTransfer = time.Duration(int(countMetric.contentTransfer) / itemCount)
		meanMetric.totalTime = time.Duration(int(countMetric.totalTime) / itemCount)

		return meanMetric
	}
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

	// If verbose flag, we output to STDOUT, but we need to convert result
	if verbose {
		printMetric(httpstatConvert(result))
	}

	return result
}

// printMetric output metrics to STDOUT
func printMetric(metric HTTPMetric) {
	fmt.Printf("DNS lookup: %d ms\n", int(metric.nsLookup/time.Millisecond))
	fmt.Printf("TCP connection: %d ms\n", int(metric.tcpConnection/time.Millisecond))
	fmt.Printf("TLS handshake: %d ms\n", int(metric.tlsHandshake/time.Millisecond))
	fmt.Printf("Server processing: %d ms\n", int(metric.serverProcessing/time.Millisecond))
	fmt.Printf("Content transfer: %d ms\n", int(metric.contentTransfer/time.Millisecond))
	fmt.Printf("Total processing: %d ms\n\n", int(metric.totalTime/time.Millisecond))
}
