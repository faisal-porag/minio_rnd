FROM golang:1.18 as builder

WORKDIR /app
COPY . /app
RUN GIT_TERMINAL_PROMPT=1 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o bin/minio_rnd main.go

# APP
FROM alpine:latest
RUN apk --no-cache add tzdata ca-certificates mailcap && addgroup -S app && adduser -S app -G app
RUN echo "Asia/Dhaka" > /etc/timezone
RUN cp /usr/share/zoneinfo/Asia/Dhaka /etc/localtime

USER app
WORKDIR /app
EXPOSE 8080
COPY --from=builder /app/bin/minio_rnd /usr/local/bin/minio_rnd
CMD ["usr/local/bin/minio_rnd"]
