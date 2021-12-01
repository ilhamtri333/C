package main

import (
	_middleware "cleanarch/app/middleware"
	"cleanarch/app/routes"
	detailFasUseCase "cleanarch/business/detail_fas"
	kostUseCase "cleanarch/business/kost"
	kostFasUseCase "cleanarch/business/kost_fasilitas"
	ownerUseCase "cleanarch/business/owner"
	payUseCase "cleanarch/business/payment_method"
	reservationUseCase "cleanarch/business/reservation"
	userUseCase "cleanarch/business/users"
	detailFasController "cleanarch/controllers/detail_fas"
	kostController "cleanarch/controllers/kost"
	kostFasController "cleanarch/controllers/kost_fasilitas"
	ownerController "cleanarch/controllers/owner"
	payController "cleanarch/controllers/payment_method"
	reservationController "cleanarch/controllers/reservation"
	userController "cleanarch/controllers/users"
	detailFasRepo "cleanarch/drivers/database/detail_fas"
	kostRepo "cleanarch/drivers/database/kost"
	kostFasRepo "cleanarch/drivers/database/kost_fasilitas"
	"cleanarch/drivers/database/mysql"
	ownerRepo "cleanarch/drivers/database/owner"
	payRepo "cleanarch/drivers/database/payment_method"
	reservationRepo "cleanarch/drivers/database/reservation"
	userRepo "cleanarch/drivers/database/users"
	"log"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile("app/config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}
func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&userRepo.User{})
	db.AutoMigrate(&kostRepo.Kost{})
	db.AutoMigrate(&payRepo.Payment{})
	db.AutoMigrate(&kostFasRepo.KostFas{})
	db.AutoMigrate(&ownerRepo.Owner{})
	db.AutoMigrate(&detailFasRepo.DetailFas{})
	db.AutoMigrate(&reservationRepo.Reservation{})
}

func main() {
	configDb := mysql.ConfigDB{
		DB_Username: viper.GetString("database.user"),
		DB_Password: viper.GetString("database.pass"),
		DB_Host:     viper.GetString("database.host"),
		DB_Port:     viper.GetString("database.port"),
		DB_Database: viper.GetString("database.name"),
	}
	db := configDb.InitialDb()
	dbMigrate(db)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()

	jwt := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString("jwt.secret"),
		ExpiresDuration: viper.GetInt("jwt.expired"),
	}

	//user
	userRepoInterface := userRepo.NewUserRepository(db)
	userUseCaseInterface := userUseCase.NewUseCase(userRepoInterface, timeoutContext, &jwt)
	userUseControllerInterface := userController.NewUserController(userUseCaseInterface)

	//kost
	kostRepoInterface := kostRepo.NewKostRepository(db)
	kostUseCaseInterface := kostUseCase.NewUseCase(kostRepoInterface, timeoutContext)
	kostUseControllerInterface := kostController.NewKostController(kostUseCaseInterface)

	//payment_method
	payRepoInterface := payRepo.NewPaymentRepository(db)
	payUseCaseInterface := payUseCase.NewUseCase(payRepoInterface, timeoutContext)
	payUseControllerInterface := payController.NewPayController(payUseCaseInterface)

	//kost_fasilitas
	kostFasRepoInterface := kostFasRepo.NewKostFasRepository(db)
	kostFasUseCaseInterface := kostFasUseCase.NewUseCase(kostFasRepoInterface, timeoutContext)
	kostFasUseControllerInterface := kostFasController.NewKostFasController(kostFasUseCaseInterface)

	//owner
	ownerRepoInterface := ownerRepo.NewOwnerRepository(db)
	ownerUseCaseInterface := ownerUseCase.NewUseCase(ownerRepoInterface, timeoutContext)
	ownerUseControllerInterface := ownerController.NewOwnerController(ownerUseCaseInterface)

	//detailfas
	detailFasRepoInterface := detailFasRepo.NewDetailFasRepository(db)
	detailFasUseCaseInterface := detailFasUseCase.NewUseCase(detailFasRepoInterface, timeoutContext)
	detailFasUseControllerInterface := detailFasController.NewDetailFasController(detailFasUseCaseInterface)

	//reservation
	reservationRepoInterface := reservationRepo.NewReservationRepository(db)
	reservationUseCaseInterface := reservationUseCase.NewUseCase(reservationRepoInterface, timeoutContext)
	reservationUseControllerInterface := reservationController.NewReservationController(reservationUseCaseInterface)

	routesInit := routes.RouteControllerList{
		UserController:        *userUseControllerInterface,
		KostController:        *kostUseControllerInterface,
		PaymentController:     *payUseControllerInterface,
		KostFasController:     *kostFasUseControllerInterface,
		OwnerController:       *ownerUseControllerInterface,
		DetailFasController:   *detailFasUseControllerInterface,
		ReservationController: *reservationUseControllerInterface,
	}
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
