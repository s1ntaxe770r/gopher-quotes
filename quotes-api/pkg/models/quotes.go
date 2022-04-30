package models

// Quote
type Quote struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type QuoteData struct {
	Id        int    `json:"-"`
	LastQuote []byte `json:"last_quote"`
}
