package public_layer_api

import (
	"fmt"
	"log"
	"os"
	"github.com/W-ptra/microservice_3service/public-layer-service/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Run(){
	app := fiber.New()
	//app.Get("/users/:id",controller.GetUserById)
	//app.Get("/users",controller.GetAllUser)
	app.Post("/users",controller.PostUser)
	app.Post("/listing",controller.PostListings)
	app.Use(controller.NotFound404)

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
	}

	app.Listen(fmt.Sprintf("%v:%v",os.Getenv("HOST"),os.Getenv("PORT")))
}