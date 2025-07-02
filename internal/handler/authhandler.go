package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (*Handler) hand_test(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	data := make(map[string]string)
	data["status"] = "ok"

	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Printf("Error %v", err)
	}

}
