package services

import (
	"context"
	"k6/eg2/interfaces"
	"k6/eg2/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	ctx             context.Context
	mongoCollection *mongo.Collection
	client          *mongo.Client
}



func InitializeCustomerService(ctx context.Context, collection *mongo.Collection, client *mongo.Client) interfaces.IDemo {
	return &CustomerService{ctx, collection, client}
}

func (c*CustomerService) CreateToken(customer *models.Sample) (*models.DBResponse, error) {
	res, err := c.mongoCollection.InsertOne(c.ctx, &customer)

	if err != nil {
		return nil, err
	}

	var newUser *models.DBResponse
	query := bson.M{"_id": res.InsertedID}

	err = c.mongoCollection.FindOne(c.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
