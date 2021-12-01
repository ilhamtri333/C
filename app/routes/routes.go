package routes

import (
	_middleware "cleanarch/app/middleware"
	detailfasController "cleanarch/controllers/detail_fas"
	kostController "cleanarch/controllers/kost"
	kostFasController "cleanarch/controllers/kost_fasilitas"
	ownerController "cleanarch/controllers/owner"
	paymentController "cleanarch/controllers/payment_method"
	reservationController "cleanarch/controllers/reservation"
	userController "cleanarch/controllers/users"

	"github.com/labstack/echo/v4"
)

type RouteControllerList struct {
	UserController        userController.UserController
	KostController        kostController.KostController
	PaymentController     paymentController.PaymentController
	KostFasController     kostFasController.KostFasController
	OwnerController       ownerController.OwnerController
	DetailFasController   detailfasController.DetailFasController
	ReservationController reservationController.ReservationController
	JWTConfig             *_middleware.ConfigJWT
}

func (controller RouteControllerList) RouteRegister(e *echo.Echo) {
	//users
	users := e.Group("/user")
	users.POST("/login", controller.UserController.Login)
	users.POST("/register", controller.UserController.Register)
	users.GET("/", controller.UserController.GetAll)
	users.GET("/getid/:id", controller.UserController.GetUserById)
	users.PUT("/update/:id", controller.UserController.UpdateUser)
	users.DELETE("/delete/:id", controller.UserController.DeleteUser)

	//kost
	kost := e.Group("/kost")
	kost.POST("/insert", controller.KostController.InsertKost)
	kost.GET("/getid/:id", controller.KostController.GetKostById)
	kost.GET("/", controller.KostController.GetAllKost)
	kost.PUT("/update/:id", controller.KostController.UpdateKost)
	kost.DELETE("/delete/:id", controller.KostController.DeleteKost)

	//payment_method
	payment := e.Group("/payment")
	payment.POST("/insert", controller.PaymentController.InsertPayment)
	payment.GET("/getid/:id", controller.PaymentController.GetPaymentById)
	payment.GET("/", controller.PaymentController.GetAllPayment)
	payment.PUT("/update/:id", controller.PaymentController.UpdatePayment)
	payment.DELETE("/delete/:id", controller.PaymentController.DeletePayment)

	//kost_fasilitas
	kostfas := e.Group("/kostfas")
	kostfas.POST("/insert", controller.KostFasController.InsertKostFas)
	kostfas.GET("/getid/:id", controller.KostFasController.GetKostFasById)
	kostfas.GET("/", controller.KostFasController.GetAllKostFas)
	kostfas.PUT("/update/:id", controller.KostFasController.UpdateKostFas)
	kostfas.DELETE("/delete/:id", controller.KostFasController.DeleteKostFas)

	//owner
	owner := e.Group("/owner")
	owner.POST("/insert", controller.OwnerController.InsertOwner)
	owner.GET("/getid/:id", controller.OwnerController.GetOwnerById)
	owner.GET("/", controller.OwnerController.GetAllOwner)
	owner.PUT("/update/:id", controller.OwnerController.UpdateOwner)
	owner.DELETE("/delete/:id", controller.OwnerController.DeleteOwner)

	//detailfas
	detailfas := e.Group("/detailfas")
	detailfas.POST("/insert", controller.DetailFasController.InsertDetailFas)
	detailfas.GET("/getid/:id", controller.DetailFasController.GetDetailFasById)
	detailfas.GET("/", controller.DetailFasController.GetAllDetailFas)
	detailfas.PUT("/update/:id", controller.DetailFasController.UpdateDetailFas)
	detailfas.DELETE("/delete/:id", controller.DetailFasController.DeleteDetailFas)

	//reservation
	reservation := e.Group("/reservation")
	reservation.POST("/insert", controller.ReservationController.InsertReservation)
	reservation.GET("/getid/:id", controller.ReservationController.GetReservationById)
	reservation.GET("/", controller.ReservationController.GetAllReservation)
	reservation.PUT("/update/:id", controller.ReservationController.UpdateReservation)
	reservation.DELETE("/delete/:id", controller.ReservationController.DeleteReservation)
}
