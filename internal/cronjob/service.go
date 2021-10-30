package cronjob

import (
	"fmt"
	_ "github.com/robfig/cron/v3"
)

c := cron.New()
c.AddFunc("0 30 * * * *", func(){ fmt.Println("Every hour on the half hour") })
