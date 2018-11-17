package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mikejoh/stryktipset"
)

func convert(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var c stryktipset.Convert

	c.Sek, _ = strconv.Atoi(req.QueryStringParameters["sek"])
	full, half := stryktipset.ConvertSekToBet(c.Sek)
	c.Full = full
	c.Half = half

	json, _ := json.Marshal(c) // Skip error handling for now

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(json),
	}, nil
}

func main() {
	lambda.Start(convert)
}
