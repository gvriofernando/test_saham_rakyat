package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/gvriofernando/test_saham_rakyat/domain"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	GetOrderItemById(ctx echo.Context, id string) (domain.OrderItem, error)
	GetAllOrderItem(ctx echo.Context) ([]domain.OrderItem, error)
	UpSertOrderItem(ctx echo.Context, user domain.OrderItem) error
	DeleteOrderItem(ctx echo.Context, id string) error
}

type orderItemRepository struct {
	postgres *gorm.DB
	redis    *redis.Client
}

type OrderItemConfig struct {
	Postgres *gorm.DB
	Redis    *redis.Client
}

func NewOrderItemRepository(cfg OrderItemConfig) OrderItemRepository {
	return &orderItemRepository{
		postgres: cfg.Postgres,
		redis:    cfg.Redis,
	}
}

func (u *orderItemRepository) GetOrderItemById(ctx echo.Context, id string) (res domain.OrderItem, err error) {
	query := u.postgres.Where("id = ?", id).First(&res)
	if query.Error == gorm.ErrRecordNotFound {
		return domain.OrderItem{}, nil
	} else if query.Error != nil {
		return domain.OrderItem{}, query.Error
	}

	return res, nil
}

func (u *orderItemRepository) GetAllOrderItem(ctx echo.Context) (res []domain.OrderItem, err error) {
	query := u.postgres.Find(&res)
	if query.Error != nil {
		return []domain.OrderItem{}, query.Error
	}

	return res, nil
}

func (u *orderItemRepository) UpSertOrderItem(ctx echo.Context, orderItem domain.OrderItem) (err error) {
	query := u.postgres.Save(&orderItem)

	return query.Error
}

func (u *orderItemRepository) DeleteOrderItem(ctx echo.Context, id string) (err error) {
	query := u.postgres.Delete(&domain.OrderItem{Id: id})

	return query.Error
}

func (u *orderItemRepository) CheckExistingOrderItemTable(ctx echo.Context) error {
	hasTable := u.postgres.Migrator().HasTable(&domain.OrderItem{})
	if !hasTable {
		err := u.postgres.Migrator().CreateTable(&domain.OrderItem{})
		if err != nil {
			return err
		}
	}

	return nil
}
