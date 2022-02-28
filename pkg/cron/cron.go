package cron

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func RunEveryMinute() {
	sched, _ := cron.ParseStandard("* * * * *")

	t := time.Now().Add(-5 * time.Minute)
	for ; !t.After(time.Now()); t = sched.Next(t) {
		fmt.Println(t)
	}
}
