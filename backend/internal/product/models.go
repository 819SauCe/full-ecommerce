package product

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dimensions struct {
	WeightGrams int `bson:"weight_grams" json:"weight_grams"`
	WidthCM     int `bson:"width_cm" json:"width_cm"`
	HeightCM    int `bson:"height_cm" json:"height_cm"`
	DepthCM     int `bson:"depth_cm" json:"depth_cm"`
}

type ProductImage struct {
	URL  string `bson:"url" json:"url"`
	Sort int    `bson:"sort" json:"sort"`
}

type Variant struct {
	ColorName string `bson:"color_name" json:"color_name"`
	ColorHex  string `bson:"color_hex" json:"color_hex"`
	Stock     int    `bson:"stock" json:"stock"`
	ImageURL  string `bson:"image_url,omitempty" json:"image_url,omitempty"`
}

type Product struct {
	ID               primitive.ObjectID `bson:"_id" json:"id"`
	SKU              string             `bson:"sku" json:"sku"`
	Name             string             `bson:"name" json:"name"`
	ShortDescription string             `bson:"short_description" json:"short_description"`
	Description      string             `bson:"description" json:"description"`
	Price            float64            `bson:"price" json:"price"`
	DiscountPrice    float64            `bson:"discount_price,omitempty" json:"discount_price,omitempty"`
	Dimensions       Dimensions         `bson:"dimensions" json:"dimensions"`
	ProfileImage     string             `bson:"profile_image" json:"profile_image"`
	Images           []ProductImage     `bson:"images" json:"images"`
	Variants         []Variant          `bson:"variants" json:"variants"`
	Tags             []string           `bson:"tags" json:"tags"`
}

type ProductQueryFilters struct {
	Search   string   `json:"search"`
	Tags     []string `json:"tags"`
	MinPrice float64  `json:"min_price"`
	MaxPrice float64  `json:"max_price"`
	Page     int      `json:"page"`
	Limit    int      `json:"limit"`
}
