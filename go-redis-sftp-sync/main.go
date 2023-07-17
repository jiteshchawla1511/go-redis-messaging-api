package main

import (
	"fmt"
	"github/com/jiteshchawla1511/go-redis-sftp-sync/model"
	"github/com/jiteshchawla1511/go-redis-sftp-sync/util"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Panic(err)
	}

}

func main() {
	redisClient := model.CreateRedisConnection()

	if redisClient.Ping().Err() != nil {
		log.Fatal(redisClient.Ping().Err())
	}

	defer redisClient.Close()

	sshConfig := ssh.ClientConfig{
		User:            os.Getenv("SFTP_USER"),
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(os.Getenv("SFTP_PASSWORD")),
		},
	}

	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", os.Getenv("SFTP_HOST"), 22), &sshConfig)
	if err != nil {
		log.Fatal(err)
	}

	defer sshClient.Close()

	sftpClient := model.CreateSftpConnection(sshClient)
	defer sftpClient.Close()

	operations := util.NewRedisSftpSyncRepo(redisClient, sftpClient)
	operations.SubscribeAndUpload("send-user-data", "sftp-new-dir")

}
