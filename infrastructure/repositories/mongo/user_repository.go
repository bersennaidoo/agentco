package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
}

/*func (u *UserRepository) PostUsers(user models.User) models.User {
}*/
