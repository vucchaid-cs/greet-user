package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	sdk "github.com/gaia-pipeline/gosdk"
)

func main() {
	if err := sdk.Serve(PipelinesJobs); err != nil {
		log.Panic(err.Error())
	}
}

var PipelinesJobs = sdk.Jobs{
	sdk.Job{
		Handler:     greetingMessage,
		Title:       "Greet User",
		Description: "Test Greet User",
	},
}

func greetingMessage(args sdk.Arguments) error {
	loc, _ := time.LoadLocation("Asia/Calcutta")
	FormattedVersion := time.Now().In(loc).Format("15:04 PM")
	if strings.Contains(FormattedVersion, "AM") {
		fmt.Printf("Oh. it's already %v in India, Good Morning!", FormattedVersion)
	}
	if strings.Contains(FormattedVersion, "PM") {
		fmt.Printf("Oh. it's already %v in India, Good Evening!", FormattedVersion)
	}
	return nil
}
