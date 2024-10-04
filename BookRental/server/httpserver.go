package server

import (
	"database/sql"
	"log"
	"bookrental/controllers"
	"bookrental/repositories"
	"bookrental/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	usersController   *controllers.UsersController
	adminsController  *controllers.AdminsController
	rentalsController *controllers.RentalsController
	ratingsController *controllers.RatingsController
	preferredBooksController *controllers.PreferredBooksController
	booksController   *controllers.BooksController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	usersRepository := repositories.NewUsersRepository(dbHandler)
	// adminsRepository := repositories.NewAdminsRepository(dbHandler)
	// rentalsRepository := repositories.NewRentalsRepository(dbHandler)
	// ratingsRepository := repositories.NewRatingsRepository(dbHandler)
	// preferredBooksRepository := repositories.NewPreferredBooksRepository(dbHandler)
	// booksRepository := repositories.NewBooksRepository(dbHandler)

	usersService := services.NewUsersService(usersRepository)
	// adminsService := services.NewAdminsService(adminsRepository)
	// rentalsService := services.NewRentalsService(rentalsRepository)
	// ratingsService := services.NewRatingsService(ratingsRepository)
	// preferredBooksService := services.NewPreferredBooksService(preferredBooksRepository)
	// booksService := services.NewBooksService(booksRepository)

	usersController := controllers.NewUsersController(usersService)
	// adminsController := controllers.NewAdminsController(adminsService)
	// rentalsController := controllers.NewRentalsController(rentalsService)
	// ratingsController := controllers.NewRatingsController(ratingsService)
	// preferredBooksController := controllers.NewPreferredBooksController(preferredBooksService)
	// booksController := controllers.NewBooksController(booksService)

	router := gin.Default()

	// User routes
	router.POST("/user", usersController.CreateUser)
	router.PUT("/user", usersController.UpdateUser)
	router.DELETE("/user/:id", usersController.DeleteUser)
	router.GET("/user/:id", usersController.GetUser)
	router.GET("/users", usersController.GetUsersBatch)

	// Admin routes
	// router.POST("/admin", adminsController.CreateAdmin)
	// router.PUT("/admin", adminsController.UpdateAdmin)
	// router.DELETE("/admin/:id", adminsController.DeleteAdmin)
	// router.GET("/admin/:id", adminsController.GetAdmin)

	// // Rental routes
	// router.POST("/rental", rentalsController.CreateRental)
	// router.GET("/rental/:id", rentalsController.GetRental)
	// router.GET("/rentals", rentalsController.GetRentals)

	// // Rating routes
	// router.POST("/rating", ratingsController.CreateRating)

	// // PreferredBook routes
	// router.POST("/preferred-book", preferredBooksController.CreatePreferredBook)

	// // Book routes
	// router.POST("/book", booksController.CreateBook)
	// router.PUT("/book", booksController.UpdateBook)
	// router.DELETE("/book/:id", booksController.DeleteBook)
	// router.GET("/book/:id", booksController.GetBook)

	return HttpServer{
		config:            config,
		router:            router,
		usersController:   usersController,
		// adminsController:  adminsController,
		// rentalsController: rentalsController,
		// ratingsController: ratingsController,
		// preferredBooksController: preferredBooksController,
		// booksController:   booksController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
