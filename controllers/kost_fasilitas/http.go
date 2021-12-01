package kostfasilitas

import (
	kostfasilitas "cleanarch/business/kost_fasilitas"
	"cleanarch/controllers"
	"cleanarch/controllers/kost_fasilitas/request"
	"cleanarch/controllers/kost_fasilitas/response"
	"cleanarch/helper/konversi"
	"net/http"

	"github.com/labstack/echo/v4"
)

type KostFasController struct {
	fasusecase kostfasilitas.KostFasUsecaseInterface
}

func NewKostFasController(kfc kostfasilitas.KostFasUsecaseInterface) *KostFasController {
	return &KostFasController{
		fasusecase: kfc,
	}
}

func (controller *KostFasController) InsertKostFas(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.InsertKostFas{}
	err := c.Bind(&req)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err)
	}
	var data kostfasilitas.Domain
	data, err = controller.fasusecase.InsertKostFas(ctx, req.ToDomain())
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *KostFasController) GetAllKostFas(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := controller.fasusecase.GetAllKostFas(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, response.KostFasAll(data))
}

func (controller *KostFasController) GetKostFasById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	data, err := controller.fasusecase.GetKostFasById(ctx, uint(konv))
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}
func (controller *KostFasController) UpdateKostFas(c echo.Context) error {
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	req := request.InsertKostFas{}
	ctx := c.Request().Context()
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	data, err1 := controller.fasusecase.UpdateKostFas(ctx, *req.ToDomain(), konv)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *KostFasController) DeleteKostFas(c echo.Context) error {
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	ctx := c.Request().Context()
	err := controller.fasusecase.DeleteKostFas(ctx, konv)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.KostFasResponse{Id: konv})
}
