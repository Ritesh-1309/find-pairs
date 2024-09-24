package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Numbers []int `json:"numbers"`
	Target  int   `json:"target"`
}

type Response struct {
	Solutions [][]int `json:"solutions"`
}

func findPairs(numbers []int, target int) [][]int {
	indexMap := make(map[int]int)
	var solutions [][]int

	for i, num := range numbers {
		rem := target - num
		if j, found := indexMap[rem]; found {
			solutions = append(solutions, []int{j, i})
		}
		indexMap[num] = i
	}

	fmt.Println("solution :=", solutions)

	return solutions
}

func findPairsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Log that a request was received
	fmt.Println("Received request")

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	fmt.Printf("Decoded request: %+v\n", req)

	// Validate input
	if len(req.Numbers) == 0 {
		json.NewEncoder(w).Encode(Response{Solutions: [][]int{}})
		return
	}

	solutions := findPairs(req.Numbers, req.Target)
	response := Response{Solutions: solutions}

	// Log the response
	fmt.Printf("Sending response: %+v\n", response)

	json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("Server started")
	http.HandleFunc("/find-pairs", findPairsHandler)

	// Run the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
