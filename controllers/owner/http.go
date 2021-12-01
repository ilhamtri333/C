package owner

import (
	"cleanarch/business/owner"
	"cleanarch/controllers"
	"cleanarch/controllers/owner/request"
	"cleanarch/controllers/owner/response"
	"cleanarch/helper/konversi"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OwnerController struct {
	ownusecase owner.OwnerUseCaseInterface
}

func NewOwnerController(oc owner.OwnerUseCaseInterface) *OwnerController {
	return &OwnerController{
		ownusecase: oc,
	}
}

func (controller *OwnerController) InsertOwner(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.InsertOwner{}
	err := c.Bind(&req)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err)
	}
	var data owner.Domain
	data, err = controller.ownusecase.InsertOwner(ctx, req.ToDomain())
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *OwnerController) GetAllOwner(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := controller.ownusecase.GetAllOwner(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error binding", err)
	}
	return controllers.SuccessResponse(c, response.OwnerAll(data))
}

func (controller *OwnerController) GetOwnerById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	data, err := controller.ownusecase.GetOwnerById(ctx, uint(konv))
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}
func (controller *OwnerController) UpdateOwner(c echo.Context) error {
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	req := request.InsertOwner{}
	ctx := c.Request().Context()
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	data, err1 := controller.ownusecase.UpdateOwner(ctx, *req.ToDomain(), konv)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(data))
}

func (controller *OwnerController) DeleteOwner(c echo.Context) error {
	id := c.Param("id")
	konv, err1 := konversi.StringToUint(id)
	if err1 != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "bad request", err1)
	}
	ctx := c.Request().Context()
	err := controller.ownusecase.DeleteOwner(ctx, konv)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "internal error", err)
	}
	return controllers.SuccessResponse(c, response.OwnerResponse{ID: konv})
}
