package usecase

import (
	"context"
	"log"
	domain2 "traileau-auth-microservices/users/domain/repository"
	domain1 "traileau-auth-microservices/users/domain/usecase"
	models "traileau-auth-microservices/users/models"
)

type UserServiceImpl struct {
	userRepo domain2.UserRepositoryInterface
	ctx      context.Context
}

// DeleteUser implements usecase.UserUsecase.
func (u *UserServiceImpl) DeleteUser(ctx context.Context, req *string) error {
	panic("unimplemented")
}

// GetAll implements usecase.UserUsecase.
func (u UserServiceImpl) GetAll(ctx context.Context) ([]models.User, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := u.userRepo.GetAllData(ctx)
	if err != nil {
		log.Println("failed to show data user with default log")
		return list, err
	}

	return list, err
}

// GetUser implements usecase.UserUsecase.
func (u *UserServiceImpl) GetUser(ctx context.Context, req *string) (*models.User, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	user, err := u.userRepo.GetData(ctx, req)
	if err != nil {
		log.Println("failed to show data user with default log")
		return user, err
	}

	return user, err
}

// UpdateUser implements usecase.UserUsecase.
func (u *UserServiceImpl) UpdateUser(ctx context.Context, req *models.User) error {
	panic("unimplemented")
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

func NewUserUsecase(userRepo domain2.UserRepositoryInterface, ctx context.Context) domain1.UserUsecase {
	return &UserServiceImpl{
		userRepo: userRepo,
		ctx:      ctx,
	}
}
