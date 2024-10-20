package listing_api

import (
	"github.com/W-ptra/microservice_3service/listing-service/controller"
	"github.com/W-ptra/microservice_3service/listing-service/model"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"fmt"
	"os"
)

func Run(){
	model.Migration()
	app := fiber.New()

	app.Use(logger.New())
	app.Get("/listing",controller.GetListings)
	app.Post("/listing",controller.CreateListing)
	app.Use(controller.NotFound404)

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
	}
	
	app.Listen(fmt.Sprintf("%v:%v",os.Getenv("HOST"),os.Getenv("PORT")))
}