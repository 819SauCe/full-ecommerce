package auth

import (
	"encoding/json"
	"errors"
	"full-ecommerce/internal/helpers"
	"full-ecommerce/pkg/response"
	"log"
	"net/http"
	"time"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.WriteError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Method not allowed. Use POST.")
		return
	}

	var body RegisterModel
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid_body", "Invalid JSON in request body.")
		return
	}

	if err := Register(body); err != nil {
		switch {
		case errors.Is(err, ErrInvalidFirstName):
			response.WriteError(w, http.StatusBadRequest, "invalid_first_name", "The provided first name is invalid.")
			return
		case errors.Is(err, ErrInvalidLastName):
			response.WriteError(w, http.StatusBadRequest, "invalid_last_name", "The provided last name is invalid.")
			return
		case errors.Is(err, ErrInvalidEmail):
			response.WriteError(w, http.StatusBadRequest, "invalid_email", "The provided email is invalid.")
			return
		case errors.Is(err, ErrEmailAlreadyUsed):
			response.WriteError(w, http.StatusConflict, "email_already_used", "This email is already in use.")
			return
		case errors.Is(err, ErrInvalidPassword):
			response.WriteError(w, http.StatusBadRequest, "invalid_password", "The password does not meet the minimum requirements.")
			return
		default:
			log.Printf("Register error: %v", err)
			response.WriteError(w, http.StatusInternalServerError, "register_error", "Error while registering user.")
			return
		}
	}

	userData, err := GetUserDataByEmail(body.Email)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "user_lookup_error", "Error while fetching user data.")
		return
	}

	token, err := helpers.GenerateToken(userData)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "token_generation_error", "Error generating authentication token.")
		return
	}

	helpers.SetAuthCookie(w, token)

	response.WriteJSON(w, http.StatusCreated, map[string]any{
		"message": "User created successfully",
		"status":  "ok",
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.WriteError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Method not allowed. Use POST.")
		return
	}

	var body LoginModel
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid_body", "Invalid JSON in request body.")
		return
	}

	if err := Login(body); err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			response.WriteError(w, http.StatusUnauthorized, "invalid_credentials", "Invalid credentials. Please check your details and try again.")
			return
		}

		response.WriteError(w, http.StatusInternalServerError, "login_error", "Error logging in.")
		return
	}

	userData, err := GetUserDataByEmail(body.Email)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			response.WriteError(w, http.StatusUnauthorized, "invalid_credentials", "Invalid credentials. Please check your details and try again.")
			return
		}

		response.WriteError(w, http.StatusInternalServerError, "user_lookup_error", "Error retrieving user data.")
		return
	}

	token, err := helpers.GenerateToken(userData)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "token_generation_error", "Error generating authentication token.")
		return
	}

	helpers.SetAuthCookie(w, token)

	response.WriteJSON(w, http.StatusAccepted, map[string]any{
		"message": "User logged successfully",
		"status":  "ok",
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "User not logged in", http.StatusBadRequest)
			return
		}
		http.Error(w, "Error reading cookie.", http.StatusBadRequest)
		return
	}

	_ = Logout(cookie.Value)

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
}

func MeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Method not allowed. Use GET.")
		return
	}

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

	response.WriteJSON(w, http.StatusAccepted, map[string]any{
		"status": "ok",
		"claims": claims,
	})
}
