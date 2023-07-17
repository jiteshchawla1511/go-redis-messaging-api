package model

import (
	"log"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func CreateSftpConnection(sshClient *ssh.Client) *sftp.Client {

	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		log.Fatal(err)
	}

	return sftpClient

}
