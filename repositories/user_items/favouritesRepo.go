package user_items

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"petstore-api/models"
	"time"
)

type favouriteRepo struct {
	collection *mongo.Collection
}

func NewFavouriteRepository(client *mongo.Client) *favouriteRepo {
	collection := client.Database("buckets").Collection("buckets")
	return &favouriteRepo{
		collection: collection,
	}
}

func (b *favouriteRepo) AddPet(userID uint, itemID uint) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userID}
	update := bson.M{"$push": bson.M{"pets": itemID}}

	_, err := b.collection.UpdateOne(ctx, filter, update)
	return err
}

func (b *favouriteRepo) RemovePet(userID uint, itemID uint) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userID}
	update := bson.M{"$pull": bson.M{"pets": itemID}}

	_, err := b.collection.UpdateOne(ctx, filter, update)
	return err
}

func (b *favouriteRepo) GetPetsByBuyerID(userID uint) ([]models.Pet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result models.Bucket

	err := b.collection.FindOne(ctx, bson.M{"userId": userID}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return []models.Pet{}, nil
		}
		return nil, err
	}

	return result.Pets, nil
}
