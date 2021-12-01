package kost

import (
	"cleanarch/business/kost"
	"cleanarch/controllers"
	"cleanarch/controllers/kost/request"
	"cleanarch/controllers/kost/response"
	"cleanarch/helper/konversi"
	"net/http"

	"github.com/labstack/echo/v4"
)

type KostController struct {
	kostusecase kost.KostUseCaseInterface
}

func NewKostController(kc kost.KostUseCaseInterface) *KostController {
	return &KostController{
		kostusecase: kc,
	}
}

func (controller *KostController) InsertKost(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.InsertKost{}
	err := c.Bind(&req)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err)
	}
	var data kost.Domain
	data, err = controller.kostusecase.InsertKost(ctx, req.ToDomain())
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *KostController) GetAllKost(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := controller.kostusecase.GetAllKost(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, response.KostAll(data))
}

func (controller *KostController) GetKostById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	data, err := controller.kostusecase.GetKostById(ctx, uint(konv))
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}
func (controller *KostController) UpdateKost(c echo.Context) error {
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	req := request.InsertKost{}
	ctx := c.Request().Context()
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	data, err1 := controller.kostusecase.UpdateKost(ctx, *req.ToDomain(), konv)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *KostController) DeleteKost(c echo.Context) error {
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	ctx := c.Request().Context()
	err := controller.kostusecase.DeleteKost(ctx, konv)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.KostResponse{ID: konv})
}
