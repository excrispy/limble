package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"limble/models"

	_ "github.com/go-sql-driver/mysql"
)

func HandleWorkerCost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	completed := r.URL.Query().Get("completed")
	workerIDsParam := r.URL.Query().Get("worker_ids")

	query := `
		SELECT 
			w.id, 
			w.username, 
			SUM(lt.time_hours * w.hourly_wage) AS total_cost
		FROM test_db.workers AS w
		JOIN test_db.logged_time AS lt ON w.id = lt.worker_id
		JOIN test_db.tasks AS t ON lt.task_id = t.id
		WHERE t.completed = ?
	`

	workerIDs := strings.Split(workerIDsParam, ",")
	args := []interface{}{completed}
	if len(workerIDs) > 0 {
		placeholders := strings.Repeat("?,", len(workerIDs))
		placeholders = placeholders[:len(placeholders)-1] // remove trailing comma
		query += " AND w.id IN (" + placeholders + ")"
		for _, id := range workerIDs {
			args = append(args, id)
		}
	}
	query += " GROUP BY w.id, w.username;"

	rows, err := db.Query(query, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []models.WorkerCost
	for rows.Next() {
		var wc models.WorkerCost
		if err := rows.Scan(&wc.WorkerID, &wc.Username, &wc.TotalCost); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		results = append(results, wc)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func HandleLocationCost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	completed := r.URL.Query().Get("completed")
	locationIDsParam := r.URL.Query().Get("location_id")

	query := `
		SELECT 
			l.id,
			l.name,
			SUM(lt.time_hours * w.hourly_wage) AS total_cost
		FROM locations l
		JOIN tasks t ON l.id = t.location_id
		JOIN logged_time lt ON t.id = lt.task_id
		JOIN workers w ON lt.worker_id = w.id
		WHERE t.completed = ?`

	locationIDs := strings.Split(locationIDsParam, ",")
	args := []interface{}{completed}
	if len(locationIDs) > 0 {
		placeholders := strings.Repeat("?,", len(locationIDs))
		placeholders = placeholders[:len(placeholders)-1] // remove trailing comma
		query += " AND l.id IN (" + placeholders + ")"
		for _, id := range locationIDs {
			args = append(args, id)
		}
	}
	query += " GROUP BY l.id;"

	rows, err := db.Query(query, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []models.LocationCost
	for rows.Next() {
		var lc models.LocationCost
		if err := rows.Scan(&lc.LocationID, &lc.Name, &lc.TotalCost); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		results = append(results, lc)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// will sum cost of all tasks for all locations and return one value
// func HandleTotalLocationCost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
// 	completed := r.URL.Query().Get("completed")
// 	locationIDsParam := r.URL.Query().Get("location_id") // instructions were unclear, allowing multiple location IDs

// 	query := `
// 		SELECT
// 			SUM(lt.time_hours * w.hourly_wage) AS total_cost
// 		FROM locations l
// 		JOIN tasks t ON l.id = t.location_id
// 		JOIN logged_time lt ON t.id = lt.task_id
// 		JOIN workers w ON lt.worker_id = w.id
// 		WHERE t.completed = ?`

// 	locationIDs := strings.Split(locationIDsParam, ",")
// 	args := []interface{}{completed}
// 	if len(locationIDs) > 0 {
// 		placeholders := strings.Repeat("?,", len(locationIDs))
// 		placeholders = placeholders[:len(placeholders)-1] // remove trailing comma
// 		query += " AND l.id IN (" + placeholders + ")"
// 		for _, id := range locationIDs {
// 			args = append(args, id)
// 		}
// 	}

// 	row := db.QueryRow(query, args...)
// 	var totalCost float64
// 	if err := row.Scan(&totalCost); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]float64{"total_cost": totalCost})
// }
