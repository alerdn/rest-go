package services

import (
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AWSService struct{}

func (as *AWSService) GetSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		log.Println("Erro ao criar sessão AWS:", err)
		return nil, err
	}

	return sess, nil
}

func (as *AWSService) UploadFile(name string, file io.Reader) error {
	sess, err := as.GetSession()
	if err != nil {
		return err
	}

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("agnusgameworks-rest"),
		Key:    aws.String(name),
		Body:   file,
	})
	if err != nil {
		log.Println("Erro ao fazer upload do arquivo:", err)
		return err
	}

	return nil
}

func (as *AWSService) DownloadFile(name string) (*os.File, error) {
	sess, err := as.GetSession()
	if err != nil {
		return nil, err
	}

	file, err := os.Create(name)
	if err != nil {
		log.Println("Erro ao criar arquivo:", err)
		return nil, err
	}

	downloader := s3manager.NewDownloader(sess)
	_, err = downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String("agnusgameworks-rest"),
		Key:    aws.String(name),
	})

	if err != nil {
		log.Println("Erro ao fazer download do arquivo:", err)
		return nil, err
	}

	return file, nil
}

func (as *AWSService) DeleteFile(name string) error {
	sess, err := as.GetSession()
	if err != nil {
		return err
	}

	svc := s3.New(sess)
	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String("agnusgameworks-rest"),
		Key:    aws.String(name),
	})
	if err != nil {
		log.Println("Erro ao deletar o arquivo:", err)
		return err
	}

	return nil
}