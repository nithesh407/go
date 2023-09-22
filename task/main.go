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
    "github.com/aws/aws-sdk-go/service/sqs"
    "github.com/google/uuid"
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

    for {
        // Receive messages from the SQS queue.
        receiveParams := &sqs.ReceiveMessageInput{
            QueueUrl:            aws.String(queueURL),
            MaxNumberOfMessages: aws.Int64(10),
            WaitTimeSeconds:     aws.Int64(20),
        }

        receiveResult, err := sqsSvc.ReceiveMessage(receiveParams)
        if err != nil {
            fmt.Println("Error receiving message:", err)
            os.Exit(1)
        }

        // Process the received messages.
        for _, message := range receiveResult.Messages {
            // Replace with your S3 bucket name.
            bucketName := "orders1"

            // Generate the folder structure based on date, shopId, and typeofpayload.
            currentDate := time.Now().Format("2006-01-02") // Format the current date as YYYY-MM-DD
            shopID := "12345"                             // Replace with the actual shop ID from your message
            payloadType := "somepayload"                  // Replace with the actual payload type from your message

            // Generate a unique UUID for the file name.
            uuid := generateUUID()

            // Construct the S3 object key.
            objectKey := fmt.Sprintf("%s/%s/%s/%s.json.gz", currentDate, shopID, payloadType, uuid)

            // Create a gzip writer to compress the data.
            var buf bytes.Buffer
            writer := gzip.NewWriter(&buf)
            _, err := writer.Write([]byte(*message.Body))
            if err != nil {
                fmt.Println("Error compressing data:", err)
                os.Exit(1)
            }
            writer.Close()

            // Put the compressed message content into an S3 object.
            _, err = s3Svc.PutObject(&s3.PutObjectInput{
                Bucket:        aws.String(bucketName),
                Key:           aws.String(objectKey),
                Body:          bytes.NewReader(buf.Bytes()),
                ContentLength: aws.Int64(int64(len(buf.Bytes()))),
            })
            if err != nil {
                fmt.Println("Error putting object in S3:", err)
                os.Exit(1)
            }

            // Delete the processed message from the SQS queue.
            _, err = sqsSvc.DeleteMessage(&sqs.DeleteMessageInput{
                QueueUrl:      aws.String(queueURL),
                ReceiptHandle: message.ReceiptHandle,
            })
            if err != nil {
                fmt.Println("Error deleting message:", err)
            } else {
                fmt.Println("Processed message:", *message.Body)
            }
        }

    }
}

// Function to generate a unique UUID
func generateUUID() string {
    return uuid.New().String()
}