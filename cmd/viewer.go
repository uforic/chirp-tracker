package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

func getInfluxClient() client.Client {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://influx:8086",
	})
	if err != nil {
		log.Fatal(err)
	}
	return c
}

var influxClient = getInfluxClient()
var databaseName = "chirps"

type message struct {
	ZipID string
	Lat   float64
	Lon   float64
	Time  int64
}

func influx(w http.ResponseWriter, r *http.Request) {

	res, _ := queryDB(influxClient, "SELECT * FROM position ORDER BY time DESC LIMIT 1")
	for _, row := range res[0].Series[0].Values {
		t, err := time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			log.Fatal(err)
		}
		lat, err := row[1].(json.Number).Float64()
		lon, err := row[2].(json.Number).Float64()
		m := message{"zip1", lat, lon, t.Unix()}
		b, err := json.Marshal(m)
		// line := fmt.Sprintf("[%2d] %s %s %d %d", i, t, row[1], row[2], row[3])
		io.WriteString(w, string(b))
	}
}

// queryDB convenience function to query the database
func queryDB(clnt client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: databaseName,
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

var indexTemplate, err = template.ParseFiles("ui/assets/index.html")

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		err = indexTemplate.Execute(w, nil)
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func main() {
	http.HandleFunc("/influx", influx)
	http.HandleFunc("/bundle.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./ui/assets/bundle.js")
	})
	http.HandleFunc("/bundle.js.map", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./ui/assets/bundle.js.map")
	})
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
