package initializers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var AwsSession *session.Session

func InitializeSession(AWS_REGION string) {
	AwsSession = session.Must(
		session.NewSession(
			&aws.Config{
				Region:     aws.String(AWS_REGION),
				MaxRetries: aws.Int(3),
			},
		),
	)
}
