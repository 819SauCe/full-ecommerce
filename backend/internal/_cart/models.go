package cart

type Cart struct {
	ID     int
	UserID string
	Status string
}

type CartItem struct {
	ID        int
	CartID    int
	ProductID string
	VariantID *string
	Quantity  int
}

type AddItemInput struct {
	ProductID string  `json:"product_id"`
	VariantID *string `json:"variant_id,omitempty"`
	Quantity  int     `json:"quantity"`
}

type UpdateItemInput struct {
	Quantity int `json:"quantity"`
}
