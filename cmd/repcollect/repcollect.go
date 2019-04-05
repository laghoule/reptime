package main

// AwsConfig contain creds and info for AWS
type AwsConfig struct {
	AccessKey string
	SecretKey string
	Region    string
	QueueURL  string
}

// RepcollectConfig struct
type RepcollectConfig struct {
	Targets  []string
	Protocol string
	Count    int
	Interval int
	Timeout  int
	AwsConfig
}