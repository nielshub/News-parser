package repositories

import (
	"context"
	"incrowd/src/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type MongoDBRepository struct {
	CollectionName string
	Database       *mongo.Collection
}

func NewMongoDBRepository(collectionName string, DB *mongo.Collection) *MongoDBRepository {
	return &MongoDBRepository{
		CollectionName: collectionName,
		Database:       DB,
	}
}

func (repo *MongoDBRepository) ClearCollectionNews(ctx context.Context) error {
	return repo.Database.Drop(ctx)
}

func (repo *MongoDBRepository) StoreNews(ctx context.Context, news []model.News) error {
	for i := range news {
		_, err := repo.Database.InsertOne(ctx, news[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *MongoDBRepository) GetNewsWithID(ctx context.Context, id string) (*model.News, error) {
	var news model.News
	filter := bson.M{"id": id}
	err := repo.Database.FindOne(ctx, filter).Decode(&news)
	if err != nil {
		return &model.News{}, err
	}

	return &news, nil
}

func (repo *MongoDBRepository) GetNews(ctx context.Context) ([]model.News, error) {
	var newsArray []model.News
	cur, err := repo.Database.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	err = cur.All(ctx, &newsArray)
	if err != nil {
		return nil, err
	}

	return newsArray, nil
}
