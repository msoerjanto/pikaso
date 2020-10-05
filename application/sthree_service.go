package application

import (
	"bytes"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/globalsign/mgo/bson"
)

// UploadMultiPartFileToS3 Uploads a multipart file to S3
func UploadMultiPartFileToS3(file multipart.File, header *multipart.FileHeader, err error, association string) string {
	if file == nil {
		return ""
	}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1"),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Upload
	fileName, err := AddFileToS3(s, file, header, association)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Image uploaded successfully: %v", fileName)
	}
	return fileName
}

// AddFileToS3 uploads file to s3
func AddFileToS3(s *session.Session, file multipart.File, header *multipart.FileHeader, association string) (string, error) {
	// get file size and read file content into buffer
	size := header.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	// create unique file name for the file
	tempFileName := association + "/" + bson.NewObjectId().Hex() + filepath.Ext(header.Filename)

	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String("pikaso-bucket"),
		Key:                  aws.String(tempFileName),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	if err != nil {
		return "", err
	}
	return tempFileName, err
}
