package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Customer struct {
	ID           string
	Name         string
	Segment      string
	CoBrandCards []string
}

type Transaction struct {
	CustomerID string
	Amount     float64
	MCC        string
	Timestamp  time.Time
}

var customers = []Customer{
	{ID: "C001", Name: "Alice Smith", Segment: "Gold Tier", CoBrandCards: []string{"SuperMart Rewards Card"}},
	{ID: "C002", Name: "Bob Johnson", Segment: "Platinum Tier", CoBrandCards: []string{}},
	{ID: "C003", Name: "Charlie Brown", Segment: "Gold Tier", CoBrandCards: []string{"Airline Miles Plus Card", "SuperMart Rewards Card"}},
	{ID: "C004", Name: "Diana Prince", Segment: "Affluent", CoBrandCards: []string{}},
	{ID: "C005", Name: "Edward Nigma", Segment: "Gold Tier", CoBrandCards: []string{}},
}

var transactions = []Transaction{
	{CustomerID: "C001", Amount: 1500.00, MCC: "5812", Timestamp: time.Now().AddDate(0, -1, 0)},
	{CustomerID: "C001", Amount: 800.00, MCC: "5411", Timestamp: time.Now().AddDate(0, -1, -10)},
	{CustomerID: "C002", Amount: 6000.00, MCC: "5812", Timestamp: time.Now().AddDate(0, -2, 0)},
	{CustomerID: "C002", Amount: 2500.00, MCC: "5812", Timestamp: time.Now().AddDate(0, 0, -15)},
	{CustomerID: "C003", Amount: 300.00, MCC: "5814", Timestamp: time.Now().AddDate(0, -1, -5)},
	{CustomerID: "C004", Amount: 12000.00, MCC: "5812", Timestamp: time.Now().AddDate(0, -3, 0)},
	{CustomerID: "C005", Amount: 4999.00, MCC: "5812", Timestamp: time.Now().AddDate(0, -1, -1)},
}

var diningMCCs = map[string]bool{"5812": true, "5813": true, "5814": true}

func filterCustomers(category string, period string, minSpend string, segmentsCSV string, excludeCSV string) ([]Customer, error) {
	var results []Customer

	months, _ := strconv.Atoi(period)
	if months == 0 {
		months = 3
	}

	spend, convErr := strconv.ParseFloat(minSpend, 64)

	targetSegments := []string{}
	if segmentsCSV != "" {
		targetSegments = strings.Split(segmentsCSV, ",")
	}
	excludeCards := []string{}
	if excludeCSV != "" {
		excludeCards = strings.Split(excludeCSV, ",")
	}

	startDate := time.Now().AddDate(0, -months, 0)

	for _, cust := range customers {
		segmentOk := false
		if len(targetSegments) == 0 || targetSegments[0] == "All" {
			segmentOk = true
		} else {
			for _, s := range targetSegments {
				if cust.Segment == s {
					segmentOk = true
					break
				}
			}
		}
		if !segmentOk {
			continue
		}

		isExcluded := false
		for _, exCard := range excludeCards {
			for _, cCard := range cust.CoBrandCards {
				if cCard == exCard {
					isExcluded = true
					break
				}
			}
			if isExcluded {
				break
			}
		}
		if isExcluded {
			continue
		}

		custSpend := 0.0
		for _, t := range transactions {
			if t.CustomerID == cust.ID && t.Timestamp.After(startDate) {
				if _, isDining := diningMCCs[t.MCC]; isDining && category == "Dining" {
					custSpend += t.Amount
				}
			}
		}

		if custSpend >= spend || convErr != nil {
			results = append(results, cust)
		}
	}

	if len(results) == 0 && convErr != nil {
		return nil, fmt.Errorf("failed to parse minimum spend and no candidates found")
	}
	return results, nil
}

func gourmetCandidatesHandler(w http.ResponseWriter, r *http.Request) {
	cat := r.URL.Query().Get("category")
	period := r.URL.Query().Get("period_months")
	minSpend := r.URL.Query().Get("min_spending")
	segments := r.URL.Query().Get("segments")
	exclude := r.URL.Query().Get("exclude_cards")

	candidates, err := filterCustomers(cat, period, minSpend, segments, exclude)
	if err != nil {
		http.Error(w, "Error finding candidates: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(candidates) == 0 {
		fmt.Fprint(w, "No candidates found matching your criteria.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(candidates)
}

func main() {
	http.HandleFunc("/gourmet-candidates", gourmetCandidatesHandler)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
