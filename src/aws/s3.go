package aws

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Upload Add a file by path to the listed bucket
func Upload(sess *session.Session, path string, bucket string, region string) (string, error) {
	client := s3.New(sess)

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	request := s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(fileInfo.Name()),
		Body:        bytes.NewReader(buffer),
		ContentType: aws.String(http.DetectContentType(buffer)),
	}
	_, err = client.PutObject(&request)

	uploadURL := fmt.Sprintf("http://s3.%s.amazonaws.com/%s/%s", region, bucket, fileInfo.Name())
	return uploadURL, err
}
