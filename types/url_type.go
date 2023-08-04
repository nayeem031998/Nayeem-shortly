package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"golang.org/x/net/ipv4"
)

type ShortUrlBody struct {
	LongUrl string `json:"long_url"`
}

type UrlDb struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UrlCode   string             `bson:"url_code"`
	LongUrl   string             `bson:"long_url"`
	Clicks    int64              `bson:"clicks"`
	Locations string           `bson:"locations"`
	ShortUrl  string             `bson:"short_url"`
	CreatedAt int64              `bson:"created_at"`
	ExpiredAt int64              `bson:"expired_at"`
}
