package main

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Request es la estructura de la solicitud
type Request struct {
	Number string `json:"number"`
}

// Response es la estructura de la respuesta
type Response struct {
	Result  int    `json:"result"`
	Mensaje string `json:"mensaje"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	numberStr := request.QueryStringParameters["number"]
	if numberStr == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       `{"error": "Number parameter is required"}`,
		}, nil
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       `{"error": "Invalid number"}`,
		}, nil
	}

	result := number * 2
	response := Response{Result: result, Mensaje: "Si funciona, esta automatizado con github actions - prueba"}
	responseBody, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       `{"error": "Internal Server Error"}`,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBody),
	}, nil
}

func main() {
	lambda.Start(handler)
}
