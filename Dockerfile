FROM golang:1.11-alpine as builder
WORKDIR /go/src/github.com/gnur/golpje/
COPY vendor vendor
COPY search search
COPY config config
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine
RUN apk add --no-cache ca-certificates ffmpeg
WORKDIR /
COPY --from=builder /go/src/github.com/gnur/golpje/app .
CMD ["/app"]
