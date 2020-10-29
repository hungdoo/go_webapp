package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type requestData struct {
	Target int `json:"target"`
}

type responseMessage struct {
	Prime   int    `json:"prime"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// Reference: https://en.wikipedia.org/wiki/Primality_test#Python_Code
func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func preLargestPrimeLookUp(target int) int {
	for i := target; i > 1; i-- {
		if isPrime(i) == true {
			return i
		}
	}
	return -1
}

func PrimeHandler(w http.ResponseWriter, r *http.Request) {

	var req requestData
	res := responseMessage{0, "", http.StatusBadRequest}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	// Look up desire Prime
	prime := preLargestPrimeLookUp(req.Target)
	if prime > 1 {
		res.Status = http.StatusOK
		res.Message = "Found prime"
		res.Prime = prime
	} else {
		res.Status = http.StatusBadRequest
		res.Message = "Please enter interger larger than 1"
		res.Prime = 0
	}
	log.Println(res)
	json.NewEncoder(w).Encode(res)

}
