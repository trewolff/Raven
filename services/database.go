package services

import (
	"database/sql"
	"fmt"
	"raven/config"
	"raven/logging"

	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type DatabaseService interface {
	WriteEvent(key string) (string, error)
	GetEvent(key string) ([]string, error)
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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, dbname, sslmode)
	logging.Logger.Debug("psqlInfo=%s", psqlInfo)
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

func (c *DefaultDatabaseService) WriteEvent(eventID string) (string, error) {
	insertStatement := `INSERT INTO events (event_id)
	VALUES($1)
	returning id AS event_id;
	`
	rows, err := c.database.Query(
		insertStatement,
		eventID,
	)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var rowID string
	for rows.Next() {
		if err := rows.Scan(&rowID); err != nil {
			logging.Logger.Error("Error scanning: ", err)
		}
	}
	return rowID, nil
}

func (c *DefaultDatabaseService) GetEvent(eventID string) ([]string, error) {
	insertStatement := `SELECT id FROM events WHERE event_id = ($1)`
	rows, err := c.database.Query(
		insertStatement,
		eventID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var rowIDs []string
	var rowID string
	for rows.Next() {
		if err := rows.Scan(&rowID); err != nil {
			logging.Logger.Error("Error scanning: ", err)
		}
		rowIDs = append(rowIDs, rowID)
	}
	return rowIDs, nil
}
