package repository

import (
	"github.com/go-redis/redis/v8"
	logger "github.com/gvriofernando/test_saham_rakyat/config/logger"
	"github.com/gvriofernando/test_saham_rakyat/domain"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserById(ctx echo.Context, id string) (domain.User, error)
	GetAllUsers(ctx echo.Context) ([]domain.User, error)
	UpSertUser(ctx echo.Context, user domain.User) error
	DeleteUser(ctx echo.Context, id string) error
}

type userRepository struct {
	postgres *gorm.DB
	redis    *redis.Client
}

type UserConfig struct {
	Postgres *gorm.DB
	Redis    *redis.Client
}

func NewUserRepository(cfg UserConfig) UserRepository {
	return &userRepository{
		postgres: cfg.Postgres,
		redis:    cfg.Redis,
	}
}

func (u *userRepository) GetUserById(ctx echo.Context, id string) (res domain.User, err error) {
	query := u.postgres.Where("id = ?", id).First(&res)
	if query.Error == gorm.ErrRecordNotFound {
		return domain.User{}, nil
	} else if query.Error != nil {
		return domain.User{}, query.Error
	}

	return res, nil
}

func (u *userRepository) GetAllUsers(ctx echo.Context) (res []domain.User, err error) {
	query := u.postgres.Find(&res)
	if query.Error != nil {
		return []domain.User{}, query.Error
	}

	return res, nil
}

func (u *userRepository) UpSertUser(ctx echo.Context, user domain.User) (err error) {
	logger.Log.Printf("Create New User From Repository, JSON Request: %v", user)
	query := u.postgres.Save(&user)
	if query.Error != nil {
		return query.Error
	}
	logger.Log.Printf("Create New User From Repository Success")
	return nil
}

func (u *userRepository) DeleteUser(ctx echo.Context, id string) (err error) {
	query := u.postgres.Delete(&domain.User{Id: id})

	return query.Error
}

func (u *userRepository) CheckExistingUserTable(ctx echo.Context) error {
	hasTable := u.postgres.Migrator().HasTable(&domain.User{})
	if !hasTable {
		err := u.postgres.Migrator().CreateTable(&domain.User{})
		if err != nil {
			return err
		}
	}

	return nil
}
