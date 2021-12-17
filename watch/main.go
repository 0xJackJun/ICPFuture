package main

import (
	"fmt"
	"github.com/robfig/cron"
	"watch/service"
)

func main() {
	// Initialize the HTTP service
	r := service.InitHttp()
	c := cron.New()

	err := c.AddFunc("@every 1h", service.Watch)
	if err != nil {
		fmt.Println(err)
	}
	c.Start()
	r.Run()
}
