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
	//_ = CreateNewBucket(minioClient, "poragbucket", "us-east-1")

	// TODO List Of Buckets
	//buckets, err := minioClient.ListBuckets(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, bucket := range buckets {
	//	fmt.Println(bucket)
	//}

	// TODO REMOVE BUCKET BY NAME
	//err = minioClient.RemoveBucket(context.Background(), "poragbucket")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	// TODO FIND BUCKET BY BUCKET NAME
	found, err := minioClient.BucketExists(context.Background(), "poragbucket")
	if err != nil {
		fmt.Println(err)
		return
	}
	if found {
		fmt.Println("Bucket found!")
	} else {
		fmt.Println("Bucket not found!")
	}

}
func CreateNewBucket(minioClient *minio.Client, bucketName, location string) bool {
	// Make a new bucket.
	ctx := context.Background()
	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location, ObjectLocking: true})
	// err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		log.Info().Msgf("bucket exists: %+v", exists)
		if errBucketExists == nil && exists {
			log.Info().Msgf("We already own %+v\n", bucketName)
		} else {
			log.Error().Err(err).Msgf("make bucket error : %s", err)
			return false
		}
	} else {
		log.Info().Msgf("Successfully created %s\n", bucketName)
	}
	return true
}
