package main

import (
	"github.com/Chanokthorn/blog-samples-efficient-report-generation/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()

	reportGenerator := internal.NewReportGenerator()
	apiHandler := internal.NewAPIHandler(reportGenerator)

	server.GET("/", apiHandler.GetReport)

	server.Run(":3000")
}
