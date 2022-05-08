package db

import (
	"github.com/s1ntaxe770r/gopher-quotes/pkg/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB : database connection object
type DB struct {
	conn *gorm.DB
}

// NewDB returns a new instance of DB
func NewDB() *DB {
	conn, err := gorm.Open(sqlite.Open("quotes.db"), &gorm.Config{})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg":    "unable to connect to database",
			"reason": err.Error(),
		}).Error("error connecting to database")
	}
	err = conn.AutoMigrate(models.Quote{}, models.QuoteData{})
	if err != nil {
		logrus.Infof("unable to migrate models, reason: %s", err.Error())
	}
	return &DB{conn}
}

// Create inserts a new quote
func (db DB) Create(quote models.Quote) error {
	result := db.conn.Create(&quote)
	return result.Error
}

// Get returns all quotes as slice
func (db DB) Get() ([]models.Quote, error) {
	var quotes []models.Quote
	result := db.conn.Find(&quotes)
	return quotes, result.Error
}

func (db DB) UpdateStats(lq models.QuoteData) {
	res := db.conn.First(&models.QuoteData{})
	if res.RowsAffected == 1 {
		err := db.conn.Model(&lq).Update("last_quote", lq.LastQuote).Error
		if err != nil {
			logrus.Errorf("failed to update quote data, %s", err.Error())
		}
		return
	}
	err := db.conn.Create(&lq).Error
	if err != nil {
		logrus.Errorf("failed to insert quote data,  %s", err.Error())
	}

}

func (db DB) GetLastQuote() models.QuoteData {
	var qd models.QuoteData
	err := db.conn.First(&qd).Error
	if err != nil {
		logrus.Errorf("failed to retrieve lastq quote, reason: %s", err.Error())
	}
	return qd
}
