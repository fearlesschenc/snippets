package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getNosService() *s3.S3 {
	cred := credentials.NewStaticCredentials("595da7952d29400680b2437a1c9fc26f", "11b22a95ee304f168b3f3d426296ca85", "")
	config := aws.NewConfig().
		WithCredentials(cred).
		WithRegion("us-east-1").
		WithEndpoint("nos-jd.163yun.com")
	return s3.New(session.Must(session.NewSession()), config)
}

func getMinioService() *s3.S3 {
	cred := credentials.NewStaticCredentials("Q3AM3UQ867SPQQA43P2F", "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG", "")
	config := aws.NewConfig().
		WithCredentials(cred).
		WithRegion("us-east-1").
		WithEndpoint("play.min.io").
		WithS3ForcePathStyle(true)
	return s3.New(session.Must(session.NewSession()), config)
}

func CreateBucket() {
	svc := getNosService()
	//svc := getMinioService()
	output, err := svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String("nzk1942"),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String("us-east-1"),
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	_ = output
}

func ListBucketObject() {
	svc := getNosService()
	//svc := getMinioService()
	output, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String("nzk1942"),
	})
	if err != nil {
		fmt.Println(err)
	}
	_ = output
}
