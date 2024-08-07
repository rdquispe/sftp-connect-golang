package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sftp-connect/internal/app"
	"sftp-connect/internal/domain"
	"sftp-connect/internal/infra"
)

type SftpHandler struct {
	service *app.Service
}

func NewSftpHandler(cfg *infra.Config) *SftpHandler {
	client, _ := infra.NewSFTPClient(cfg)
	handler := app.NewService(client)
	return &SftpHandler{
		service: handler,
	}
}

func (s *SftpHandler) Upload(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file: " + err.Error()})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open file: " + err.Error()})
		return
	}
	defer file.Close()

	domainFile := &domain.File{
		Name:    fileHeader.Filename,
		Content: file,
	}

	if err := s.service.Upload(domainFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to SFTP server: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File successfully uploaded to SFTP server"})
}
