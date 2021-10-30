package cronjob

import (
	"fmt"
	"github.com/robfig/cron/v3"
	_ "github.com/robfig/cron/v3"
	"time"
)


type CronJob struct {
}

func (cj CronJob) Init(){
	c := cron.New()
	c.AddFunc("0 30 * * * *", func(){ fmt.Println("Every hour on the half hour") })


	// Start cron with one scheduled job
	c.Start()
	printCronEntries(c.Entries())
	time.Sleep(2 * time.Minute)
}

func printCronEntries(cronEntries []cron.Entry) {
	fmt.Printf("Cron Info: %+v\n", cronEntries)
}