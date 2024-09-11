package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	model "social-network/Model"
)

func Register(w http.ResponseWriter, r *http.Request) {
	contextValue := r.Context().Value(model.RegisterCtx).([]byte)

	var register model.Register
	if err := json.Unmarshal(contextValue, &register); err != nil {
		fmt.Println(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Internal Error: go see the terminal of the server")
		return
	}

	// INSERT in BDD in the Auth table and in the UserInfo table

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(register)
}
