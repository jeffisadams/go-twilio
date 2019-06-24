// binFile {{sid}} {{token}} {{from}} {{to}} {{image path}}
package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"./aws"

	twilio "github.com/kevinburke/twilio-go"
)

var sid = os.Args[1]
var token = os.Args[2]
var from = os.Args[3]
var to = os.Args[4]
var path = os.Args[5]

const region = "us-east-1"
const bucket = "test-upload-bucket-frontdoor"

func main() {
	sess, err := aws.Session(region)
	if err != nil {
		log.Fatalln(err)
	}

	uploadURL, uploadErr := aws.Upload(sess, path, bucket, region)
	fmt.Printf("Where to look for the upload: %s", uploadURL)
	if uploadErr != nil {
		log.Fatalln(uploadErr)
	}

	// Start the Twilio upload part
	client := twilio.NewClient(sid, token, nil)

	currentTime := time.Now()
	formattedTime := fmt.Sprintf("formatted time: %s", currentTime.Format("Jan _2 15:04:05"))

	mediaList := make([]*url.URL, 0, 1)
	mediaLink, _ := url.Parse(uploadURL)
	mediaList = append(mediaList, mediaLink)

	mms, err := client.Messages.SendMessage(
		fmt.Sprintf("+%s", from),
		fmt.Sprintf("+%s", to),
		formattedTime,
		mediaList,
	)
	if err != nil {
		fmt.Println("Got here and couldn't do something in sending the message")
		log.Fatal(err)
	}
	fmt.Println(mms)
}
