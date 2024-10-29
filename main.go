package main

import (
	"encoding/json"
	"fmt"
	ai_orchestration_service "go-app/ai-orchestration-service"
	"log"
	"net/http"
	"time"
)

var cache = ai_orchestration_service.NewCache()

func modelHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var reqBody ai_orchestration_service.RequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	results := []ai_orchestration_service.Response{}

	for _, choice := range reqBody.Choices {
		cacheKey := fmt.Sprintf("%s-%s-%s", choice.Model, choice.Version, choice.Region)
		url, found := cache.Get(cacheKey)
		if !found {
			log.Printf("Cache miss for key: %s", cacheKey)
			results = append(results, ai_orchestration_service.Response{
				Model:    choice.Model,
				Version:  choice.Version,
				Region:   choice.Region,
				Response: "error: cache miss",
				Status:   "failed",
			})
			continue
		}

		log.Printf("Cache hit for key: %s, URL: %s", cacheKey, url)

		body, err := json.Marshal(ai_orchestration_service.RequestBody{
			Model:      reqBody.Model,
			Version:    reqBody.Version,
			Timeout:    reqBody.Timeout,
			Choices:    []ai_orchestration_service.Choice{choice},
			CallOpenAI: reqBody.CallOpenAI,
		})
		if err != nil {
			http.Error(w, "Error marshalling request body", http.StatusInternalServerError)
			return
		}

		respBody, err := ai_orchestration_service.CallAzureOpenAI(url, body, time.Duration(reqBody.Timeout)*time.Second)
		status := "failed"
		if err == nil && respBody != "" {
			status = "success"
		} else {
			log.Printf("Error calling Azure OpenAI: %v", err)
		}

		results = append(results, ai_orchestration_service.Response{
			Model:    choice.Model,
			Version:  choice.Version,
			Region:   choice.Region,
			Response: respBody,
			Status:   status,
		})

		if status == "success" {
			break
		}
	}

	responseData, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error marshalling response data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}

func main() {
	http.HandleFunc("/model", modelHandler)
	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
