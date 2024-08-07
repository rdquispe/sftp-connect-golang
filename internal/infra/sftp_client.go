package infra

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"sftp-connect/internal/domain"
)

type SFTPClient struct {
	client *sftp.Client
}

func NewSFTPClient(cfg *Config) (*SFTPClient, error) {
	key, err := os.ReadFile(cfg.PrivateKeyPath)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	sshConfig := &ssh.ClientConfig{
		User: cfg.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", cfg.Host+":"+cfg.Port, sshConfig)
	if err != nil {
		return nil, err
	}

	client, err := sftp.NewClient(conn)
	if err != nil {
		return nil, err
	}

	return &SFTPClient{client: client}, nil
}

func (s *SFTPClient) UploadFile(file *domain.File, remotePath string) error {
	dstFile, err := s.client.OpenFile(remotePath+file.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, file.Content)
	return err
}
