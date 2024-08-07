package main

import (
	"github.com/gin-gonic/gin"
	"sftp-connect/internal/handler"
	"sftp-connect/internal/infra"
)

func main() {

	cfg := infra.LoadConfig()
	router := gin.Default()
	sftpHandler := handler.NewSftpHandler(cfg)
	router.POST("/upload", sftpHandler.Upload)

	_ = router.Run(":8080")
}
