package db

import (
	"github.com/s1ntaxe770r/gopher-quotes/pkg/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB : database connection object
type DB struct {
	conn *gorm.DB
}

// NewDB returns a new instance of DB
func NewDB(dburi string) *DB {
	conn, err := gorm.Open(postgres.Open(dburi),&gorm.Config{})
	if err != nil {
		logrus.Panic(logrus.WithFields(logrus.Fields{
			"msg":"unable to connect to database",
			"reason": err.Error(),
		}))
	}
	err = conn.AutoMigrate(models.Quote{})
	if err != nil {
		logrus.Infof("unable to migrate models, reason: %s",err.Error())
	}
	return &DB{conn}
}
// Create inserts a new quote
func ( db DB) Create(quote models.Quote) error {
	result := db.conn.Create(&quote) 
	return result.Error
}

// Get returns all quotes as slice
func ( db DB ) Get()([]models.Quote,error) {
     var quotes []models.Quote
     result := db.conn.Find(&quotes) 
     return quotes, result.Error
}
