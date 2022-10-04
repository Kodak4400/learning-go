package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func ListBucket(client *s3.Client) {
	// Get the first page of results for ListObjectsV2 for a bucket
	listBucketsResult, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to AWS S3 ListBuckets")
	}

	log.Println("first page results:")
	for _, bucket := range listBucketsResult.Buckets {
		fmt.Printf("Bucket name: %s\t\tcreated at: %v\n", *bucket.Name, bucket.CreationDate)
	}
}
