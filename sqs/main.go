package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func sendMessage(i int, svc *sqs.SQS, queueURL string) {
	message := fmt.Sprintf(`{"customerCreate": {"userErrors": [], "customer": {"id": "gid://shopify/Customer/%d", "email": "customer%d@example.com", "phone": "+123456789%d", "taxExempt": false, "acceptsMarketing": true, "firstName": "Customer%d", "lastName": "Lastname%d", "ordersCount": "0", "totalSpent": "0.00", "smsMarketingConsent": {"marketingState": "NOT_SUBSCRIBED", "marketingOptInLevel": "SINGLE_OPT_IN"}, "addresses": [{"address1": "123 Main St", "city": "New York", "country": "USA", "phone": "+11223344%d", "zip": "10001"}}]}}`, i, i, i, i, i, i)

	// Include a unique MessageGroupId and MessageDeduplicationId for each message
	messageGroupId := fmt.Sprintf("Group%d", i)
	messageDeduplicationId := fmt.Sprintf("Dedupe%d", i)

	_, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageBody:           aws.String(message),
		QueueUrl:              &queueURL,
		MessageGroupId:        aws.String(messageGroupId),
		MessageDeduplicationId: aws.String(messageDeduplicationId),
		DelaySeconds:          aws.Int64(0),
	})
	if err != nil {
		fmt.Printf("Error sending message %d: %v\n", i, err)
		return
	}
	fmt.Printf("Message %d sent successfully\n", i)
}

func main() {
	// Initialize an AWS session based on your AWS credentials and region
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Replace with your desired region
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	// Create an SQS service client
	svc := sqs.New(sess)

	// Define the URL of your FIFO SQS queue
	queueURL := "https://sqs.us-east-1.amazonaws.com/992925780889/orders.fifo"

	// Send 500 messages with unique MessageGroupIds and MessageDeduplicationIds using Goroutines
	for i := 1; i <= 500; i++ {
		go sendMessage(i, svc, queueURL)
	}

	// Sleep for some time to allow Goroutines to complete
	// You may need to adjust this duration based on your actual use case
	// or use a WaitGroup to wait for all Goroutines to finish.
	// For this example, we'll keep it simple with a Sleep.
	select {}
}
