package cart

import (
	"context"
	"full-ecommerce/internal/config"
)

func GetOrCreateCart(ctx context.Context, userID string) (int, error) {
	var cartID int

	err := config.DB.QueryRowContext(ctx,
		`SELECT id FROM carts WHERE user_id = $1 AND status = 'open'`,
		userID,
	).Scan(&cartID)

	if err == nil {
		return cartID, nil
	}

	err = config.DB.QueryRowContext(ctx,
		`INSERT INTO carts (user_id) VALUES ($1) RETURNING id`,
		userID,
	).Scan(&cartID)

	return cartID, err
}

func AddItem(ctx context.Context, cartID int, input AddItemInput) error {
	_, err := config.DB.ExecContext(ctx, `
        INSERT INTO cart_items (cart_id, product_id, variant_id, quantity)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (cart_id, product_id, variant_id)
        DO UPDATE SET quantity = cart_items.quantity + EXCLUDED.quantity
    `, cartID, input.ProductID, input.VariantID, input.Quantity)

	return err
}

func UpdateItem(ctx context.Context, itemID int, quantity int) error {
	_, err := config.DB.ExecContext(ctx,
		`UPDATE cart_items SET quantity = $1 WHERE id = $2`,
		quantity, itemID,
	)
	return err
}

func RemoveItem(ctx context.Context, itemID int) error {
	_, err := config.DB.ExecContext(ctx,
		`DELETE FROM cart_items WHERE id = $1`,
		itemID,
	)
	return err
}

func ClearCart(ctx context.Context, cartID int) error {
	_, err := config.DB.ExecContext(ctx,
		`DELETE FROM cart_items WHERE cart_id = $1`,
		cartID,
	)
	return err
}

func GetCartItems(ctx context.Context, cartID int) ([]CartItem, error) {
	rows, err := config.DB.QueryContext(ctx,
		`SELECT id, product_id, variant_id, quantity
         FROM cart_items WHERE cart_id = $1`,
		cartID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []CartItem

	for rows.Next() {
		var item CartItem
		err := rows.Scan(&item.ID, &item.ProductID, &item.VariantID, &item.Quantity)
		if err != nil {
			return nil, err
		}
		item.CartID = cartID
		items = append(items, item)
	}

	return items, nil
}
