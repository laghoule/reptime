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
type AwsConfig struct {
	AccessKey string
	SecretKey string
	Region 		string
	QueueURL 	string
}

// Config struct 
type Config struct {
	Targets 	[]string
	Protocol 	string
	Count 		int
	Timeout 	int
	AwsConfig
}