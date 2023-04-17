package http

import (
	"net/http"

	logger "github.com/gvriofernando/test_saham_rakyat/config/logger"
	"github.com/gvriofernando/test_saham_rakyat/domain"
	"github.com/labstack/echo/v4"
)

type userController struct {
	userUseCase domain.UserUseCase
}

func NewUserController(e *echo.Echo, userUseCase domain.UserUseCase) {
	controller := &userController{
		userUseCase: userUseCase,
	}

	userHttp := e.Group("/users")
	userHttp.POST("", controller.CreateUser)
	userHttp.GET("/:id", controller.GetUserById)
	userHttp.GET("", controller.GetAllUser)
	userHttp.PUT("/:id", controller.UpdateUser)
	userHttp.DELETE("/:id", controller.DeleteUser)
}

func (controller *userController) CreateUser(ctx echo.Context) (err error) {
	u := new(domain.User)
	if err := ctx.Bind(&u); err != nil {
		logger.Log.Fatalf("Bad Request Create User, JSON Request: %v", u)
		return ctx.JSON(http.StatusBadRequest, u)
	}
	logger.Log.Printf("Create New User From Handler, JSON Request: %v", u)
	err = controller.userUseCase.CreateUser(ctx, *u)
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

func (controller *userController) GetUserById(ctx echo.Context) (err error) {
	id := ctx.Param("id")
	res, err := controller.userUseCase.GetUserDetail(ctx, id)
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

func (controller *userController) GetAllUser(ctx echo.Context) (err error) {
	res, err := controller.userUseCase.GetUserList(ctx)
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

func (controller *userController) UpdateUser(ctx echo.Context) (err error) {
	u := new(domain.User)
	if err := ctx.Bind(&u); err != nil {
		return ctx.JSON(http.StatusBadRequest, u)
	}
	id := ctx.Param("id")
	u.Id = id
	err = controller.userUseCase.UpdateUser(ctx, *u)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	apiRes := domain.Response{
		ErrorCode: 0,
		Message:   "success",
	}
	return ctx.JSON(http.StatusCreated, apiRes)
}

func (controller *userController) DeleteUser(ctx echo.Context) (err error) {
	id := ctx.Param("id")
	err = controller.userUseCase.DeleteUser(ctx, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	apiRes := domain.Response{
		ErrorCode: 0,
		Message:   "success",
	}
	return ctx.JSON(http.StatusOK, apiRes)
}
