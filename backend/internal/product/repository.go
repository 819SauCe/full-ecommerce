package product

import (
	"context"
	"full-ecommerce/internal/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func collection() *mongo.Collection {
	return config.MongoDB.Collection("products")
}

func InsertProduct(ctx context.Context, p Product) (primitive.ObjectID, error) {
	res, err := collection().InsertOne(ctx, p)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func FindProductBySKU(ctx context.Context, sku string) (Product, error) {
	var p Product
	err := collection().FindOne(ctx, bson.M{"sku": sku}).Decode(&p)
	return p, err
}

func FindProductByID(ctx context.Context, id primitive.ObjectID) (Product, error) {
	var p Product
	err := collection().FindOne(ctx, bson.M{"_id": id}).Decode(&p)
	return p, err
}

func ProductExistsBySKU(ctx context.Context, sku string) bool {
	err := collection().FindOne(ctx, bson.M{"sku": sku}).Err()
	return err == nil
}

func UpdateProduct(ctx context.Context, id primitive.ObjectID, update bson.M) error {
	_, err := collection().UpdateByID(ctx, id, bson.M{"$set": update})
	return err
}

func DeleteProduct(ctx context.Context, id primitive.ObjectID) error {
	_, err := collection().DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func QueryProducts(ctx context.Context, filters ProductQueryFilters) ([]Product, error) {
	query := bson.M{}

	if filters.Search != "" {
		query["$or"] = []bson.M{
			{"name": bson.M{"$regex": filters.Search, "$options": "i"}},
			{"description": bson.M{"$regex": filters.Search, "$options": "i"}},
			{"tags": bson.M{"$regex": filters.Search, "$options": "i"}},
		}
	}

	if len(filters.Tags) > 0 {
		query["tags"] = bson.M{"$in": filters.Tags}
	}

	priceFilter := bson.M{}
	if filters.MinPrice > 0 {
		priceFilter["$gte"] = filters.MinPrice
	}
	if filters.MaxPrice > 0 {
		priceFilter["$lte"] = filters.MaxPrice
	}
	if len(priceFilter) > 0 {
		query["price"] = priceFilter
	}

	page := filters.Page
	limit := filters.Limit
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	skip := (page - 1) * limit

	opts := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(limit)).
		SetSort(bson.D{{Key: "_id", Value: -1}})

	cursor, err := collection().Find(ctx, query, opts)
	if err != nil {
		return nil, err
	}

	var products []Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}
