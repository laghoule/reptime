package corelibs

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// SendToQueue send the metrics information to the SQS Queue URL
func SendToQueue(metrics []HTTPMetric, queueURL string) error {
	
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	for _, metric := range metrics {
		result, err := svc.SendMessage(&sqs.SendMessageInput{
			DelaySeconds: aws.Int64(10),
			MessageAttributes: map[string]*sqs.MessageAttributeValue{
					"target": &sqs.MessageAttributeValue{
							DataType:    aws.String("String"),
							StringValue: aws.String(metric.target),
					},
					"timestamp": &sqs.MessageAttributeValue{
						DataType:    aws.String("String"),
						StringValue: aws.String(metric.timestamp.String()),
					},
					"nsLookup": &sqs.MessageAttributeValue{
							DataType:    aws.String("String"),
							StringValue: aws.String(metric.nsLookup.String()),
					},
					"tcpConnection": &sqs.MessageAttributeValue{
							DataType:    aws.String("String"),
							StringValue: aws.String(metric.tcpConnection.String()),
					},
					"tlsHandshake": &sqs.MessageAttributeValue{
						DataType:    aws.String("String"),
						StringValue: aws.String(metric.tlsHandshake.String()),
					},
					"serverProcessing": &sqs.MessageAttributeValue{
						DataType:    aws.String("String"),
						StringValue: aws.String(metric.serverProcessing.String()),
					},
					"contentTransfer": &sqs.MessageAttributeValue{
						DataType:    aws.String("String"),
						StringValue: aws.String(metric.contentTransfer.String()),
					},
					"total": &sqs.MessageAttributeValue{
						DataType:    aws.String("String"),
						StringValue: aws.String(metric.total.String()),
					},
			},
			MessageBody: aws.String("Information about URL response time."),
			QueueUrl:    &queueURL,
		})

		if err != nil {
			fmt.Println("Error", err)
			return err
		}

		fmt.Println("Success", *result.MessageId)
	}

	return nil
}