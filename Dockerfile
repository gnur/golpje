FROM golang:1.8.1 as builder
WORKDIR /go/src/github.com/gnur/golpje/
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
WORKDIR /
COPY --from=builder /go/src/github.com/gnur/snost/app .
EXPOSE 8080
CMD ["/app"]
