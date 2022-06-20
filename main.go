package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

func main() {
	endpoint := "play.min.io"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	//useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: true,
	})
	if err != nil {
		log.Error().Err(err).Msg("minio client initiate error!")
	}

	log.Info().Msg("Successfully initiated.")
	log.Info().Msgf("Result: %#v\n", minioClient) // minioClient is now setup

	// TODO BUCKET CREATION
	// Create a bucket at region 'us-east-1' with object locking enabled.
	//err = minioClient.MakeBucket(context.Background(), "poragbucket", minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: true})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("Successfully created poragbucket.")

	// TODO List Of Buckets
	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}

}
