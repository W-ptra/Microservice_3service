package public_layer_api

import (
	"fmt"
	"log"
	"os"
	"github.com/W-ptra/microservice_3service/public-layer-service/controller"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Run(){
	app := fiber.New()

	app.Use(logger.New())
	app.Post("/public-api/users",controller.PostUser)
	app.Get("/public-api/listing",controller.GetListings)
	app.Post("/public-api/listing",controller.PostListings)
	app.Use(controller.NotFound404)

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
	}

	app.Listen(fmt.Sprintf("%v:%v",os.Getenv("HOST"),os.Getenv("PORT")))
}