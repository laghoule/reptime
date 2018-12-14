package corelibs

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// SQSMessage contain the information about SQS and msg
type SQSMessage struct {
	metric []HTTPMetric
}

// CreateSQSQueue create the SQS queue, if allowed in the configuration
func CreateSQSQueue() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	// Create a SQS service client.
	svc := sqs.New(sess)

	// Create the Queue
	result, err := svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String("Pgauthier_SQS_QUEUE"),
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("60"),
			"MessageRetentionPeriod": aws.String("86400"),
		},
	})
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success QueueUrl:", *result.QueueUrl)
}
