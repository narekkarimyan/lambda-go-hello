package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"encoding/json"
	"net/http"
)

func hello(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	type BodyMessage struct {
		Message string
	}
	response := BodyMessage{
		Message: "Здорово!",
	}
	return apiResponse(http.StatusOK, response)
}

func main() {
	lambda.Start(hello)
}

func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status
	
	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return &resp, nil
}