package mongo

import "github.com/bersennaidoo/agentco/domain/models"

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
}

func (u *UserRepository) PostUsers(user models.User) models.User {
}
