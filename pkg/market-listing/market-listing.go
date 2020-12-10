package market_listing

import (
	"github.com/jackc/pgtype"
	"github.com/scotthaley/globofactory/internal/database"
)

type ListingType string

const (
	Buy ListingType = "buy"
	Sell = "sell"
)

type MarketListing struct {
	Code        string
	OrderType   ListingType
	Price       float64
	Amount      int64
	ListingDate pgtype.Timestamp
	ExpiryDate  pgtype.Timestamp
}

func Create(listing MarketListing) {
	database.DBCon.Create(&listing)
}