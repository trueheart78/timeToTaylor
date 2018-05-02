package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/trueheart78/timeToTaylor/taylorTime"
)

type Request struct {
	Location string `json:"location"`
}

type TimeToShow struct {
	Showtime   string `json:"showtime"`
	TimeToShow string `json:"time_to_show"`
}

func HandleLambdaEvent(event Request) (TimeToShow, error) {
	return TimeToShow{Showtime: taylorTime.Banner(), TimeToShow: taylorTime.TimeRemaining()}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
