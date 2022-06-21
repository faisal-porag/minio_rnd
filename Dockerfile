# Create build stage based on image
FROM golang:1.18 as builder

# Create working directory under /app
WORKDIR /app
# Copy over all go config (go.mod, go.sum etc.)
COPY . /app

# Install any required modules
RUN go mod download

# Run the Go build and output binary
RUN GIT_TERMINAL_PROMPT=1 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o bin/minio_rnd main.go

# APP
FROM alpine:latest
RUN apk --no-cache add tzdata ca-certificates mailcap && addgroup -S app && adduser -S app -G app
RUN echo "Asia/Dhaka" > /etc/timezone
RUN cp /usr/share/zoneinfo/Asia/Dhaka /etc/localtime


WORKDIR /app
# Make sure to expose the port the HTTP server is using
EXPOSE 8080
# Copy over the binary built from the previous stage
COPY --from=builder /app/bin/minio_rnd /usr/local/bin/minio_rnd
# Run the app binary when we run the container
CMD ["usr/local/bin/minio_rnd"]
