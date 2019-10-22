package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		req := make(map[string]interface{})
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		respText := "hello"
		reqText := req["request"].(map[string]interface{})["original_utterance"]
		if reqText != "" {
			respText = reqText.(string)
		}

		resp := map[string]interface{}{
			"response": map[string]interface{}{
				"text": respText,
			},
			"session": req["session"],
			"version": req["version"],
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":5000", nil)
}
