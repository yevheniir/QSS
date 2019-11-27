// Queries list
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client/v2"
)

//Query JSON query
type Query struct {
	Db       string `json:"db"`
	From     string `json:"from"`
	To       string `json:"to"`
	Interval int    `json:"interval"`
}

func parseQueries(file string) []Query {

	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var queries = []Query{}
	json.Unmarshal(byteValue, &queries)

	return queries
}

func getQuery(query Query) client.Query {
	return client.NewQuery(fmt.Sprintf(`SELECT SUM(*), COUNT(*), MIN(*), MAX(*)
			INTO "%s".:MEASUREMENT 
			FROM "%s"./.*/ 
			GROUP BY time(%vs)`, query.To, query.From, query.Interval), query.Db, "")
}

func runQuery(c client.Client, query client.Query) {
	if response, err := c.Query(query); err == nil && response.Error() == nil {
		fmt.Println(response.Results)
	} else {
		fmt.Println(err, response.Error())
	}
}
