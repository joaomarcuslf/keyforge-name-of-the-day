package routines

import (
	"fmt"

	"github.com/joaomarcuslf/keyforge-name-of-the-day/usecases"
	"github.com/robfig/cron/v3"
)

type CronJob struct {
	cron *cron.Cron
}

func NewCronJob() *CronJob {
	return &CronJob{cron: cron.New()}
}

func (c *CronJob) Run() {
	c.cron.AddFunc("@every 1d", func() {
		kfr, _ := usecases.SendTweet()
		fmt.Println("Sent tweet: '" + kfr.Name + "'")
	})

	c.cron.Start()
}
