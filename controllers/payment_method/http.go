package payment

import (
	payment "cleanarch/business/payment_method"
	"cleanarch/controllers"
	"cleanarch/controllers/payment_method/request"
	"cleanarch/controllers/payment_method/response"
	"cleanarch/helper/konversi"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	payusecase payment.PaymentMetodUsecaseInterface
}

func NewPayController(pc payment.PaymentMetodUsecaseInterface) *PaymentController {
	return &PaymentController{
		payusecase: pc,
	}
}

func (controller *PaymentController) InsertPayment(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.InsertPayment{}
	err := c.Bind(&req)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err)
	}
	var data payment.Domain
	data, err = controller.payusecase.InsertPayment(ctx, req.ToDomain())
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *PaymentController) GetAllPayment(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := controller.payusecase.GetAllPayment(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, response.PaymentAll(data))
}

func (controller *PaymentController) GetPaymentById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	data, err := controller.payusecase.GetPaymentById(ctx, uint(konv))
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}
func (controller *PaymentController) UpdatePayment(c echo.Context) error {
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	req := request.InsertPayment{}
	ctx := c.Request().Context()
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	data, err1 := controller.payusecase.UpdatePayment(ctx, *req.ToDomain(), konv)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *PaymentController) DeletePayment(c echo.Context) error {
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	ctx := c.Request().Context()
	err := controller.payusecase.DeletePayment(ctx, konv)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.PaymentResponse{Id: konv})
}
