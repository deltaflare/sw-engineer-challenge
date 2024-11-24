package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"math/rand"

	"github.com/OtchereDev/deltaflare/daemon/models"
	influxdb "github.com/influxdata/influxdb-client-go/v2"
)

func CreateEvent() models.Event {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	return models.Event{
		Criticality:  r.Intn(5) + 1,
		Timestamp:    time.Now(),
		EventMessage: "New event created",
	}
}

func SaveEventToDB(client influxdb.Client, event models.Event) error {
	orgName := os.Getenv("INFLUXDB_ORG")
	bucketName := os.Getenv("INFLUXDB_BUCKET")

	if orgName == "" || bucketName == "" {
		return errors.New("INFLUXDB_ORG or INFLUXDB_BUCKET is not provided")
	}

	writeAPI := client.WriteAPIBlocking(orgName, bucketName)

	p := influxdb.NewPoint(
		"security_event",
		map[string]string{
			"source":       "daemon",
			"eventMessage": event.EventMessage,
		},
		map[string]interface{}{
			"criticality": event.Criticality,
		},
		event.Timestamp,
	)

	if err := writeAPI.WritePoint(context.Background(), p); err != nil {
		fmt.Println("Error writing event to InfluxDB:", err)
		return err
	}

	return nil
}
