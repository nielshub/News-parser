package repositories

import (
	"context"
	"incrowd/src/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type MongoDBRepository struct {
	CollectionName string
	Database       *mongo.Database
}

func NewMongoDBRepository(collectionName string, DB *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		CollectionName: collectionName,
		Database:       DB,
	}
}

func (repo *MongoDBRepository) StoreNews(ctx context.Context, news []model.News) error {
	_, err := repo.Database.Collection(repo.CollectionName).InsertOne(ctx, news)
	//Check insertManny with
	//newValue := make([]interface{}, len(statements))
	//for i := range statements {
	// newValue[i] = statements[i]
	//}
	if err != nil {
		return err
	}

	return nil
}

func (repo *MongoDBRepository) GetNewsWithID(ctx context.Context, id string) (*model.News, error) {
	var news model.News
	filter := bson.M{"id": id}
	err := repo.Database.Collection(repo.CollectionName).FindOne(ctx, filter).Decode(&news)
	if err != nil {
		return &model.News{}, err
	}

	return &news, nil
}

func (repo *MongoDBRepository) GetNews(ctx context.Context) ([]model.News, error) {
	var newsArray []model.News
	cur, err := repo.Database.Collection(repo.CollectionName).Find(ctx, bson.D{})
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
