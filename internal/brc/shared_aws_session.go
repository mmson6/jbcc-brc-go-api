package brc

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var awsSess *session.Session
var awsSessOnce sync.Once

func SharedAWSSession(ctx context.Context) *session.Session {
	awsSessOnce.Do(func() {
		// uoCfg := BuildConfig(ctx)

		// Build the AWS configuration from the UO configuration

		conf := aws.NewConfig().
			WithMaxRetries(3)
		// if uoCfg.Debug {
		// 	conf = conf.WithLogLevel(aws.LogDebugWithRequestErrors)
		// }

		// Create the session

		awsSess = session.Must(session.NewSession(conf))
	})

	return awsSess
}
