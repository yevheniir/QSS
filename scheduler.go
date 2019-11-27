// Scheduler for running queries
package main

import (
	"time"

	_ "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client/v2"
)

type scheduler = func(client.Query, int)

func getScheduler(send chan client.Query) scheduler {
	return func(query client.Query, interval int) {
		for range time.Tick(time.Duration(interval) * time.Second) {
			send <- query
		}
	}
}
