package main

import (
	"encoding/json"
	"flag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
)

func main() {
	queue := flag.String("q", "", "The name of the queue")
	timeout := flag.Int64("t", 5, "How long, in seconds, that the message is hidden from others")
	waitTime := flag.Int64("w", 20, "How long in seconds, the queue waits for messages")
	flag.Parse()

	if *queue == "" {
		log.Fatal("You must supply the name of a queue (-q QUEUE)")
	}

	if *timeout < 0 {
		*timeout = 0
	}

	if *timeout > 12*60*60 {
		*timeout = 12 * 60 * 60
	}

	if *waitTime < 0 {
		*waitTime = 0
	}

	// Create a session that gets credential values from ~/.aws/credentials
	// and the default region from ~/.aws/config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	urlResult, err := getQueueURL(svc, queue)
	if err != nil {
		log.Fatalf("Got an error getting the queue URL: %v", err)
	}

	queueURL := urlResult.QueueUrl

	msgResult, err := getMessages(svc, queueURL, timeout, waitTime)
	if err != nil {
		log.Fatalf("Got an error receiving messages: %v", err)
	}

	for _, msg := range msgResult.Messages {
		filename := getUploadedFilename(msg)

		if _, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      queueURL,
			ReceiptHandle: msg.ReceiptHandle,
		}); err != nil {
			log.Fatalf("Got an error removing the message from the queue")
		}

		log.Printf("New file %s uploaded", filename)
	}

}

func getMessages(svc *sqs.SQS, queueURL *string, timeout *int64, waitTime *int64) (*sqs.ReceiveMessageOutput, error) {
	return svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            queueURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   timeout,
		WaitTimeSeconds:     waitTime,
	})
}

func getQueueURL(svc *sqs.SQS, queue *string) (*sqs.GetQueueUrlOutput, error) {
	return svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queue,
	})
}

func getUploadedFilename(msg *sqs.Message) string {
	record := struct {
		Records []struct {
			S3 struct {
				Object struct {
					Name string `json:"key"`
				} `json:"object"`
			} `json:"s3"`
		} `json:"Records"`
	}{}

	if err := json.Unmarshal([]byte(*msg.Body), &record); err != nil {
		log.Fatalf("Unable to parse message: %s", *msg.Body)
	}
	return record.Records[0].S3.Object.Name
}
