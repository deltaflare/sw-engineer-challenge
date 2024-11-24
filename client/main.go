package main

import (
	"log"
	"os"

	"github.com/OtchereDev/deltaflare/cient/services"
	influxdb "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	dbUrl := os.Getenv("INFLUXDB_URL")
	influxToken := os.Getenv("INFLUX_TOKEN")

	if dbUrl == "" || influxToken == "" {
		log.Fatal("INFLUXDB_URL or INFLUX_TOKEN is not provided")
	}

	client := influxdb.NewClient(dbUrl, influxToken)
	defer client.Close()

	// Query for critical events
	err := services.FetchEvents(client)

	if err != nil {
		log.Fatal(err)
	}
}
