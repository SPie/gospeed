package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectionHandler is responsible for database connections
type ConnectionHandler interface {
	AutoMigrate(model interface{})
	Create(model interface{})
	Find(model interface{}) ConnectionHandler
	Close() error
}

type connectionHandler struct {
	db *gorm.DB
}

// NewConnectionHandler creates a new ConnectionHandler with an open connection
func NewConnectionHandler(username string, password string, host string, port string, database string) (ConnectionHandler, error) {
	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		return &connectionHandler{}, err
	}

	return &connectionHandler{db: db}, nil
}

func (ch connectionHandler) AutoMigrate(model interface{}) {
	ch.db.AutoMigrate(model)
}

func (ch connectionHandler) Create(model interface{}) {
	ch.db.Create(model)
}

func (ch connectionHandler) Find(model interface{}) ConnectionHandler {
	db := ch.db.Find(model)

	return connectionHandler{db: db}
}

func (ch connectionHandler) Close() error {
	return nil
}
