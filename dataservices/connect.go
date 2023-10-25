package dataservices

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/lib/pq" // blank import required for postgres database driver and connection
	log "github.com/sirupsen/logrus"
)

// Connect opens a connection to the DB
func (pc *PostgresClient) Connect() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])
	log.Info(psqlInfo)
	pc.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Connecting to "+strings.ToUpper(config[dbname])+" DB. Fatal error encountered: ", err.Error())
	}
	err = pc.DB.Ping()
	if err != nil {
		log.Fatal("Pinging " + strings.ToUpper(config[dbname]) + " DB. Unable to ping database. Shutting down server.")
	}
	pc.DB.SetMaxOpenConns(20)
	pc.DB.SetMaxIdleConns(5)
	pc.DB.SetConnMaxLifetime(5 * time.Minute)
	log.Info("Successfully connected to " + strings.ToUpper(config[dbname]) + " database.")
}
