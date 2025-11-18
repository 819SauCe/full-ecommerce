package auth

import (
	"encoding/json"
	"full-ecommerce/internal/helpers"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body RegisterModel

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	if err := Register(body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, role, err := GetUserIDAndRoleByEmail(body.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := helpers.GenerateToken(id, role)
	if err != nil {
		http.Error(w, "error generating token", http.StatusInternalServerError)
		return
	}

	helpers.SetAuthCookie(w, token)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "User created successfully",
		"status":  "ok",
	})
}
