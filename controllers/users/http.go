package controllers

import (
	"cleanarch/business/users"
	"cleanarch/controllers"
	"cleanarch/controllers/users/request"
	"cleanarch/controllers/users/response"
	"cleanarch/helper/konversi"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	usecase users.UserUseCaseInterface
}

func NewUserController(uc users.UserUseCaseInterface) *UserController {
	return &UserController{
		usecase: uc,
	}
}

func (controller *UserController) Login(c echo.Context) error {
	ctx := c.Request().Context()
	var userLogin request.UserLogin
	err := c.Bind(&userLogin)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	user, _ := controller.usecase.Login(*userLogin.ToDomain(), ctx)
	return controllers.SuccessResponse(c, response.FromDomain(user))
}

func (controller *UserController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := controller.usecase.GetAllUser(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, response.FromListDomain(data))
}

func (controller *UserController) Register(c echo.Context) error {
	ctx := c.Request().Context()
	reqRegist := request.UserRegister{}
	err := c.Bind(&reqRegist)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "Bad request", err)
	}
	user, err1 := controller.usecase.Register(reqRegist.ToDomain(), ctx)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, response.FromDomainToRegist(user))
}
func (controller *UserController) GetUserById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	data, err := controller.usecase.GetUserById(ctx, uint(konv))
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}
func (controller *UserController) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	req := request.UserRegister{}
	ctx := c.Request().Context()
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	data, err1 := controller.usecase.UpdateUser(ctx, *req.ToDomain(), konv)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	ctx := c.Request().Context()
	err := controller.usecase.DeleteUser(ctx, konv)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.UserResponse{Id: konv})
}
