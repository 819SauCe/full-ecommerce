package middlewares

import (
	"full-ecommerce/internal/helpers"
	"full-ecommerce/pkg/response"
	"net/http"
)

func RequireAuth(w http.ResponseWriter, r *http.Request) (map[string]interface{}, bool) {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		response.WriteError(w, http.StatusUnauthorized, "not_authenticated", "User is not authenticated.")
		return nil, false
	}

	claims, err := helpers.ValidateToken(cookie.Value)
	if err != nil {
		response.WriteError(w, http.StatusUnauthorized, "invalid_token", "Invalid or expired token.")
		return nil, false
	}

	return claims, true
}

func RequireRole(w http.ResponseWriter, r *http.Request, allowedRoles ...string) (map[string]interface{}, bool) {
	claims, ok := RequireAuth(w, r)
	if !ok {
		return nil, false
	}

	role, _ := claims["role"].(string)

	if len(allowedRoles) == 0 {
		return claims, true
	}

	for _, allowed := range allowedRoles {
		if role == allowed {
			return claims, true
		}
	}

	response.WriteError(w, http.StatusForbidden, "forbidden", "You do not have permission to perform this action.")
	return nil, false
}
