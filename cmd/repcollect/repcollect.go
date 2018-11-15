package main

// AwsConfig contain creds and info for AWS
type AwsConfig struct {
	AccessKey string
	SecretKey string
	Region    string
	QueueURL  string
}

// Config struct
type Config struct {
	Targets  []string
	Protocol string
	Count    int
	Timeout  int
	AwsConfig
}
