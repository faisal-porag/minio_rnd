minio_run:
	go run main.go

image_convert:
	go run image_convert/main.go

encode_base64:
	go run image_convert/encode_base64.go

encode_v1:
	go run image_convert/encode_version1.go

build:
	go build main.go

clean:
	rm -r bin


