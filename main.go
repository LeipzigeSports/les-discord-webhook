package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type Config struct {
	Data map[string]JobConfig `json:"data"`
}

type JobConfig struct {
	Cron       string `json:"cron"`
	WebhookURL string `json:"webhook_url"`
	// RoleID     string `json:"role_id"`
	Message string `json:"message"`
}

func main() {
	fmt.Println("Hola!")

	slog.Info("Config: read file...")
	data, err := os.ReadFile("/app/config.json")
	if err != nil {
		slog.Error("Config: read failure!")
		log.Fatal(err)
	} else {
		slog.Info("Config: read.")
	}

	var config Config

	err = json.Unmarshal(data, &config)
	if err != nil {
		slog.Error("Config: unmarshal failure!")
		log.Fatal(err)
	} else {
		slog.Info("Config: struct built.")
	}

	count := len(config.Data)
	slog.Info("Config: " + strconv.Itoa(count) + " jobs loaded.")

	slog.Info("Scheduler: create...")
	s, err := gocron.NewScheduler()
	defer func() { _ = s.Shutdown() }()

	if err != nil {
		slog.Error("Scheduler: creation failed!")
		log.Fatal(err)
	} else {
		slog.Info("Scheduler: creation successful.")
	}

	var i = 1

	for key, job := range config.Data {
		slog.Info("Job: create... " + key + " [" + strconv.Itoa(i) + "/" + strconv.Itoa(count) + "]")

		// TODO: ensure WebhookURL starts with https://discrod.com
		j, err := s.NewJob(gocron.CronJob(job.Cron, false),
			gocron.NewTask(
				func() {
					// source: https://stackoverflow.com/questions/66842959/submit-variable-in-payload-of-golang-http-newrequest
					type dingsStruct struct {
						Content string `json:"content"`
					}
					dingsData := dingsStruct{
						// Content: "<@&" + job.RoleID + "> " + job.Message,
						Content: job.Message,
					}
					dingsBytes, err := json.Marshal(dingsData)
					dingsPayload := bytes.NewBuffer(dingsBytes)

					// fmt.Println("Payload:", string(dingsBytes)) // troubleshoot dings

					// source: https://stackoverflow.com/questions/16895294/how-to-set-timeout-for-http-get-requests-in-golang
					client := http.Client{
						Timeout: 5 * time.Second,
						Transport: &http.Transport{
							DisableKeepAlives: true, // Force new connection each time
						},
					}
					resp, err := client.Post(job.WebhookURL, "application/json", dingsPayload)
					if err != nil {
						// panic(err)
						// log.Fatal(err)
						slog.Error("Job: FAIL: "+key+" --> network error;", "details", err)
						return
					}
					defer resp.Body.Close()

					// TODO: handle http response by status code
					slog.Info("Job: OK: " + key + " [HttpStatus:" + strconv.Itoa(resp.StatusCode) + "] " + job.Message)
				},
			),
		)
		if err != nil {
			slog.Error("Job: creation failed!")
			log.Fatal(err)
		}

		slog.Info("Job: created. " + key + " [" + strconv.Itoa(i) + "/" + strconv.Itoa(count) + "] ID: " + j.ID().String())

		i++
	}

	// gnJob := config.Data["gn"]
	// fmt.Println(gnJob.Cron)

	s.Start()
	slog.Info("Scheduler: lesgooo!")

	// source: https://jacobtomlinson.dev/posts/2022/golang-block-until-interrupt-with-ctrl-c/
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig // wait for interupt signal (ctrl-c)

	slog.Info("Scheduler: shutting down...")
	s.Shutdown()

	err = s.Shutdown()
	if err != nil {
		log.Fatal(err)
		s.Shutdown()
	}
}
