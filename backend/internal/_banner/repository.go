package banner

import (
	"context"
	"full-ecommerce/internal/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func collection() *mongo.Collection {
	return config.MongoDB.Collection("banners")
}

func InsertBanner(ctx context.Context, b Banner) (primitive.ObjectID, error) {
	res, err := collection().InsertOne(ctx, b)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func FindBannerByID(ctx context.Context, id primitive.ObjectID) (Banner, error) {
	var b Banner
	err := collection().FindOne(ctx, bson.M{"_id": id}).Decode(&b)
	return b, err
}

func FindAllBanners(ctx context.Context) ([]Banner, error) {
	cursor, err := collection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var banners []Banner
	if err := cursor.All(ctx, &banners); err != nil {
		return nil, err
	}

	return banners, nil
}

func FindActiveBanners(ctx context.Context) ([]Banner, error) {
	cursor, err := collection().Find(ctx, bson.M{"active": true}, options.Find().SetSort(bson.D{{Key: "position", Value: 1}}))
	if err != nil {
		return nil, err
	}

	var banners []Banner
	if err := cursor.All(ctx, &banners); err != nil {
		return nil, err
	}

	return banners, nil
}

func DeleteBanner(ctx context.Context, id primitive.ObjectID) error {
	_, err := collection().DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func UpdateBannerByID(ctx context.Context, id primitive.ObjectID, update bson.M) error {
	_, err := collection().UpdateByID(ctx, id, bson.M{"$set": update})
	return err
}
