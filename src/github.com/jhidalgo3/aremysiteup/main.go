package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/jhidalgo3/aremysiteup/notify"
	"github.com/jhidalgo3/aremysiteup/params"
)

func check(client *http.Client, u string) error {
	resp, err := client.Get(u)
	if err != nil || resp.StatusCode != 200 {
		var checkError error
		if err != nil {
			checkError = err
		} else {
			checkError = &url.Error{Op: "Get", URL: u, Err: fmt.Errorf("status code is %d", resp.StatusCode)}
		}
		return checkError
	}
	return nil
}

func main() {
	var (
		config params.Config

		sleep    time.Duration
		mailer   *notify.Mailer
		client   *http.Client
		errors   []error
		errCount int
	)
	params.Load(&config)

	if len(config.Urls) == 0 {
		log.Println("Nothing to check. Exiting...")
		return
	}
	client = &http.Client{Timeout: time.Second * time.Duration(config.Timeout)}
	mailer = &notify.Mailer{Mailgun: config.Mailgun, To: config.To, From: config.From}
	log.Printf("Starting to check these urls => %v...\n", config.Urls)
	for {
		errors = make([]error, 0)
		for _, url := range config.Urls {
			err := check(client, url)
			if err != nil {
				errors = append(errors, err)
				log.Printf("Error for '%s': %s\n", url, err.Error())
			}
		}
		errCount = len(errors)
		if errCount == 0 {
			sleep = time.Second * time.Duration(config.Sleep)
		} else {
			mailer.ComposeAndSendEmail(errors)
			sleep = time.Second * time.Duration(config.SleepWithError)
		}
		if config.OneCheck {
			// f(x) = [e^(-1/x^2)]
			//exitStatus := int(math.Ceil(math.Pow(2.7182818284, -1/math.Pow(float64(errCount), 2.0))))
			//os.Exit(exitStatus)
			if errCount == 0 {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		}
		time.Sleep(sleep)
	}
}
