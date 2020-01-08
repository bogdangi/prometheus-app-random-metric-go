# Introduction

Simple service to produce random metrics to learn how to produce Prometheus metrics in go applications.


# Run tests

`docker run -it -v $PWD:/go/src/app -w /go/src/app -e CGO_ENABLED=0 -e GO111MODULE=on golang go test .`
