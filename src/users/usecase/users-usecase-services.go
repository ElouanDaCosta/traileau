package users_usecase

import (
	"context"
	"log"
	models "traileau/src/users/models"
	repository "traileau/src/users/repository"
)

type UserServiceImpl struct {
	userRepo repository.UserRepositoryI
	ctx      context.Context
}

func (u UserServiceImpl) CreateUser(ctx context.Context, req *models.User) error {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}
	//insert data
	err := u.userRepo.InsertData(ctx, req)
	if err != nil {
		return err
	}
	log.Println("Successfully Inserted Data User")

	return nil
}
