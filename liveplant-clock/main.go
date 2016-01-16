package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func main() {
	c := cron.New()
	c.AddFunc("30 * * * * *", func() {
		t := time.Now()
		fmt.Println(t.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
	})
	c.Start()
	defer c.Stop()
	select {}
}
