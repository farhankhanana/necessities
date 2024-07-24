package ssh

import (
	"fmt"

	"github.com/gat/necessities/logger"
	"golang.org/x/crypto/ssh"
)

type ClientSSH struct {
	Client *ssh.Client
}

// NewClientSSHWithPassword sets and creates configuration for SSH client.
func NewClientSSHWithPassword(host, port, username, password string) (*ClientSSH, error) {
	logger := logger.NewLogger("")
	sshConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}

	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", host, port), sshConfig)
	if err != nil {
		defer sshClient.Close()
		logger.LogError("ssh client dial", err)
		return nil, err
	}

	return &ClientSSH{
		Client: sshClient,
	}, err
}
