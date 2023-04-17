package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type OrderItem struct {
	Id        string    `json:"id" gorm:"primary_key; column:id;"`
	Name      string    `json:"name" gorm:"name"`
	Price     float64   `json:"price" gorm:"price"`
	ExpiredAt string    `json:"expiredAt" gorm:"column:expired_at"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt time.Time `json:"deletedAt" gorm:"column:deleted_at"`
}

type OrderItemUseCase interface {
	CreateOrderItem(context echo.Context, orderItem OrderItem) error
	GetOrderItemList(context echo.Context) ([]OrderItem, error)
	GetOrderItemDetail(context echo.Context, id string) (OrderItem, error)
	UpdateOrderItem(context echo.Context, editedOrderItem OrderItem) error
	DeleteOrderItem(context echo.Context, id string) error
}
