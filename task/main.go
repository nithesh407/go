package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
)

const (
	maxBatchSize  = 10
	workerTimeout = 30
)

func main() {
	// Initialize AWS session and clients.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	sqsSvc := sqs.New(sess)
	s3Svc := s3.New(sess)
	uploader := s3manager.NewUploaderWithClient(s3Svc, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024 // 5MB part size for multipart upload
	})

	// Replace with your SQS queue URL.
	queueURL := "https://sqs.us-east-1.amazonaws.com/992925780889/orders.fifo"

	for {
		messages, err := receiveMessages(sqsSvc, queueURL)
		if err != nil {
			fmt.Println("Error receiving messages:", err)
			os.Exit(1)
		}

		if len(messages) == 0 {
			fmt.Println("No messages to process.")
			time.Sleep(time.Second * 10) // Sleep for 10 seconds before checking again.
			continue
		}

		processBatch(messages, uploader)
	}
}

func receiveMessages(sqsSvc *sqs.SQS, queueURL string) ([]*sqs.Message, error) {
	receiveParams := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: aws.Int64(maxBatchSize),
		WaitTimeSeconds:     aws.Int64(20),
	}

	receiveResult, err := sqsSvc.ReceiveMessage(receiveParams)
	if err != nil {
		return nil, err
	}

	return receiveResult.Messages, nil
}

func processBatch(messages []*sqs.Message, uploader *s3manager.Uploader) {
	// Replace with your S3 bucket name.
	bucketName := "orders1"

	// Generate the folder structure based on date, shopId, and typeofpayload.
	currentDate := time.Now().Format("2006-01-02")
	shopID := "12345"
	payloadType := "somepayload"

	for _, message := range messages {
		uuid := generateUUID()
		objectKey := fmt.Sprintf("%s/%s/%s/%s.json.gz", currentDate, shopID, payloadType, uuid)

		if err := compressAndUpload(message, uploader, bucketName, objectKey); err != nil {
			fmt.Printf("Error processing message ID %s: %v\n", *message.MessageId, err)
			continue
		}

		fmt.Printf("Processed message ID: %s, S3 object key: %s\n", *message.MessageId, objectKey)
	}
}

func compressAndUpload(message *sqs.Message, uploader *s3manager.Uploader, bucketName, objectKey string) error {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	_, err := writer.Write([]byte(*message.Body))
	if err != nil {
		return err
	}
	writer.Close()

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   bytes.NewReader(buf.Bytes()),
	})
	if err != nil {
		return err
	}

	return nil
}

func generateUUID() string {
	return uuid.New().String()
}
