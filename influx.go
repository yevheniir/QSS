//All influxDb tools
package main

import (
	"fmt"
	"os"

	_ "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client/v2"
)

func getConnection(addr string, username string, password string) client.Client {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     addr,
		Username: username,
		Password: password,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
		os.Exit(1)
	}

	return c
}
