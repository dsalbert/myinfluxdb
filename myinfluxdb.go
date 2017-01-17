package myinfluxdb

import (
	"log"

	influx "github.com/influxdata/influxdb/client/v2"
)

// InfluxConfig - Initial configuration for influxdb
type InfluxConfig struct {
	Hostname string            `json:"hostname"`
	Protocol string            `json:"protocol"`
	Port     string            `json:"port"`
	Db       string            `json:"db"`
	Username string            `json:"username"`
	Password string            `json:"password"`
	Tags     map[string]string `json:"tags"`
}

// InfluxMetric - Metric strcuture for InfluxDB
type InfluxMetric struct {
	Tags        map[string]string
	Values      map[string]interface{}
	Measurement string
}

// Send - Send metric according to the configuration in receiver
func (config *InfluxConfig) Send(metrics *[]InfluxMetric) {

	var (
		Client      influx.Client
		BatchPoints influx.BatchPoints
		Point       *influx.Point
		err         error
	)

	BatchPoints, err = influx.NewBatchPoints(influx.BatchPointsConfig{
		Database:  config.Db,
		Precision: "s",
	})

	if err != nil {
		log.Panic("Error: Cannot create BatchPoints", err)
	}

	// Inject tags in case if there are some
	if len(config.Tags) > 0 {
		for _, MetricValue := range *metrics {
			MetricValue.AddTags(config.Tags)
		}
	}

	for _, MetricVal := range *metrics {
		Point, err = influx.NewPoint(MetricVal.Measurement,
			MetricVal.Tags,
			MetricVal.Values)

		if err != nil {
			log.Println("Warning: Problem with data point", err)
		} else {
			BatchPoints.AddPoint(Point)
		}
	}
	// Create Client

	switch config.Protocol {

	case "udp":
		Client, err = influx.NewUDPClient(influx.UDPConfig{
			Addr: config.Hostname + ":" + config.Port,
		})

	case "http":
		Client, err = influx.NewHTTPClient(influx.HTTPConfig{
			Addr:     config.Hostname + ":" + config.Port,
			Username: config.Username,
			Password: config.Password,
		})

	default:
		log.Panic("Error: Unknown protocol ", config.Protocol)
	}

	if err != nil {
		log.Panic("Error: Cannot connect ", err)
	} else {
		Client.Write(BatchPoints)
		Client.Close()
	}
}

// AddTags - Inject tags to current metric
func (metric *InfluxMetric) AddTags(tags map[string]string) {
	for tagKey, tagValue := range tags {
		metric.Tags[tagKey] = tagValue
	}
}
