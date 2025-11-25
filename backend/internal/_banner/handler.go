package banner

import (
	"encoding/json"
	"full-ecommerce/internal/helpers"
	"full-ecommerce/pkg/response"
	"net/http"
)

func RegisterBannerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ecommerce/banner", BannerHandler)
}

func BannerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch r.Method {

	case http.MethodGet:
		banners, err := ListBanners(ctx)
		if err != nil {
			response.WriteError(w, http.StatusBadRequest, "db_error", err.Error())
			return
		}

		response.WriteJSON(w, http.StatusOK, banners)
		return

	case http.MethodPost:
		// verify if user is admin
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			http.Error(w, "User not logged in", http.StatusBadRequest)
			return
		}

		claims, err := helpers.ValidateToken(cookie.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if claims["role"] != "admin" {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		// decode body
		var body Banner
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			response.WriteError(w, http.StatusBadRequest, "invalid_body", "Invalid JSON.")
			return
		}

		id, err := CreateBanner(ctx, body)
		if err != nil {
			response.WriteError(w, http.StatusInternalServerError, "db_error", err.Error())
			return
		}

		response.WriteJSON(w, http.StatusCreated, map[string]interface{}{
			"id":     id,
			"status": "banner_created",
		})
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
