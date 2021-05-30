package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"

	"os"

	"github.com/sirupsen/logrus"
	"go.elastic.co/apm"
	_ "go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgorilla"
	"go.elastic.co/apm/transport"
)

var logger = &logrus.Logger{
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

func init() {
	godotenv.Load()
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
	println("APM ", os.Getenv("ELASTIC_APM_SERVICE_NAME"))

	r := mux.NewRouter()

	//tracer, err := apm.NewTracer(os.Getenv("ELASTIC_APM_SERVICE_NAME"), "")
	//tracer.SetCaptureBody(2)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//apmgorilla.Instrument(r, apmgorilla.WithTracer(tracer))

	apmgorilla.Instrument(r)
	apm.DefaultTracer.SetLogger(logger)
	// Ping test
	r.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Hello world!")
	})
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
