package product

import (
	"context"
	"errors"
	"full-ecommerce/internal/helpers"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidSKU  = errors.New("invalid sku")
	ErrInvalidName = errors.New("invalid name")
	ErrSKUExists   = errors.New("sku already exists")
)

type CreateProductInput struct {
	SKU              string         `json:"sku"`
	Name             string         `json:"name"`
	ShortDescription string         `json:"short_description"`
	Description      string         `json:"description"`
	Price            float64        `json:"price"`
	DiscountPrice    float64        `json:"discount_price"`
	Dimensions       Dimensions     `json:"dimensions"`
	ProfileImage     string         `json:"profile_image"`
	Images           []ProductImage `json:"images"`
	Variants         []Variant      `json:"variants"`
	Tags             []string       `json:"tags"`
}

func CreateProduct(input CreateProductInput) (primitive.ObjectID, error) {
	ctx := context.Background()

	if len(input.SKU) < 2 {
		return primitive.NilObjectID, ErrInvalidSKU
	}

	isValid, err := helpers.NameIsValid(input.Name)
	if err != nil || !isValid {
		return primitive.NilObjectID, ErrInvalidName
	}

	if ProductExistsBySKU(ctx, input.SKU) {
		return primitive.NilObjectID, ErrSKUExists
	}

	product := Product{
		ID:               primitive.NewObjectID(),
		SKU:              input.SKU,
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		Price:            input.Price,
		DiscountPrice:    input.DiscountPrice,
		Dimensions:       input.Dimensions,
		ProfileImage:     input.ProfileImage,
		Images:           input.Images,
		Variants:         input.Variants,
		Tags:             input.Tags,
	}

	return InsertProduct(ctx, product)
}

func ListProducts(filters ProductQueryFilters) ([]Product, error) {
	ctx := context.Background()
	return QueryProducts(ctx, filters)
}
