package models

import "github.com/shopspring/decimal"

type Statistic struct {
	Date   string          `json:"date"`
	Views  int             `json:"views"`
	Clicks int             `json:"clicks"`
	Cost   decimal.Decimal `json:"cost"`
	CPC    decimal.Decimal `json:"cpc"`
	CPM    decimal.Decimal `json:"cpm"`
}
