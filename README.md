# Example usage

```
package main

import (
    db "github.com/DanKans/myinfluxdb"
)

func main() {
 
    var (
        InfluxConfig db.InfluxConfig
        Metric db.InfluxMetric
        Metrics []db.InfluxMetric
    )

    // Connection 
    InfluxConfig.Hostname = "127.0.0.1"
    InfluxConfig.Protocol = "udp"
    InfluxConfig.Port = "8086"
    InfluxConfig.Tags = map[string]string{"name": "Collector1"}

    // Sample metric
    Metric.Measurement = "temperature"
    Metric.Tags = map[string]string {
        "room":     "bathroom",
    }
    Metric.Values = map[string]interface{} {
        "probe1": 22,
        "probe2": 21,
    }
 
    // Append sample Metric to array of Metrics   
    Metrics = append(Metrics, Metric)

    InfluxConfig.Send(&Metrics)

}
```
