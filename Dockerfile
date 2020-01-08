FROM golang:alpine as builder
WORKDIR /go/src/app
COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/app .

FROM scratch
COPY --from=builder /go/bin/app /
CMD ["/app"]
