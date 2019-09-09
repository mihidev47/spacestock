package datasource

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Storage struct {
	Name       string `yaml:"name"`
	AccessKey  string `yaml:"access_key"`
	SecretKey  string `yaml:"secret_key"`
	Region     string `yaml:"region"`
	BucketName string `yaml:"bucket_name"`
	RootFolder string `yaml:"root_folder"`
	*session.Session
}

// Init implement FileUploader on FileStorage datasource
func (s *S3Storage) Init() FileUploader {
	// Create static credentials
	c := credentials.NewStaticCredentials(s.AccessKey, s.SecretKey, "")
	_, err := c.Get()
	if err != nil {
		log.Errorf("Unable to get aws credential. Error: %s", err.Error())
		os.Exit(19)
	}
	// Setup configuration
	config := aws.NewConfig().WithRegion(s.Region).WithCredentials(c)
	s3Session, err := session.NewSession(config)
	if err != nil {
		log.Errorf("Unable to initiate aws session. Error: %s", err.Error())
		os.Exit(20)
	}
	s.Session = s3Session
	return s
}

// Upload store file from multipart/form request into filesystem
func (s *S3Storage) Upload(folder string, file multipart.File, fileName string) error {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error(err)
		return err
	}
	// filename
	p := s.GetFolderPath(folder) + fileName
	// Push to aws
	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(s.BucketName),
		Key:                  aws.String(p),
		Body:                 bytes.NewReader(data),
		ContentType:          aws.String(http.DetectContentType(data)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		ACL:                  aws.String("public-read"),
	})
	if err != nil {
		log.Error(err)
		return err
	}
	// Success
	return nil
}

// UploadBulk handles multiple file uploads
func (s *S3Storage) UploadBulk(params ...UploadParam) error {
	// Get uploader
	uploader := s3manager.NewUploader(s)
	// Prepare object iterator
	iterator, err := s.newUploadObjectsIterator(params)
	if err != nil {
		return err
	}
	// Upload
	err = uploader.UploadWithIterator(aws.BackgroundContext(), iterator)
	if err != nil {
		return err
	}
	return nil
}

// GetFolderPath returns folder path in filesystem
func (s *S3Storage) GetFolderPath(folder string) string {
	return s.RootFolder + folder + "/"
}

// newUploadObjectsIterator takes arrays of UploadParam, converts to BatchUploadObject and returns iterator
func (s *S3Storage) newUploadObjectsIterator(files []UploadParam) (*s3manager.UploadObjectsIterator, error) {
	// Init result
	var result []s3manager.BatchUploadObject
	// Iterate files
	for _, p := range files {
		// Read data
		data, err := ioutil.ReadAll(p.File)
		if err != nil {
			return nil, err
		}
		// Key
		key := s.GetFolderPath(p.Folder) + p.FileName
		// Create new BatchUploadObject
		o := s3manager.BatchUploadObject{
			Object: &s3manager.UploadInput{
				Bucket:               aws.String(s.BucketName),
				Key:                  aws.String(key),
				Body:                 bytes.NewReader(data),
				ContentType:          aws.String(http.DetectContentType(data)),
				ContentDisposition:   aws.String("attachment"),
				ServerSideEncryption: aws.String("AES256"),
				ACL:                  aws.String("public-read"),
			},
		}
		// Push to result
		result = append(result, o)
	}
	return &s3manager.UploadObjectsIterator{Objects: result}, nil
}
