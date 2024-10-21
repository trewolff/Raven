package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Specification struct {
	SERVER_HOST_PORT string       `default:"localhost:8080"`
	SERVER_URL       string       `default:"ws://localhost:8080/socket"`
	PG_HOST          string       `default:"localhost"`
	PG_PORT          int          `default:"5432"`
	PG_USER          string       `default:"postgres"`
	PG_PASS          string       `default:"password"`
	PG_DB_NAME       string       `default:"postgres"`
	PG_SSL_MODE      string       `default:"require"`
	PG_MAX_CONNS     int          `default:"100"`
	REDIS_URL        string       `default:"localhost:236379"`
	REDIS_PASSWORD   string       `default:""`
	LOG_LEVEL        logrus.Level `default:"debug" logrus:"logrus.ParseLevel()"`
}

func GetConfig() (Specification, error) {
	var s Specification
	err := envconfig.Process("raven", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	return s, nil
}
