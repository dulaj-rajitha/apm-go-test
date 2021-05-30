package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgin"
	"go.elastic.co/apm/transport"
	"log"
	"net/http"
	"os"
)

var glogger = &logrus.Logger{
	Out:   os.Stderr,
	Hooks: make(logrus.LevelHooks),
	Level: logrus.DebugLevel,
	Formatter: &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "log.level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "function.name", // non-ECS
		},
	},
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Reload Environment variables to Elastic Apm.
	if _, err := transport.InitDefault(); err != nil {
		log.Fatal(err)
	}
	print("APM ", os.Getenv("ELASTIC_APM_SERVICE_NAME"))
	r := gin.Default()
	apm.DefaultTracer.SetLogger(glogger)
	r.Use(apmgin.Middleware(r))
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.Run(":8080")
}
