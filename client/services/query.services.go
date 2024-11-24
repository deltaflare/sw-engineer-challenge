package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/OtchereDev/deltaflare/cient/models"
	influxdb "github.com/influxdata/influxdb-client-go/v2"
)

func FetchEvents(client influxdb.Client) error {
	orgName := os.Getenv("INFLUXDB_ORG")
	bucketName := os.Getenv("INFLUXDB_BUCKET")

	if orgName == "" || bucketName == "" {
		return errors.New("INFLUXDB_ORG or INFLUXDB_BUCKET is not provided")
	}

	queryAPI := client.QueryAPI(orgName)
	query := fmt.Sprintf(`from(bucket: "%s")
			|> range(start: -4h)
			|> filter(fn: (r) => r._measurement == "security_event")
		 	|> filter(fn: (r) => r._field == "criticality")
			|> sort(columns: ["_value"], desc: true)
			|> limit(n: 10)`, bucketName)

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		fmt.Println("Error querying InfluxDB:", err)
		return err
	}
	defer result.Close()

	for result.Next() {
		event := models.Event{
			Criticality:  int(result.Record().Values()["_value"].(int64)),
			EventMessage: result.Record().Values()["eventMessage"].(string),
			Timestamp:    result.Record().Values()["_time"].(time.Time).Format(time.RFC3339),
		}

		DisplayEvent(event)
	}

	return result.Err()

}

func DisplayEvent(event models.Event) {
	fmt.Printf("Criticality: %v | EventMessage: %v | Timestamp: %v\n",
		event.Criticality, event.EventMessage, event.Timestamp)
}
