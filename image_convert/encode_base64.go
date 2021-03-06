package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"os"
)

func main() {
	// Open file on disk.
	f, err := os.Open("../flower.jpg")

	if err != nil {
		log.Error().Err(err).Msgf("file open error!")
	}

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	// Print encoded data to console.
	// ... The base64 image can be used as a data URI in a browser.
	fmt.Println("ENCODED: " + encoded)
}
