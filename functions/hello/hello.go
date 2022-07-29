package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	mm_url := "https://chat.gameloft.org/hooks/bcbu761yb3bufkmianijxjjfze"
	type MatterMostPayload struct {
		Channel string
		Text    string
	}
	var payload MatterMostPayload = MatterMostPayload{
		Channel: "@vu.nguyenhoang",
		Text:    "Your Computer booted up",
	}

	var byteArr, _ = json.Marshal(payload)
	req, err := http.NewRequest("POST", mm_url, bytes.NewBuffer(byteArr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, World!",
	}, nil
}

func main() {
	lambda.Start(handler)
}
