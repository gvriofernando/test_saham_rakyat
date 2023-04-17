package http

import (
	"net/http"

	logger "github.com/gvriofernando/test_saham_rakyat/config/logger"
	"github.com/gvriofernando/test_saham_rakyat/domain"
	"github.com/labstack/echo/v4"
)

type orderItemController struct {
	orderItemUseCase domain.OrderItemUseCase
}

func NewOrderItemController(e *echo.Echo, orderItemUseCase domain.OrderItemUseCase) {
	controller := &orderItemController{
		orderItemUseCase: orderItemUseCase,
	}

	userHttp := e.Group("/order-items")
	userHttp.POST("", controller.CreateOrderItem)
	userHttp.GET("/:id", controller.GetOrderItemById)
	userHttp.GET("", controller.GetAllOrderItem)
	userHttp.PUT("/:id", controller.UpdateOrderItem)
	userHttp.DELETE("/:id", controller.DeleteOrderItem)
}

func (controller *orderItemController) CreateOrderItem(ctx echo.Context) (err error) {
	u := new(domain.OrderItem)
	if err := ctx.Bind(&u); err != nil {
		logger.Log.Fatalf("Bad Request Create User, JSON Request: %v", u)
		return ctx.JSON(http.StatusBadRequest, u)
	}
	logger.Log.Printf("Create New User From Handler, JSON Request: %v", u)
	err = controller.orderItemUseCase.CreateOrderItem(ctx, *u)
	if err != nil {
		logger.Log.Fatalf("Server Error Create User, JSON Request: %v", u)
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	logger.Log.Printf("Create New User Success")
	apiRes := domain.Response{
		ErrorCode: 0,
		Message:   "success",
	}

	return ctx.JSON(http.StatusCreated, apiRes)
}

func (controller *orderItemController) GetOrderItemById(ctx echo.Context) (err error) {
	id := ctx.Param("id")
	res, err := controller.orderItemUseCase.GetOrderItemDetail(ctx, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	apiRes := domain.Response{
		ErrorCode: 0,
		Message:   "success",
		Data:      res,
	}

	return ctx.JSON(http.StatusOK, apiRes)
}

func (controller *orderItemController) GetAllOrderItem(ctx echo.Context) (err error) {
	res, err := controller.orderItemUseCase.GetOrderItemList(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	apiRes := domain.Response{
		ErrorCode: 0,
		Message:   "success",
		Data:      res,
	}

	return ctx.JSON(http.StatusOK, apiRes)
}

func (controller *orderItemController) UpdateOrderItem(ctx echo.Context) (err error) {
	u := new(domain.OrderItem)
	if err := ctx.Bind(&u); err != nil {
		return ctx.JSON(http.StatusBadRequest, u)
	}
	id := ctx.Param("id")
	u.Id = id
	err = controller.orderItemUseCase.UpdateOrderItem(ctx, *u)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	apiRes := domain.Response{
		ErrorCode: 0,
		Message:   "success",
	}
	return ctx.JSON(http.StatusCreated, apiRes)
}

func (controller *orderItemController) DeleteOrderItem(ctx echo.Context) (err error) {
	id := ctx.Param("id")
	err = controller.orderItemUseCase.DeleteOrderItem(ctx, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	apiRes := domain.Response{
		ErrorCode: 0,
		Message:   "success",
	}
	return ctx.JSON(http.StatusOK, apiRes)
}
