package cart

import (
	"encoding/json"
	"errors"
	"full-ecommerce/internal/middlewares"
	"full-ecommerce/pkg/response"
	"net/http"
	"strconv"
	"strings"
)

func RegisterCartRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/cart", handleCart)
	mux.HandleFunc("/cart/items", handleCartItems)
	mux.HandleFunc("/cart/items/", handleCartItemByID)
}

func handleCart(w http.ResponseWriter, r *http.Request) {
	claims, ok := middlewares.RequireAuth(w, r)
	if !ok {
		return
	}

	userID := claims["sub"].(string)

	switch r.Method {
	case http.MethodGet:
		cartID, items, err := GetUserCart(r.Context(), userID)
		if err != nil {
			response.WriteError(w, 500, "cart_error", err.Error())
			return
		}

		response.WriteJSON(w, 200, map[string]any{
			"cart_id": cartID,
			"items":   items,
		})

	case http.MethodDelete:
		err := ClearUserCart(r.Context(), userID)
		if err != nil {
			response.WriteError(w, 500, "cart_error", err.Error())
			return
		}
		response.WriteJSON(w, 200, map[string]string{"status": "cleared"})

	default:
		response.WriteError(w, 405, "method_not_allowed", "Use GET or DELETE")
	}
}

func handleCartItems(w http.ResponseWriter, r *http.Request) {
	claims, ok := middlewares.RequireAuth(w, r)
	if !ok {
		return
	}

	userID := claims["sub"].(string)

	if r.Method != http.MethodPost {
		response.WriteError(w, 405, "method_not_allowed", "Use POST")
		return
	}

	var input AddItemInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.WriteError(w, 400, "invalid_body", "Invalid JSON")
		return
	}

	err := AddToCart(r.Context(), userID, input)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidQuantity):
			response.WriteError(w, 400, "invalid_quantity", "Quantity must be greater than 0")
		default:
			response.WriteError(w, 500, "cart_error", err.Error())
		}
		return
	}

	response.WriteJSON(w, 201, map[string]string{"status": "item_added"})
}

func handleCartItemByID(w http.ResponseWriter, r *http.Request) {
	_, ok := middlewares.RequireAuth(w, r)
	if !ok {
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/cart/items/")
	itemID, _ := strconv.Atoi(idStr)

	switch r.Method {

	case http.MethodPut:
		var input UpdateItemInput
		json.NewDecoder(r.Body).Decode(&input)
		UpdateCartItem(r.Context(), itemID, input.Quantity)
		response.WriteJSON(w, 200, map[string]string{"status": "item_updated"})

	case http.MethodDelete:
		RemoveCartItem(r.Context(), itemID)
		response.WriteJSON(w, 200, map[string]string{"status": "item_deleted"})

	default:
		response.WriteError(w, 405, "method_not_allowed", "Use PUT or DELETE")
	}
}
