package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	model "social-network/Model"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		var register model.Register
		json.Unmarshal(body, &register)
		json.Unmarshal(body, &register.Auth)

		if register.Auth.Password != register.Auth.ConfirmPassword {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Password and password confirmation do not match")
			return
		}

		if register.Auth.Email == "" || register.Auth.Password == "" || register.FirstName == "" || register.LastName == "" || register.BirthDate == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Password and password confirmation do not match")
			return
		}

		cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Auth.Password), 12)
		if err != nil {
			fmt.Println(err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Internal Error: There is a probleme with bcrypt")
			return
		}
		register.Auth.Password = string(cryptedPassword)


		uuid, err := uuid.NewV7()
		if err != nil {
			fmt.Println(err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Internal Error: There is a probleme with the generation of the uuid")
			return
		}
		register.Auth.Id = uuid.String()

		
		json, _ := json.Marshal(register)
		ctx := context.WithValue(r.Context(), model.RegisterCtx, json)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
