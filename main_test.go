package main

import (
    "fmt"
    "strings"
    "net/http"
    "net/http/httptest"

    "testing"

    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func Test_RandomPrometheusMetric(t *testing.T) {
    req, err := http.NewRequest("GET", "http://localhost:3000/metrics", nil)
    if err != nil {
        t.Fatal(err)
    }

    randomizer := Randomizer{
        RandImpl: func() float64 { return 1234567.89 },
    }
    recordMetrics(randomizer)

    res := httptest.NewRecorder()
    promhttp.Handler().ServeHTTP(res, req)

    must_include := "app_random_metric 1.23456789e+06"
    act := res.Body.String()

    if !strings.Contains(act, must_include) {
        t.Fatalf("Expected `%s` inside %s", must_include, act)
    }
}

func Test_Healthz(t *testing.T) {
    req, err := http.NewRequest("GET", "http://localhost:3000/healthz", nil)
    if err != nil {
        t.Fatal(err)
    }

    res := httptest.NewRecorder()
    Healthz(res, req)

    exp := fmt.Sprintf("OK")
    act := res.Body.String()
    if exp != act {
        t.Fatalf("Expected %s got %s", exp, act)
    }
}
