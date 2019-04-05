package corelibs

import (
	"time"
)

// HTTPMetric is the metric we collect from target
type HTTPMetric struct {
	target					  string
	timestamp					time.Time
	nsLookup         	time.Duration
	tcpConnection    	time.Duration
	tlsHandshake     	time.Duration
	serverProcessing 	time.Duration
	contentTransfer  	time.Duration
	total							time.Duration
}