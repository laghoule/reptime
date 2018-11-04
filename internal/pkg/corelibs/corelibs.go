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

// HTTPMetric is the metrics
type HTTPMetric struct {
	nsLookup         time.Duration
	tcpConnection    time.Duration
	tlsHandshake     time.Duration
	serverProcessing time.Duration
	contentTransfer  time.Duration
	totalTime        time.Duration
}

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

	// Slice of the metrics, will have len of `count`
	var metric []HTTPMetric

	for i := 0; i < int(count); i++ {
		metric = append(metric, httpstatConvert(getBobyResponseTime(target, verbose)))
	}

	return getMeanTimes(metric, verbose)
}

// getMeanTimes collect metrics and return average response times
func getMeanTimes(metrics []HTTPMetric, verbose bool) HTTPMetric {

	var meanMetric HTTPMetric

	for _, metric := range metrics {
		meanMetric.nsLookup += metric.nsLookup
		meanMetric.tcpConnection += metric.tcpConnection
		meanMetric.tlsHandshake += metric.tlsHandshake
		meanMetric.serverProcessing += metric.serverProcessing
		meanMetric.contentTransfer += metric.contentTransfer
		meanMetric.totalTime += metric.totalTime
	}

	meanMetric.nsLookup = time.Duration(int(meanMetric.nsLookup) / len(metrics))
	meanMetric.tcpConnection = time.Duration(int(meanMetric.tcpConnection) / len(metrics))
	meanMetric.tlsHandshake = time.Duration(int(meanMetric.tlsHandshake) / len(metrics))
	meanMetric.serverProcessing = time.Duration(int(meanMetric.serverProcessing) / len(metrics))
	meanMetric.contentTransfer = time.Duration(int(meanMetric.contentTransfer) / len(metrics))
	meanMetric.totalTime = time.Duration(int(meanMetric.totalTime) / len(metrics))

	if verbose {
		fmt.Println("Mean time:")
		printMetric(meanMetric)
	}

	return meanMetric
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
