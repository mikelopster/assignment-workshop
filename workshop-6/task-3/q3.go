package main

import (
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

func findCandidatesLast3Months(w http.ResponseWriter, r *http.Request) {
	minSpendingStr := r.URL.Query().Get("min_spending")
	targetSegmentsStr := r.URL.Query().Get("segments")
	excludeCardsStr := r.URL.Query().Get("exclude_cards")

	minSpending, _ := strconv.ParseFloat(minSpendingStr, 64)
	targetSegments := strings.Split(targetSegmentsStr, ",")
	excludeCards := strings.Split(excludeCardsStr, ",")
	if targetSegmentsStr == "" {
		targetSegments = []string{}
	}
	if excludeCardsStr == "" {
		excludeCards = []string{}
	}

	var candidates []Customer
	startDate := time.Now().AddDate(0, -3, 0)

	for _, cust := range customers {
		segmentOk := false
		if len(targetSegments) == 0 || (len(targetSegments) == 1 && targetSegments[0] == "All Segments") {
			segmentOk = true
		} else {
			for _, seg := range targetSegments {
				if cust.Segment == seg {
					segmentOk = true
					break
				}
			}
		}
		if !segmentOk {
			continue
		}

		excludedByCard := false
		if len(excludeCards) > 0 {
			for _, excludedCard := range excludeCards {
				for _, heldCard := range cust.CoBrandCards {
					if heldCard == excludedCard {
						excludedByCard = true
						break
					}
				}
				if excludedByCard {
					break
				}
			}
		}
		if excludedByCard {
			continue
		}

		totalSpendingInCategory := 0.0
		for _, trans := range transactions {
			if trans.CustomerID == cust.ID && trans.Timestamp.After(startDate) {
				if _, isDining := diningMCCs[trans.MCC]; isDining {
					totalSpendingInCategory += trans.Amount
				}
			}
		}

		if totalSpendingInCategory >= minSpending {
			candidates = append(candidates, cust)
		}
	}

	fmt.Fprintf(w, "Found %d candidates (Last 3 Months):\n", len(candidates))
	for _, c := range candidates {
		fmt.Fprintf(w, "- ID: %s, Name: %s\n", c.ID, c.Name)
	}
}

func findCandidatesLast6Months(w http.ResponseWriter, r *http.Request) {
	minSpendingStr := r.URL.Query().Get("min_spending")
	targetSegmentsStr := r.URL.Query().Get("segments")
	excludeCardsStr := r.URL.Query().Get("exclude_cards")

	minSpending, _ := strconv.ParseFloat(minSpendingStr, 64)
	targetSegments := strings.Split(targetSegmentsStr, ",")
	excludeCards := strings.Split(excludeCardsStr, ",")
	if targetSegmentsStr == "" {
		targetSegments = []string{}
	}
	if excludeCardsStr == "" {
		excludeCards = []string{}
	}

	var candidates []Customer
	startDate := time.Now().AddDate(0, -6, 0)

	for _, cust := range customers {
		segmentOk := false
		if len(targetSegments) == 0 || (len(targetSegments) == 1 && targetSegments[0] == "All Segments") {
			segmentOk = true
		} else {
			for _, seg := range targetSegments {
				if cust.Segment == seg {
					segmentOk = true
					break
				}
			}
		}
		if !segmentOk {
			continue
		}

		excludedByCard := false
		if len(excludeCards) > 0 {
			for _, excludedCard := range excludeCards {
				for _, heldCard := range cust.CoBrandCards {
					if heldCard == excludedCard {
						excludedByCard = true
						break
					}
				}
				if excludedByCard {
					break
				}
			}
		}
		if excludedByCard {
			continue
		}

		totalSpendingInCategory := 0.0
		for _, trans := range transactions {
			if trans.CustomerID == cust.ID && trans.Timestamp.After(startDate) {
				if _, isDining := diningMCCs[trans.MCC]; isDining {
					totalSpendingInCategory += trans.Amount
				}
			}
		}

		if totalSpendingInCategory >= minSpending {
			candidates = append(candidates, cust)
		}
	}

	fmt.Fprintf(w, "Found %d candidates (Last 6 Months):\n", len(candidates))
	for _, c := range candidates {
		fmt.Fprintf(w, "- ID: %s, Name: %s\n", c.ID, c.Name)
	}
}

func main() {
	http.HandleFunc("/find-candidates-3m", findCandidatesLast3Months)
	http.HandleFunc("/find-candidates-6m", findCandidatesLast6Months)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
