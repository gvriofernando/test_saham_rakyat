package usecase

import (
	"github.com/google/uuid"
	logger "github.com/gvriofernando/test_saham_rakyat/config/logger"
	"github.com/gvriofernando/test_saham_rakyat/domain"
	"github.com/gvriofernando/test_saham_rakyat/service/user/repository"
	"github.com/labstack/echo/v4"
)

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) domain.UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (u *userUseCase) CreateUser(ctx echo.Context, newUser domain.User) (err error) {
	newUser.Id = uuid.NewString()
	logger.Log.Printf("Create New User From Use Case, JSON Request: %v", newUser)
	err = u.userRepo.UpSertUser(ctx, newUser)
	if err != nil {
		return err
	}
	logger.Log.Printf("Create New User Succes From Use Case")
	return nil
}

func (u *userUseCase) GetUserList(ctx echo.Context) ([]domain.User, error) {
	res, err := u.userRepo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *userUseCase) GetUserDetail(ctx echo.Context, id string) (domain.User, error) {
	res, err := u.userRepo.GetUserById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return res, nil
}

func (u *userUseCase) UpdateUser(ctx echo.Context, editedUser domain.User) error {
	err := u.userRepo.UpSertUser(ctx, editedUser)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUseCase) DeleteUser(ctx echo.Context, id string) error {
	err := u.userRepo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
