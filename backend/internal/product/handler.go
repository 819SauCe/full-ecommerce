package product

import (
	"encoding/json"
	"errors"
	"full-ecommerce/pkg/response"
	"log"
	"net/http"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterProductRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/products", productsHandler)
	mux.HandleFunc("/products/", productByIDHandler)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreateProductHandler(w, r)
	case http.MethodGet:
		ListProductsHandler(w, r)
	default:
		response.WriteError(w, http.StatusMethodNotAllowed, "method_not_allowed",
			"Method not allowed. Use GET or POST.")
	}
}

func productByIDHandler(w http.ResponseWriter, r *http.Request) {
	idHex := strings.TrimPrefix(r.URL.Path, "/products/")
	if idHex == "" {
		response.WriteError(w, http.StatusBadRequest, "missing_id", "Product ID is required.")
		return
	}

	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid_id", "Invalid product ID.")
		return
	}

	switch r.Method {
	case http.MethodGet:
		GetProductHandler(w, r, id)
	case http.MethodDelete:
		DeleteProductHandler(w, r, id)
	default:
		response.WriteError(w, http.StatusMethodNotAllowed, "method_not_allowed",
			"Method not allowed. Use GET or DELETE.")
	}
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.WriteError(w, http.StatusMethodNotAllowed, "method_not_allowed",
			"Method not allowed. Use POST.")
		return
	}

	var body CreateProductInput
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid_body", "Invalid JSON in request body.")
		return
	}

	id, err := CreateProduct(body)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidSKU):
			response.WriteError(w, http.StatusBadRequest, "invalid_sku", "The provided SKU is invalid.")
			return
		case errors.Is(err, ErrInvalidName):
			response.WriteError(w, http.StatusBadRequest, "invalid_name", "The provided product name is invalid.")
			return
		case errors.Is(err, ErrSKUExists):
			response.WriteError(w, http.StatusConflict, "sku_exists", "This SKU is already in use.")
			return
		default:
			log.Printf("CreateProduct error: %v", err)
			response.WriteError(w, http.StatusInternalServerError, "product_error",
				"An error occurred while creating the product.")
			return
		}
	}

	response.WriteJSON(w, http.StatusCreated, map[string]any{
		"message": "Product created successfully.",
		"id":      id.Hex(),
		"status":  "ok",
	})
}

func GetProductHandler(w http.ResponseWriter, r *http.Request, id primitive.ObjectID) {
	p, err := FindProductByID(r.Context(), id)
	if err != nil {
		response.WriteError(w, http.StatusNotFound, "product_not_found", "Product not found.")
		return
	}

	response.WriteJSON(w, http.StatusOK, p)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request, id primitive.ObjectID) {
	err := DeleteProduct(r.Context(), id)
	if err != nil {
		response.WriteError(w, http.StatusNotFound, "product_not_found", "Product not found.")
		return
	}

	response.WriteJSON(w, http.StatusOK, map[string]any{
		"deleted": true,
		"status":  "ok",
	})
}

func ListProductsHandler(w http.ResponseWriter, r *http.Request) {
	filters := ProductQueryFilters{
		Search: r.URL.Query().Get("search"),
		Page:   parseInt(r.URL.Query().Get("page")),
		Limit:  parseInt(r.URL.Query().Get("limit")),
	}

	tagsParam := r.URL.Query().Get("tags")
	if tagsParam != "" {
		filters.Tags = strings.Split(tagsParam, ",")
	}

	filters.MinPrice = parseFloat(r.URL.Query().Get("min_price"))
	filters.MaxPrice = parseFloat(r.URL.Query().Get("max_price"))

	products, err := ListProducts(filters)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "list_error",
			"Could not retrieve products.")
		return
	}

	response.WriteJSON(w, http.StatusOK, map[string]any{
		"page":     filters.Page,
		"limit":    filters.Limit,
		"count":    len(products),
		"products": products,
	})
}

func parseInt(val string) int {
	n, _ := strconv.Atoi(val)
	return n
}

func parseFloat(val string) float64 {
	f, _ := strconv.ParseFloat(val, 64)
	return f
}
