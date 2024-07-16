package user

import (
	"api/autentiacion/internal/domain"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"os"
	"time"
)

func (s Service) Create(ctx context.Context, params domain.UserCreateParams) (*domain.User, error) {
	now := time.Now().UTC()
	_, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	user := &domain.User{
		ID:           "123", // TODO: generate user ID,
		Username:     params.Username,
		Email:        params.Email,
		PasswordHash: params.Password,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	// TODO: connect to mongo and save user
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("mydb").Collection("users")
	insertResult, err2 := collection.InsertOne(ctx, user)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer client.Disconnect(ctx)

	return &domain.User{
		ID:           insertResult.InsertedID.(string),
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}, nil
}

func (s Service) Register(ctx context.Context, params domain.UserCreateParams) (*domain.User, error) {
	return nil, nil
}

func (s Service) Login(ctx context.Context, username string, password string) (*domain.User, error) {
	return nil, nil
}
