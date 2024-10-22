package services

import (
	"database/sql"
	"fmt"
	"raven/config"

	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type DatabaseService interface {
	WriteEvent(key string) (int8, error)
}

type DefaultDatabaseService struct {
	database *sql.DB
}

func NewDatabaseConnection() *sql.DB {
	var db *sql.DB
	conf, _ := config.GetConfig()
	host := conf.PG_HOST
	port := conf.PG_PORT
	user := conf.PG_USER
	pass := conf.PG_PASS
	dbname := conf.PG_DB_NAME
	sslmode := conf.PG_SSL_MODE
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s pass=%s dbname=%s sslmode=%s", host, port, user, pass, dbname, sslmode)
	db, err := sql.Open(conf.PG_CONNECTOR_TYPE, psqlInfo)
	if err != nil {
		log.Fatal("Error opening database: ", err.Error())
		panic(err)
	}
	db.SetMaxIdleConns(conf.PG_MAX_CONNS)
	err = db.Ping()
	if err != nil {
		log.Error("Ping error: ", err.Error())
		panic(err)
	}
	return db
}

func NewDatabaseService(database *sql.DB) DatabaseService {
	return &DefaultDatabaseService{
		database: database,
	}
}

func (c *DefaultDatabaseService) WriteEvent(entityID string) (int8, error) {
	insertStatement := `INSERT INTO events (entity_id)
	VALUES($1)
	returning id AS event_id;
	`
	rows, err := c.database.Query(
		insertStatement,
		entityID,
	)
	if err != nil {
		return int8(0), err
	}
	defer rows.Close()
	var eventID int8
	for rows.Next() {
		rows.Scan(&eventID)
	}
	return eventID, nil
}
