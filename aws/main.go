package main

import (
	"aws/pkg"
)

func main() {
	cfg := pkg.AwsLoadConfig()
	pkg.ListBucket(pkg.S3Client(cfg))
}
