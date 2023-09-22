package main

import (
	"bytes"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// BenchmarkS3Upload benchmarks the performance of uploading data to an S3 bucket.
func BenchmarkS3Upload(b *testing.B) {
	// Create an AWS session.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create an S3 client.
	s3Svc := s3.New(sess)

	// Create an S3 uploader with a buffer.
	uploader := s3manager.NewUploaderWithClient(s3Svc, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024 // 5MB part size for multipart upload
	})

	// Replace with your S3 bucket name.
	bucketName := "orders1"

	// Create sample data to upload.
	data := []byte("This is a sample data to upload to S3.")

	// Reset the timer to exclude setup time.
	b.ResetTimer()

	// Run the code you want to benchmark in a loop.
	for i := 0; i < b.N; i++ {
		// Generate a unique object key for each upload.
		objectKey := generateUUID() + ".txt"

		// Upload the data to S3.
		_, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
			Body:   bytes.NewReader(data),
		})
		if err != nil {
			b.Fatalf("Error uploading to S3: %v", err)
		}
	}
}

