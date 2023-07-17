package util

import (
	"context"
	"log"
	"net/smtp"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/pkg/sftp"
)

var (
	ctx    = context.Background()
	snsSvc *sns.SNS
)

type redisSftpRepo struct {
	redisClient *redis.Client
	sftpClient  *sftp.Client
}

func NewRedisSftpSyncRepo(redisClient *redis.Client, sftpClient *sftp.Client) *redisSftpRepo {
	return &redisSftpRepo{
		redisClient: redisClient,
		sftpClient:  sftpClient,
	}
}

func (rdSftp *redisSftpRepo) SubscribeAndUpload(topicName string, folderName string) error {
	pubSub := rdSftp.redisClient.WithContext(ctx).Subscribe(topicName)
	for {
		msg, err := pubSub.ReceiveMessage()
		if err != nil {
			log.Print(err.Error())
			return err
		}

		sftpFile, err := rdSftp.sftpClient.Create("/upload/" + uuid.NewString() + ".txt")
		if err != nil {
			log.Printf("Error 1 %s", err.Error())
			return err
		}

		bytesWritten, err := sftpFile.Write([]byte(msg.Payload))
		if err != nil {
			log.Printf("Error 2 %s", err.Error())
			return err
		}

		log.Printf("Total Bytes written : %d", bytesWritten)
		// err = SendEmail(msg.Payload)
		// if err != nil {
		// 	return err
		// }

		err = PushToSnsTopic(msg.Payload)
		if err != nil {
			return err
		}

		// fileInfo, err := rdSftp.sftpClient.ReadDir("/")
		// if err != nil {
		// 	log.Println(err.Error())
		// 	return err
		// }

		// for _, list := range fileInfo {
		// 	log.Print(list.Name())
		// }

	}
}

func SendEmail(message string) error {
	smtpAuth := smtp.PlainAuth("", os.Getenv("SMTP_FROM"), os.Getenv("SMTP_PASS"), os.Getenv("SMTP_HOST"))

	// Sending Email
	err := smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_HOST"), smtpAuth, os.Getenv("SMTP_FROM"), []string{"jiteshchawla1511@gmail.com"}, []byte(message))
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Email sent successfully")
	return err
}

func PushToSnsTopic(message string) error {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	if err != nil {
		log.Fatal(err)
	}

	snsSvc = sns.New(sess)
	snsInput := sns.PublishInput{
		TopicArn: aws.String(os.Getenv("SNS_ARN")),
		Subject:  aws.String("File uploaded"),
		Message:  aws.String(message),
	}

	_, err = snsSvc.Publish(&snsInput)
	if err != nil {
		log.Print(err)
		return err
	}

	log.Println("Message sent")

	return nil
}
