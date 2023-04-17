package domain

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	Id         string         `json:"id" gorm:"primary_key; column:id;"`
	FullName   string         `json:"fullName" gorm:"full_name"`
	FirstOrder string         `json:"firstOrder" gorm:"column:first_order"`
	CreatedAt  time.Time      `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt  time.Time      `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at"`
}

type UserUseCase interface {
	CreateUser(context echo.Context, user User) error
	GetUserList(context echo.Context) ([]User, error)
	GetUserDetail(context echo.Context, id string) (User, error)
	UpdateUser(context echo.Context, editedUser User) error
	DeleteUser(context echo.Context, id string) error
}
