package reservation

import (
	"cleanarch/business/reservation"
	"cleanarch/controllers"
	"cleanarch/controllers/reservation/request"
	"cleanarch/controllers/reservation/response"
	"cleanarch/helper/konversi"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ReservationController struct {
	rsusecase reservation.ReservationUseCaseInterface
}

func NewReservationController(rsc reservation.ReservationUseCaseInterface) *ReservationController {
	return &ReservationController{
		rsusecase: rsc,
	}
}

func (controller *ReservationController) InsertReservation(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.InsertReservation{}
	err := c.Bind(&req)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err)
	}
	var data reservation.Domain
	data, err = controller.rsusecase.InsertReservation(ctx, req.ToDomain())
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *ReservationController) GetAllReservation(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := controller.rsusecase.GetAllReservation(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, response.ReservationAll(data))
}

func (controller *ReservationController) GetReservationById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	data, err := controller.rsusecase.GetReservationById(ctx, uint(konv))
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}
func (controller *ReservationController) UpdateReservation(c echo.Context) error {
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	req := request.InsertReservation{}
	ctx := c.Request().Context()
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	data, err1 := controller.rsusecase.UpdateReservation(ctx, *req.ToDomain(), konv)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *ReservationController) DeleteReservation(c echo.Context) error {
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	ctx := c.Request().Context()
	err := controller.rsusecase.DeleteReservation(ctx, konv)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.ReservationResponse{Id: konv})
}
