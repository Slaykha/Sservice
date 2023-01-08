package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Slaykha/Poll-App-Service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
}

func NewRepository(dbReplicaSetUrl string) *Repository {
	uri := dbReplicaSetUrl
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &Repository{client}
}

func (r *Repository) CreateUser(user models.User) error {
	collection := r.client.Database("pollapp").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindUser(email string) (*models.User, error) {
	collection := r.client.Database("pollapp").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"email": email}

	result := collection.FindOne(ctx, filter)

	user := models.User{}
	err := result.Decode(&user)
	if err != nil {
		fmt.Println("1")
		return nil, err
	}

	return &user, err
}

func (r *Repository) CreatePoll(poll models.Poll) error {
	collection := r.client.Database("pollapp").Collection("polls")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, poll)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetPolls() ([]models.Poll, error) {
	collection := r.client.Database("pollapp").Collection("polls")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var polls []models.Poll

	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for result.Next(ctx) {
		poll := models.Poll{}
		err := result.Decode(&poll)
		if err != nil {
			fmt.Println("3")
			return nil, err
		}

		polls = append(polls, poll)
	}

	if err != nil {
		fmt.Println("4")
		return nil, err
	}

	return polls, nil
}
