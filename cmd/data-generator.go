package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

const (
	DB = "chirps"
)

var coordsMuhanga = []float64{-1.9209437, 29.5748399}

func main() {
	time.Sleep(time.Second * time.Duration(3))
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://influx:8086",
	})
	if err != nil {
		log.Fatal(err)
	}
	_, err = queryDB(c, fmt.Sprintf("CREATE DATABASE %s", DB))
	if err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(time.Second)
		writePoints(c)
	}
}

// queryDB convenience function to query the database
func queryDB(clnt client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: DB,
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

func writePoints(clnt client.Client) {
	numberOfZips := 2

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "chirps",
		Precision: "us",
	})
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numberOfZips; i++ {
		tags := map[string]string{
			"zip": fmt.Sprintf("zip%d", i),
		}

		idle := rand.Float64() * .1
		fields := map[string]interface{}{
			"lat": coordsMuhanga[0] + idle,
			"lon": coordsMuhanga[1] + idle,
		}

		pt, err := client.NewPoint(
			"position",
			tags,
			fields,
			time.Now(),
		)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)
	}

	if err := clnt.Write(bp); err != nil {
		log.Fatal(err)
	}
}
