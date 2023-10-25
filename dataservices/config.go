package dataservices

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/backend-services/config"
)

const (
	dbhost string = "DBHOST"
	dbport string = "DBPORT"
	dbuser string = "DBUSER"
	dbpass string = "DBPASS"
	dbname string = "DBNAME"
	env    string = "ENV"
)

func dbConfig() map[string]string {
	var host, port, user, password, name string
	conf := make(map[string]string)
	env, ok := os.LookupEnv(env)
	if !ok {
		log.Info("Setting env to localhost...")
		env = "localhost"
	}
	log.Info("Applying " + env + " DB configuration")
	if env == "localhost" {
		config.SetEnvFromDotEnv()
	}
	host, ok = os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	port, ok = os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}
	user, ok = os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	password, ok = os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}
	name, ok = os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}

	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}
