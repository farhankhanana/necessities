package sftp

import (
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/gat/necessities/logger"
	"github.com/gat/necessities/ssh"
	"github.com/pkg/sftp"
)

type ClientSFTP struct {
	Client *sftp.Client
}

// NewClientSFTP creates new client object SFTP.
func NewClientSFTP(sshClient *ssh.ClientSSH) (*ClientSFTP, error) {
	logger := logger.NewLogger("")
	sftpClient, err := sftp.NewClient(sshClient.Client)
	if err != nil {
		logger.LogError("failed create new sftp client", err)
		return nil, err
	}

	c := &ClientSFTP{
		Client: sftpClient,
	}

	return c, err
}

// Upload uploads file to designated path and folder.
func (c *ClientSFTP) Upload(directoryPath, folderName, fileName string, file multipart.File) error {
	logger := logger.NewLogger("")
	directoryPath = strings.TrimPrefix(directoryPath, ".")
	fullPath := fmt.Sprintf(".%s/%s/%s", directoryPath, folderName, fileName)

	fDestination, err := c.Client.Create(fullPath)
	if err != nil {
		logger.LogError("failed to create destination file", err)
		return err
	}

	_, err = io.Copy(fDestination, file)
	if err != nil {
		logger.LogError("failed copy source file into destination file", err)
		return err
	}

	return err
}
