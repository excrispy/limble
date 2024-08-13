package models

type WorkerCost struct {
	WorkerID  int     `json:"worker_id"`
	Username  string  `json:"username"`
	TotalCost float64 `json:"total_cost"`
}