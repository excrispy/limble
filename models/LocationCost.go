package models

type LocationCost struct {
	LocationID int     `json:"location_id"`
	Name       string  `json:"name"`
	TotalCost  float64 `json:"total_cost"`
}