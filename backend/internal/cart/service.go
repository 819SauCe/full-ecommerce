package cart

import (
	"context"
	"errors"
)

var (
	ErrInvalidQuantity = errors.New("invalid quantity")
)

func GetUserCart(ctx context.Context, userID string) (int, []CartItem, error) {
	cartID, err := GetOrCreateCart(ctx, userID)
	if err != nil {
		return 0, nil, err
	}

	items, err := GetCartItems(ctx, cartID)
	return cartID, items, err
}

func AddToCart(ctx context.Context, userID string, input AddItemInput) error {
	if input.Quantity <= 0 {
		return ErrInvalidQuantity
	}

	cartID, err := GetOrCreateCart(ctx, userID)
	if err != nil {
		return err
	}

	return AddItem(ctx, cartID, input)
}

func UpdateCartItem(ctx context.Context, itemID int, quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}
	return UpdateItem(ctx, itemID, quantity)
}

func RemoveCartItem(ctx context.Context, itemID int) error {
	return RemoveItem(ctx, itemID)
}

func ClearUserCart(ctx context.Context, userID string) error {
	cartID, err := GetOrCreateCart(ctx, userID)
	if err != nil {
		return err
	}
	return ClearCart(ctx, cartID)
}
