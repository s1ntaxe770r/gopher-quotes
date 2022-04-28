package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/s1ntaxe770r/gopher-quotes/pkg/db"
	"github.com/s1ntaxe770r/gopher-quotes/pkg/models"
	"github.com/sirupsen/logrus"
)

type QuotesHandler struct {
	db *db.DB
	l  *logrus.Logger
}

// NewQoutesHandler returns a new quotes handler
func NewQoutesHandler(db *db.DB) *QuotesHandler {
	l := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
	return &QuotesHandler{db, l}
}

func (qh *QuotesHandler) CreateQuote(rw http.ResponseWriter, r *http.Request) {
	qh.l.Info("create handler called")
	quote := models.Quote{}
	err := json.NewDecoder(r.Body).Decode(&quote)
	if err != nil {
		qh.l.Error("unable to process request")
		http.Error(rw, "unable to process quote", http.StatusBadRequest)
		return
	}
	err = qh.db.Create(quote)
	if err != nil {
		qh.l.WithFields(logrus.Fields{"msg": "unable to create quote", "reason": err.Error()}).Error()
		http.Error(rw, "unable to create quote", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(quote)
}

func (qh *QuotesHandler) GetQuotes(rw http.ResponseWriter, r *http.Request) {
	qh.l.Info("get handler called")
	quotes, err := qh.db.Get()
	if err != nil {
		qh.l.WithFields(logrus.Fields{"msg": "unable to get quote", "reason": err.Error()}).Error()
		http.Error(rw, "unable to retrieve quotes", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(quotes)
}
