package datastructs

import "time"

type OHLCData struct {
	Opening float64   `json:"o"`
	High    float64   `json:"h"`
	Low     float64   `json:"l"`
	Closing float64   `json:"c"`
	Time    time.Time `json:"t"`
}

type VenueResult struct {
	Name         string       `json:"name"`
	Title        string       `json:"title"`
	Mic          string       `json:"mic"`
	IsOpen       bool         `json:"is_open"`
	OpeningHours OpeningHours `json:"opening_hours"`
	OpeningDays  []string     `json:"opening_days"`
}

type OpeningHours struct {
	Start    string `json:"start"`
	End      string `json:"end"`
	Timezone string `json:"timezone"`
}
