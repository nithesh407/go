package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"os"
	"sync"
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
	maxWorkers    = 5
	workerTimeout = 30
)

func main() {
	// Initialize an AWS session with your credentials.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create an SQS client.
	sqsSvc := sqs.New(sess)

	// Replace with your SQS queue URL.
	queueURL := "https://sqs.us-east-1.amazonaws.com/992925780889/orders.fifo"

	// Create an S3 client.
	s3Svc := s3.New(sess)

	// Create a reusable S3 uploader with a buffer.
	uploader := s3manager.NewUploaderWithClient(s3Svc, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024 // 5MB part size for multipart upload
	})

	for {
		// Receive messages from the SQS queue.
		receiveParams := &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueURL),
			MaxNumberOfMessages: aws.Int64(maxBatchSize),
			WaitTimeSeconds:     aws.Int64(20),
		}

		receiveResult, err := sqsSvc.ReceiveMessage(receiveParams)
		if err != nil {
			fmt.Println("Error receiving messages:", err)
			os.Exit(1)
		}

		// Check if there are any messages to process.
		if len(receiveResult.Messages) == 0 {
			fmt.Println("No messages to process.")
			time.Sleep(time.Second * 10) // Sleep for 10 seconds before checking again.
			continue
		}

		// Process the received messages in batches.
		var batches [][]*sqs.Message

		for len(receiveResult.Messages) > 0 {
			batchSize := min(maxBatchSize, len(receiveResult.Messages))
			batches = append(batches, receiveResult.Messages[:batchSize])
			receiveResult.Messages = receiveResult.Messages[batchSize:]
		}

		// Process each batch of messages concurrently using worker goroutines.
		var wg sync.WaitGroup

		for _, batch := range batches {
			wg.Add(1)

			go func(batch []*sqs.Message) {
				defer wg.Done()
				processBatch(batch, uploader)
			}(batch)
		}

		// Wait for all batches to be processed.
		wg.Wait()
	}
}

func processBatch(batch []*sqs.Message, uploader *s3manager.Uploader) {
	// Replace with your S3 bucket name.


	// Generate unique UUIDs for file names.
	// var uuids []string
	// for range batch {
	// 	uuids = append(uuids, generateUUID())
	// }

	// Process the messages and upload to S3 in parallel using worker goroutines.
	var wg sync.WaitGroup

	// Create a worker channel for processing messages.
	workerCh := make(chan *sqs.Message, len(batch))

	// Start five worker goroutines.
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go worker(i, workerCh, uploader, &wg)
	}

	// Send messages to workers for processing.
	for _, message := range batch {
		workerCh <- message
	}

	// Close the worker channel to signal workers to exit after processing.
	close(workerCh)

	// Wait for all workers to finish processing.
	wg.Wait()
}

func worker(workerID int, workerCh <-chan *sqs.Message, uploader *s3manager.Uploader, wg *sync.WaitGroup) {
	defer wg.Done()

	// Create a buffer for compressed data.
	var buf bytes.Buffer

	for message := range workerCh {
		// Replace with your S3 bucket name.
		bucketName := "orders1"

		// Generate the folder structure based on date, shopId, and typeofpayload.
		currentDate := time.Now().Format("2006-01-02")
		shopID := "12345"
		payloadType := "somepayload"

		// Generate a unique UUID for the file name.
		uuid := generateUUID()

		// Construct the S3 object key.
		objectKey := fmt.Sprintf("%s/%s/%s/%s.json.gz", currentDate, shopID, payloadType, uuid)

		// Create a gzip writer to compress the data.
		buf.Reset()
		writer := gzip.NewWriter(&buf)
		_, err := writer.Write([]byte(*message.Body))
		if err != nil {
			fmt.Printf("Error compressing data for message ID %s: %v\n", *message.MessageId, err)
			continue
		}
		writer.Close()

		// Upload the compressed message content to an S3 object.
		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
			Body:   bytes.NewReader(buf.Bytes()),
		})
		if err != nil {
			fmt.Printf("Error putting object in S3 for message ID %s: %v\n", *message.MessageId, err)
			continue
		}

		fmt.Printf("Worker %d processed message ID: %s, S3 object key: %s\n", workerID, *message.MessageId, objectKey)
	}
}

// Function to generate a unique UUID
func generateUUID() string {
	return uuid.New().String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
