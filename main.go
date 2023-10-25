package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/backend-services/controllers/dataservice"
	dataservices "github.com/backend-services/dataservices"
	"github.com/backend-services/router"
	"github.com/backend-services/utils/constants"
)

var serviceName = constants.ServiceName

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
	env    = "ENV"
)

func main() {
	fmt.Printf("Starting %v\n", serviceName)
	initializeLogrus()
	initializeDatabase()
	StartServer("9090")
}

func redirect(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		// see @andreiavrammsd comment: often 307 > 301
		http.StatusTemporaryRedirect)
}

func initializeDatabase() {
	log.Info("Initializating Database")
	dataservice.DataService = &dataservices.PostgresClient{}
	dataservice.DataService.Connect()
}

func initializeLogrus() {
	Formatter := new(log.JSONFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	log.SetFormatter(Formatter)

	var filename = constants.ServiceLogFileName
	dir, errWD := os.Getwd()
	if errWD != nil {
		log.Error("Error occurred while getting present working directory, logging to stderror" + " - " + errWD.Error())
	}
	log.Info(dir)
	f, err := os.OpenFile(dir+"/"+filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Error("Error occurred while opening log file, logging to stderror" + " - " + err.Error())
	} else {
		multiWriter := io.MultiWriter(os.Stdout, f)
		log.SetOutput(multiWriter)
	}
}

// StartServer - startup the HTTP server
func StartServer(port string) {
	go http.ListenAndServe(":9091", http.HandlerFunc(redirect))
	go http.ListenAndServe(":8082", router.NewHealthcheckRouter())

	r := router.New()
	http.Handle("/", r)
	log.Info("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Error("An error occurred starting HTTP listener at port " + port + " error: " + err.Error())
	}
}
