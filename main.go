package main

import (
    "fmt"
    "math/rand"
    "time"

    "net/http"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func Healthz(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "OK")
}

type Randomizer struct {
    RandImpl func() float64
}


func recordMetrics(randomizer Randomizer) {
    gauge := prometheus.NewGauge(
        prometheus.GaugeOpts{
            Namespace: "app",
            Name:      "random_metric",
            Help:      "This is a random metric",
        })

    prometheus.MustRegister(gauge)

    go func() {
        for {
            gauge.Set(randomizer.RandImpl())

            time.Sleep(time.Second)
        }
    }()

}

func main() {
    rand.Seed(time.Now().UnixNano())
    randomizer := Randomizer{
        RandImpl: func() float64 { return rand.Float64() },
    }

    recordMetrics(randomizer)

    http.Handle("/metrics", promhttp.Handler())
    http.HandleFunc("/healthz", Healthz)
    
    http.ListenAndServe(":3000", nil)
}
