package main

// RepcollectConfig is the yaml primary components
type RepcollectConfig struct {
	AwsConfig AwsConfig `yaml:"aws"`
	Targets   []Targets `yaml:"targets"`
}

// AwsConfig contain creds and info for AWS
type AwsConfig struct {
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	Region    string `yaml:"region"`
	QueueURL  string `yaml:"queue_url"`
}

// Targets contain informations about the URL to scrape
type Targets struct {
	URL      string `yaml:"url"`
	Count    int    `yaml:"count"`
	Interval int    `yaml:"interval"`
	Timeout  int    `yaml:"timeout"`
}
