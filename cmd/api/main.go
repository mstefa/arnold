package main

import (
	"arnold/cmd/api/bootstrap"
	"arnold/internal/cronjob"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}

	cronjob.CronJob.Init()
}
