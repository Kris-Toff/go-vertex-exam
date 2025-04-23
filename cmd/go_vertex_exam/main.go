package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DateData struct {
	Day  string
	Date string
}

func main() {
	http.HandleFunc("/", dateHandler)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

func dateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	dateStr := r.FormValue("date")
	layout := "02/01/2006" // matches "day/month/year"

	// Parse the input string
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		http.Error(w, "Error parsing date:", http.StatusBadRequest)
		return
	}

	dateThousandDays := parsedTime.AddDate(0, 0, 1000) // add date by 1000 days

	data := DateData{
		Day:  dateThousandDays.Format("Monday"), // format day of the week
		Date: dateThousandDays.Format(layout),   // format back to DD/MM/YYYY
	}

	json.NewEncoder(w).Encode(data)
}
