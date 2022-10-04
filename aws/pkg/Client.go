package pkg

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func S3Client(cfg aws.Config) *s3.Client {
	// Create an Amazon S3 service client
	return s3.NewFromConfig(cfg)
}
