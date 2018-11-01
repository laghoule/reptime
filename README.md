# reptime
[![Go Report Card](https://goreportcard.com/badge/github.com/laghoule/reptime)](https://goreportcard.com/report/github.com/laghoule/reptime)

*repcollect* agent will collect response time of web/api from different part of the world via AWS Lambda or container (VM for old-school pal) and send it to an AWS SQS queue. *reppush* will process SQS Queue and send it to an InfluxDB server. You will be abel to visualise the data in a Grafana, via the provided dashboard.
