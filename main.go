package main

import (
	"log"

	"github.com/fleimkeipa/challengers-api/controller"
	"github.com/fleimkeipa/challengers-api/pkg"
	"github.com/fleimkeipa/challengers-api/repositories"
	"github.com/fleimkeipa/challengers-api/uc"
	"github.com/fleimkeipa/challengers-api/util"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	loadEnv()

	serveApplication()
}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println(".env file loaded successfully")
}

func serveApplication() {
	var e = echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Mongo
	var mongo = initDB()

	var userRepo = repositories.NewUserRepository(mongo)
	var userCommands = uc.NewUserUC(userRepo)
	var userHandlers = controller.NewUserHandlers(userCommands)

	var authRoutes = e.Group("/auth")
	authRoutes.POST("/register", userHandlers.Register)
	authRoutes.POST("/login", userHandlers.Login)

	var adminRoutes = e.Group("/admin")
	adminRoutes.Use(util.JWTAuth)
	adminRoutes.GET("", controller.Welcome)

	var challengerRoutes = e.Group("challenger")
	challengerRoutes.Use(util.JWTAuthChallenger)
	challengerRoutes.GET("", controller.Welcome)

	e.Logger.Fatal(e.Start(":8080"))
}

func initDB() *mongo.Database {
	mongo, err := pkg.MongoConnect()
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	return mongo
}
