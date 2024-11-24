package main

import (
	"log"
	"os"
	"time"

	"github.com/OtchereDev/deltaflare/daemon/service"
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

	log.Println("Starting Daemon service ...")

	for {
		event := service.CreateEvent()
		log.Printf("Generated event: %+v\n", event)

		if err := service.SaveEventToDB(client, event); err != nil {
			log.Printf("Failed to save event: %v\n", err)
		}

		time.Sleep(5 * time.Second)
	}
}
