package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/trueheart78/timeToTaylor/taylorTime"
)

type Request struct {
	Location int `json:"location"`
}

type TimeToShow struct {
	Showtime   string `json:"showtime"`
	TimeToShow string `json:"time_to_show"`
}

func Handler(request Request) (TimeToShow, error) {
	tts := TimeToShow{taylorTime.Banner(), taylorTime.TimeRemaining()}

	return tts, nil
}

func main() {
	lambda.Start(Handler)
}
