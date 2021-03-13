package models

import "github.com/shopspring/decimal"

type Statistic struct {
	Date   string          `json:"date"`
	Views  int             `json:"views"`
	Clicks int             `json:"clicks"`
	Cost   decimal.Decimal `json:"cost"`
	CPC    int             `json:"cpc"`
	CPM    int             `json:"cpm"`
}
