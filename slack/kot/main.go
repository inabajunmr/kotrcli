package main

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/inabajunmr/kotrcli"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// valification token
	token := os.Getenv("VERIFICATION_TOKEN")
	vals, _ := url.ParseQuery(request.Body)
	if vals.Get("token") != token {
		// invalid token
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Invalid token."),
			StatusCode: 200,
		}, nil
	}

	kotToken := os.Getenv(vals.Get("user_id"))
	if len(kotToken) == 0 {
		return events.APIGatewayProxyResponse{
			Body:       "You must register token.",
			StatusCode: 200,
		}, nil
	}

	// syukkin or taikin
	kind := vals.Get("text")

	// kot token format is "user_token:token"
	tokenSet := strings.Split(kotToken, ":")
	if len(tokenSet) != 2 {
		return events.APIGatewayProxyResponse{
			Body:       "Wrong token format. Token format is 'user_token:token'",
			StatusCode: 200,
		}, nil
	}

	buf := new(bytes.Buffer)
	kotrcli.Dakoku(buf, getSyukkinTaikin(kind), tokenSet[0], tokenSet[1])

	return events.APIGatewayProxyResponse{
		Body:       buf.String(),
		StatusCode: 200,
	}, nil
}

func getSyukkinTaikin(kind string) kotrcli.Type {
	if kind == "syukkin" {
		return kotrcli.SYUKKIN
	} else if kind == "taikin" {
		return kotrcli.TAIKIN
	}

	return -1
}

func main() {
	lambda.Start(handler)
}
