package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/s1ntaxe770r/gopher-quotes/pkg/models"
	"github.com/sirupsen/logrus"
)

type StatsResponse struct {
	QuoteCount int          `json:"quote_count"`
	LastQuote  models.Quote `json:"last_quote"`
}

func (qh *QuotesHandler) GetStats(rw http.ResponseWriter, r *http.Request) {
	Quotes, err := qh.db.Get()
	if err != nil {
		qh.l.WithFields(logrus.Fields{"msg": "unable to retrieve quotes", "reason": err.Error()}).Error()
	}
	// get quote count by calculating the length of the slice returned
	QC := len(Quotes)
	// get the last quote in the slice
	LQ := qh.db.GetLastQuote()
	qoute := models.Quote{}
	_ = json.Unmarshal(LQ.LastQuote, &qoute)
	response := StatsResponse{QC, qoute}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response)

}
