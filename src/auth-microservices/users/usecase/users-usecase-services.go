package users_usecase

import (
	"context"
	"log"
	domain "traileau/users/domain/repository"
	models "traileau/users/models"
)

type UserServiceImpl struct {
	userRepo domain.UserRepositoryInterface
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
