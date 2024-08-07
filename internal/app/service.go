package app

import (
	"sftp-connect/internal/domain"
	"sftp-connect/internal/infra"
)

type Service struct {
	sftpClient *infra.SFTPClient
}

func NewService(sftpClient *infra.SFTPClient) *Service {
	return &Service{sftpClient: sftpClient}
}

func (u *Service) Upload(file *domain.File) error {
	return u.sftpClient.UploadFile(file, "/bucket-example-transfer/rodrigo/")
}
