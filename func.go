package main

import (
	"context"
	"encoding/json"
	"io"
	"log"

	weatherapi "github.com/belovai/goopenweathermapapi"
	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))
}

type Request struct {
	City  string `json:"city"`
	Units string `json:"units"`
	Lang  string `json:"lang"`
}

type contextKey string

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {
	var request Request
	json.NewDecoder(in).Decode(&request)

	apiKey := fdk.GetContext(ctx).Config()["OPEN_WEATHER_MAP_API_KEY"]

	if apiKey == "" {
		log.Fatal("OpenWeatherMap api key is not configured")
	}

	client := weatherapi.NewClient(apiKey)

	if request.City == "" {
		log.Fatal("City name is required")
	}

	response, err := client.GetWeatherByCityName(request.City, request.Units, request.Lang)

	if err != nil {
		log.Fatal(err)
	}

	out.Write([]byte(response))
}
