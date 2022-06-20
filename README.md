# Initialize MinIO Client object instructions

 The MinIO Go Client SDK provides simple APIs to access any Amazon S3 compatible object storage.
This quickstart guide will show you how to install the MinIO client SDK, 
connect to MinIO, and provide a walk-through for a simple file uploader. 
For a complete list of APIs and examples, please take a look at the 
[Go Client API Reference](https://docs.min.io/docs/golang-client-api-reference).

This document assumes that you have a working [Go development environment](https://go.dev/doc/install).

[Official Document Resource](https://docs.min.io/docs/golang-client-api-reference.html)

### Download from GitHub
```shell
go get github.com/minio/minio-go/v7
```

### Initialize MinIO Client
MinIO client requires the following four parameters specified to connect to an Amazon S3 compatible object storage.

```go
package main

import (
    "log"

    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
    endpoint := "play.min.io"
    accessKeyID := "Q3AM3UQ867SPQQA43P2F"
    secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
    useSSL := true

    // Initialize minio client object.
    minioClient, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
        Secure: useSSL,
    })
    if err != nil {
        log.Fatalln(err)
    }

    log.Printf("%#v\n", minioClient) // minioClient is now set up
}
```

### API Reference : Bucket Operations
* MakeBucket
* ListBuckets
* BucketExists
* RemoveBucket
* ListObjects
* ListIncompleteUploads

FPutObject(ctx context.Context, bucketName, objectName, filePath, opts PutObjectOptions) (info UploadInfo, err error)
Uploads contents from a file to objectName.

FPutObject uploads objects that are less than 128MiB in a single PUT operation. For objects that are greater than the 128MiB in size, FPutObject seamlessly uploads the object in chunks of 128MiB or more depending on the actual file size. The max upload size for an object is 5TB.

#### Parameters
Param | Type             | Description
------ |------------------| ---------
ctx | context.Context  | Custom context for timeout/cancellation of the call
bucketName | string | Name of the bucket
objectName | string | Name of the object
filePath | string | Path to file to be uploaded
opts | minio.PutObjectOptions | Pointer to struct that allows user to set optional custom metadata, content-type, content-encoding, content-disposition, content-language and cache-control headers, pass encryption module for encrypting objects, and optionally configure number of threads for multipart put operation.

#### minio.UploadInfo
Field | Type             | Description
------ |------------------| ---------
info.ETag | string | The ETag of the new object
info.VersionID | string | The version identifyer of the new object


#### Example
```go
uploadInfo, err := minioClient.FPutObject(context.Background(), "my-bucketname", "my-objectname", "my-filename.csv", minio.PutObjectOptions{
    ContentType: "application/csv",
});
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println("Successfully uploaded object: ", uploadInfo)
```