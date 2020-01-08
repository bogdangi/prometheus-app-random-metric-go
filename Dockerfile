FROM golang as builder

WORKDIR /go/src/app

COPY main.go main.go
COPY go.mod go.mod
COPY go.sum go.sum

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE=on

RUN go build -a -installsuffix cgo -o /go/bin/app .

FROM scratch
COPY --from=builder /go/bin/app /
CMD ["/app"]
