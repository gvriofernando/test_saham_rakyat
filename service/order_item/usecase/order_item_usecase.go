package usecase

import (
	"github.com/google/uuid"
	"github.com/gvriofernando/test_saham_rakyat/domain"
	"github.com/gvriofernando/test_saham_rakyat/service/order_item/repository"
	"github.com/labstack/echo/v4"
)

type orderItemUseCase struct {
	orderItemRepo repository.OrderItemRepository
}

func NewOrderItemUseCase(orderItemRepo repository.OrderItemRepository) domain.OrderItemUseCase {
	return &orderItemUseCase{
		orderItemRepo: orderItemRepo,
	}
}

func (u *orderItemUseCase) CreateOrderItem(ctx echo.Context, newOrderItem domain.OrderItem) (err error) {
	newOrderItem.Id = uuid.NewString()
	err = u.orderItemRepo.UpSertOrderItem(ctx, newOrderItem)
	if err != nil {
		return err
	}
	return nil
}

func (u *orderItemUseCase) GetOrderItemList(ctx echo.Context) ([]domain.OrderItem, error) {
	res, err := u.orderItemRepo.GetAllOrderItem(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *orderItemUseCase) GetOrderItemDetail(ctx echo.Context, id string) (domain.OrderItem, error) {
	res, err := u.orderItemRepo.GetOrderItemById(ctx, id)
	if err != nil {
		return domain.OrderItem{}, err
	}
	return res, nil
}

func (u *orderItemUseCase) UpdateOrderItem(ctx echo.Context, editedOrderItem domain.OrderItem) error {
	err := u.orderItemRepo.UpSertOrderItem(ctx, editedOrderItem)
	if err != nil {
		return err
	}
	return nil
}

func (u *orderItemUseCase) DeleteOrderItem(ctx echo.Context, id string) error {
	err := u.orderItemRepo.DeleteOrderItem(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
