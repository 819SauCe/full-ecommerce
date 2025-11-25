package banner

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBanner(ctx context.Context, b Banner) (primitive.ObjectID, error) {
	if b.Tittle == "" {
		return primitive.NilObjectID, errors.New("tittle is required")
	}
	return InsertBanner(ctx, b)
}

func GetBannerByID(ctx context.Context, id primitive.ObjectID) (Banner, error) {
	return FindBannerByID(ctx, id)
}

func ListBanners(ctx context.Context) ([]Banner, error) {
	return FindAllBanners(ctx)
}

func ListActiveBanners(ctx context.Context) ([]Banner, error) {
	return FindActiveBanners(ctx)
}

func UpdateBanner(ctx context.Context, id primitive.ObjectID, update bson.M) error {
	if len(update) == 0 {
		return errors.New("nothing to update")
	}
	return UpdateBannerByID(ctx, id, update)
}

func DeleteBannerByID(ctx context.Context, id primitive.ObjectID) error {
	return DeleteBanner(ctx, id)
}
