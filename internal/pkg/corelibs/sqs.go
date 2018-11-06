package corelibs

import (
	"fmt"
)

// transformToJSON will transform metric to json format
func sendMetricToSQS (metric HTTPMetric) {
	fmt.Println("will send metric to SQS queue")
}
