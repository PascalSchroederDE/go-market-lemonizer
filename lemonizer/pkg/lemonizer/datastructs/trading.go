package datastructs

import "time"

type PlaceOrderResults struct {
	CreatedAt             time.Time             `json:"created_at"`
	ID                    string                `json:"id"`
	Status                string                `json:"status"`
	RegulatoryInformation RegulatoryInformation `json:"regulatory_information"`
	Isin                  string                `json:"isin"`
	ExpiresAt             time.Time             `json:"expires_at"`
	Side                  string                `json:"side"`
	Quantity              int                   `json:"quantity"`
	StopPrice             interface{}           `json:"stop_price"`
	LimitPrice            interface{}           `json:"limit_price"`
	Venue                 string                `json:"venue"`
	EstimatedPrice        int                   `json:"estimated_price"`
	EstimatedPriceTotal   int                   `json:"estimated_price_total"`
	Notes                 interface{}           `json:"notes"`
	Charge                int                   `json:"charge"`
	ChargeableAt          interface{}           `json:"chargeable_at"`
	KeyCreationID         string                `json:"key_creation_id"`
	Idempotency           interface{}           `json:"idempotency"`
}

type RegulatoryInformation struct {
	CostsEntry                      float64 `json:"costs_entry"`
	CostsEntryPct                   string  `json:"costs_entry_pct"`
	CostsRunning                    float64 `json:"costs_running"`
	CostsRunningPct                 string  `json:"costs_running_pct"`
	CostsProduct                    float64 `json:"costs_product"`
	CostsProductPct                 string  `json:"costs_product_pct"`
	CostsExit                       float64 `json:"costs_exit"`
	CostsExitPct                    string  `json:"costs_exit_pct"`
	YieldReductionYear              int     `json:"yield_reduction_year"`
	YieldReductionYearPct           string  `json:"yield_reduction_year_pct"`
	YieldReductionYearFollowing     int     `json:"yield_reduction_year_following"`
	YieldReductionYearFollowingPct  string  `json:"yield_reduction_year_following_pct"`
	YieldReductionYearExit          int     `json:"yield_reduction_year_exit"`
	YieldReductionYearExitPct       string  `json:"yield_reduction_year_exit_pct"`
	EstimatedHoldingDurationYears   string  `json:"estimated_holding_duration_years"`
	EstimatedYieldReductionTotal    int     `json:"estimated_yield_reduction_total"`
	EstimatedYieldReductionTotalPct string  `json:"estimated_yield_reduction_total_pct"`
	Kiid                            string  `json:"KIID"`
	LegalDisclaimer                 string  `json:"legal_disclaimer"`
}
