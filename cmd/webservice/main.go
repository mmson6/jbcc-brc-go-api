// package main

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"

// 	"github.com/aws/aws-lambda-go/events"
// 	"github.com/aws/aws-lambda-go/lambda"
// )

// // Response is of type APIGatewayProxyResponse since we're leveraging the
// // AWS Lambda Proxy Request functionality (default behavior)
// //
// // https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
// type Response events.APIGatewayProxyResponse

// // Handler is our lambda handler invoked by the `lambda.Start` function call
// func Handler(ctx context.Context) (Response, error) {
// 	var buf bytes.Buffer

// 	body, err := json.Marshal(map[string]interface{}{
// 		"message": "Go Serverless v1.0! Your function executed successfully!",
// 	})
// 	if err != nil {
// 		return Response{StatusCode: 404}, err
// 	}
// 	json.HTMLEscape(&buf, body)

// 	resp := Response{
// 		StatusCode:      200,
// 		IsBase64Encoded: false,
// 		Body:            buf.String(),
// 		Headers: map[string]string{
// 			"Content-Type":           "application/json",
// 			"X-MyCompany-Func-Reply": "hello-handler",
// 		},
// 	}

// 	return resp, nil
// }

// func main() {
// 	lambda.Start(Handler)
// }

package main

import (
	"context"

	"github.com/jbcc/brc-api/internal/webservice"
)

////////////////////////////////////////////////////////////////////////////////
// MAIN

func main() {
	ctx := context.Background()

	// Start the web service
	svc := webservice.BRC(ctx)
	svc.Start()
}

/*
aws lambda create-function --region us-east-1 --function-name test-brc-api --zip-file fileb://./deployment.zip --runtime go1.x --role arn:aws:iam::666885501734:role/service-role/lambda-apigateway-role --handler main
*/
