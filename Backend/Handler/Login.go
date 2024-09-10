package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	model "social-network/Model"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	fmt.Println(string(body))

	var data model.Auth
	json.Unmarshal(body, &data)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
