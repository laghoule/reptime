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

//var httpInterface interface {
//	GetMeanTimes()
//}

// httpMetric is the metrics
type httpMetric struct {
	nsLookup         time.Duration
	tcpConnection    time.Duration
	tlsHandshake     time.Duration
	serverProcessing time.Duration
	contentTransfer  time.Duration
	totalTime        time.Duration
}

// httpstatConvert convert httpstat to httpMetrics
func httpstatConvert(result httpstat.Result) httpMetric {
	var metrics httpMetric

	metrics.nsLookup = result.DNSLookup
	metrics.tcpConnection = result.TCPConnection
	metrics.tlsHandshake = result.TLSHandshake
	metrics.serverProcessing = result.ServerProcessing
	metrics.contentTransfer = result.ContentTransfer(time.Now())
	metrics.totalTime = result.Total(time.Now())

	return metrics
}

// GetMetrics call getBodyResponse and call getMeanTimes for getting average
// response time
func GetMetrics(target string, count uint, verbose bool) {

	// Slice of the metrics, will have len of `count`
	var metric []httpMetric

	for i := 0; i < int(count); i++ {
		metric = append(metric, httpstatConvert(getBobyResponseTime(target, verbose)))
	}

	// Only print to STDOUT for now
	getMeanTimes(metric)
}

// getMeanTimes collect metrics and return average response times
func getMeanTimes(metrics []httpMetric) {

	var meanMetrics httpMetric

	for _, metric := range metrics {
		meanMetrics.nsLookup += metric.nsLookup
		meanMetrics.tcpConnection += metric.tcpConnection
		meanMetrics.tlsHandshake += metric.tlsHandshake
		meanMetrics.serverProcessing += metric.serverProcessing
		meanMetrics.contentTransfer += metric.contentTransfer
		meanMetrics.totalTime += metric.totalTime
	}

	meanMetrics.nsLookup = time.Duration(int(meanMetrics.nsLookup) / len(metrics))
	meanMetrics.tcpConnection = time.Duration(int(meanMetrics.tcpConnection) / len(metrics))
	meanMetrics.tlsHandshake = time.Duration(int(meanMetrics.tlsHandshake) / len(metrics))
	meanMetrics.serverProcessing = time.Duration(int(meanMetrics.serverProcessing) / len(metrics))
	meanMetrics.contentTransfer = time.Duration(int(meanMetrics.contentTransfer) / len(metrics))
	meanMetrics.totalTime = time.Duration(int(meanMetrics.totalTime) / len(metrics))

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

	// If verbose flag, we output to STDOUT, but we need to convert result
	if verbose {
		printMetrics(httpstatConvert(result))
	}

	return result
}

// printMetrics output metrics to STDOUT
func printMetrics(metric httpMetric) {
	fmt.Printf("DNS lookup: %d ms\n", int(metric.nsLookup/time.Millisecond))
	fmt.Printf("TCP connection: %d ms\n", int(metric.tcpConnection/time.Millisecond))
	fmt.Printf("TLS handshake: %d ms\n", int(metric.tlsHandshake/time.Millisecond))
	fmt.Printf("Server processing: %d ms\n", int(metric.serverProcessing/time.Millisecond))
	fmt.Printf("Content transfer: %d ms\n", int(metric.contentTransfer/time.Millisecond))
	fmt.Printf("Total processing: %d ms\n\n", int(metric.totalTime/time.Millisecond))
}
