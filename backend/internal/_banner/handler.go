package banner

import (
	"encoding/json"
	"full-ecommerce/internal/helpers"
	"full-ecommerce/pkg/response"
	"net/http"
)

func BannerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetBanner()
	}

	//verify if user is admin
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "User not logged in", http.StatusBadRequest)
			return
		}
		http.Error(w, "Error reading cookie.", http.StatusBadRequest)
		return
	}

	claims, err := helpers.ValidateToken(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if claims["role"] != "admin" {
		http.Error(w, "not authorized.", http.StatusUnauthorized)
	}

	//decode body and active post function
	var body Banner
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid_body", "Invalid JSON in request body.")
		return
	}

	PostBanner(body)
}
