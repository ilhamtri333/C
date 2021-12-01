package detailfas

import (
	detailfas "cleanarch/business/detail_fas"
	"cleanarch/controllers"
	"cleanarch/controllers/detail_fas/request"
	"cleanarch/controllers/detail_fas/response"
	"cleanarch/helper/konversi"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DetailFasController struct {
	dfusecase detailfas.DetailFasUseCaseInterface
}

func NewDetailFasController(dfc detailfas.DetailFasUseCaseInterface) *DetailFasController {
	return &DetailFasController{
		dfusecase: dfc,
	}
}

func (controller *DetailFasController) InsertDetailFas(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.InsertDetailFas{}
	err := c.Bind(&req)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err)
	}
	var data detailfas.Domain
	data, err = controller.dfusecase.InsertDetailFas(ctx, req.ToDomain())
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *DetailFasController) GetAllDetailFas(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := controller.dfusecase.GetAllDetailFas(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, response.DetailFasAll(data))
}

func (controller *DetailFasController) GetDetailFasById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	data, err := controller.dfusecase.GetDetailFasById(ctx, uint(konv))
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}
func (controller *DetailFasController) UpdateDetailFas(c echo.Context) error {
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	req := request.InsertDetailFas{}
	ctx := c.Request().Context()
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	data, err1 := controller.dfusecase.UpdateDetailFas(ctx, *req.ToDomain(), konv)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *DetailFasController) DeleteDetailFas(c echo.Context) error {
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	ctx := c.Request().Context()
	err := controller.dfusecase.DeleteDetailFas(ctx, konv)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.DetailFasResponse{ID: konv})
}
