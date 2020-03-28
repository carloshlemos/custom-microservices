/*
 * @File: main.go
 * @Description: Api de Teste com Golang + Gin Web(https://github.com/gin-gonic/gin)
 * @Author: Carlos Henrique Lemos (chenriquelemos@gmail.com)
 */
package main

import (
	"io"
	"os"

	"./common"
	"./controllers"
	"./databases"
	"github.com/gin-gonic/gin"

	_ "./docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error
	// Load config file
	err = common.LoadConfig()
	if err != nil {
		return err
	}

	// Initialize User database
	databases.Database.Connect()
	if err != nil {
		return err
	}

	// Setting Gin Logger
	if common.Config.EnableGinFileLog {
		f, _ := os.Create("logs/gin.log")
		if common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
		} else {
			gin.DefaultWriter = io.MultiWriter(f)
		}
	} else {
		if !common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter()
		}
	}

	m.router = gin.Default()

	return nil
}

// @title Experimento com Golang pra MicroServices
// @version 1.0
// @description Api de Teste com Golang + Gin Web(https://github.com/gin-gonic/gin)
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /api/v1
func main() {
	m := Main{}

	// Initialize server
	if m.initServer() != nil {
		return
	}

	defer databases.Database.Close()

	c := controllers.Custom{}
	// Simple group: v1
	v1 := m.router.Group("/api/v1")
	{
		test := v1.Group("/test")
		{
			test.GET("/hello", c.Test)
		}
	}

	m.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	m.router.Run(common.Config.Port)
}
