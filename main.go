//Query Scheduler Service
package main

import (
	"fmt"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

func main() {
	fmt.Print("QSS started \n")
	conf := parceYAML()

	conn := getConnection("http://"+conf.Influx.Host+":"+conf.Influx.Port, conf.Influx.Login, conf.Influx.Password)
	defer conn.Close()

	send := make(chan client.Query)
	scheduler := getScheduler(send)
	queries := parseQueries(conf.Queries)

	for _, query := range queries {
		go scheduler(getQuery(query), query.Interval)
	}

	for {
		query := <-send
		fmt.Print("TICK \n")
		go runQuery(conn, query)
	}

}
