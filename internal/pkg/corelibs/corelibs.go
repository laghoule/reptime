package corelibs

import (
	"time"
)

// HTTPMetric is the metric we collect from target
type HTTPMetric struct {
	nsLookup         time.Duration
	tcpConnection    time.Duration
	tlsHandshake     time.Duration
	serverProcessing time.Duration
	contentTransfer  time.Duration
	totalTime        time.Duration
}

// Creds and info for AWS connection and services
type awsConfig struct {
	accessKey string
	secretKey string
	region 		string
	queueURL 	string
}

// Config struct 
type Config struct {
	target 		[]string
	protocol 	string
	count 		uint8
	timeout 	uint8
	awsConfig
}