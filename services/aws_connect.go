package services

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func ConnectAWS() *session.Session{
	AccessKeyId := os.Getenv("ACCESS_KEY_ID")
	SecretAccessKey := os.Getenv("ACCESS_KEY_KEY")
	MyRegion := os.Getenv("AWS_REGION")

	fmt.Println(MyRegion)

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyId,
				SecretAccessKey,
				"",
			),
		})
	
	if err != nil {
		panic(err)
	}

	return sess
}