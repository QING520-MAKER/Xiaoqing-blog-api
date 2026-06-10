package main

// @title Xiaoqing Blog API
// @version 1.0
// @description Xiaoqing Blog 后端 API（Admin + Public V1）。
// @BasePath /api
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @securityDefinitions.apikey AdminSession
// @in cookie
// @name fiber_session

import (
	"log"

	"github.com/QING520-MAKER/Xiaoqing-blog-api/config"
	_ "github.com/QING520-MAKER/Xiaoqing-blog-api/docs"
	"github.com/QING520-MAKER/Xiaoqing-blog-api/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(cfg)
}
