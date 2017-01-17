package myinfluxdb

import (
	"reflect"
	"testing"
)

var influxDBTestHTTP = InfluxConfig{Hostname: "localhost", Protocol: "http", Port: "8092",
	Db:       "test",
	Username: "username",
	Password: "password",
	Tags:     map[string]string{"addkey1": "value1", "addkey2": "value2"}}

var influxDBTestUDP = InfluxConfig{Hostname: "localhost", Protocol: "udp", Port: "8092",
	Db:       "test",
	Username: "username",
	Password: "password",
	Tags:     map[string]string{"addkey1": "value1", "addkey2": "value2"}}

var InfluxMetricSample = InfluxMetric{
	Tags: map[string]string{
		"key1": "value1",
		"key2": "value2",
	},
	Values: map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}, Measurement: "key1"}

var TT []InfluxMetric

func TestSend(t *testing.T) {
	TT = append(TT, InfluxMetricSample)
	TT = append(TT, InfluxMetricSample)

	t.Log("Testing: Send()")
	//influxDBTestHTTP.Send(&TT)
	//	influxDBTestUDP.Send(&TT)

}

func TestAddTags(t *testing.T) {

	result := map[string]string{"addkey1": "value1", "addkey2": "value2", "key1": "value1", "key2": "value2"}

	t.Log("Testing: AddTags()")
	metric := InfluxMetricSample

	metric.AddTags(influxDBTestHTTP.Tags)

	if !reflect.DeepEqual(metric.Tags, result) {
		t.Errorf("Added tags are different from what we expect")
	}
}
