package main

import (
	"fmt"
	"github.com/rs/zerolog/log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	// Initialize minio client object.
	s3Client, err := minio.New("s3.amazonaws.com", &minio.Options{
		Creds:  credentials.NewStaticV4("YOUR-ACCESSKEYID", "YOUR-SECRETACCESSKEY", ""),
		Secure: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Info().Msg("AWS S3 initialize success.")
	log.Info().Msgf("%v", s3Client)
}
