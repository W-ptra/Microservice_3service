package user_api

import (
	"github.com/W-ptra/microservice_3service/user-service/controller"
	"github.com/W-ptra/microservice_3service/user-service/model"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"fmt"
	"os"
)

func Run(){
	model.Migration()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
    }))

	app.Use(logger.New())
	app.Get("/users/:id",controller.GetUserById)
	app.Get("/users",controller.GetAllUsers)
	app.Post("/users",controller.CreateUser)
	app.Use(controller.NotFound404)

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
	}

	app.Listen(fmt.Sprintf("%v:%v",os.Getenv("HOST"),os.Getenv("PORT")))
}