package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"github.com/W-ptra/microservice_3service/public-layer-service/model"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func GetListings(c *fiber.Ctx) error{
	err := godotenv.Load()
	if err!=nil{
		log.Fatalln("cant load .env variable")
	}

	pageNumStr := c.Query("pageNum")
	pageSizeStr := c.Query("pageSize")
	userIdStr := c.Query("userId")

	pageNum,pageSize := 1,10

	if pageNumStr != ""{
		num,err := strconv.Atoi(pageNumStr)

		if err != nil{
			return c.Status(400).JSON(fiber.Map{
				"result":false,
				"message":"invalid query string pageNum: "+pageNumStr,
			})
		}
		pageNum = num
	}

	if pageSizeStr != ""{
		num,err := strconv.Atoi(pageSizeStr)

		if err != nil{
			return c.Status(400).JSON(fiber.Map{
				"result":false,
				"message":"invalid query string pageSize: "+pageSizeStr,
			})
		}
		pageSize = num
	}

	var user model.User
	if userIdStr != ""{
		link := fmt.Sprintf("%v:%v/users/%v",os.Getenv("USER_SERVICE_HOST"),os.Getenv("USER_SERVICE_PORT"),userIdStr)

		resp,err := http.Get(link)
		if err != nil{
			return c.Status(404).JSON(fiber.Map{
				"result":false,
				"message":"user not found",
			})
		}
		defer resp.Body.Close()

		body,err := io.ReadAll(resp.Body)
		if err != nil{
			return c.Status(500).JSON(fiber.Map{
				"result":false,
				"message":"can't create user",
			})
		}

		var userRespond model.UserRespond
		if err := json.Unmarshal(body,&userRespond); err!=nil{
			return c.Status(500).JSON(fiber.Map{
				"result":false,
				"message":"failed to unmarshal json",
			})
		}

		user.Id = userRespond.User.Id
		user.Name = userRespond.User.Name
		user.CreatedAt = userRespond.User.CreatedAt
		user.UpdatedAt = userRespond.User.CreatedAt
	}

	link := fmt.Sprintf("%v:%v/listing?pageNum=%v&pageSize",
		os.Getenv("LISTING_SERVICE_HOST"),
		os.Getenv("LISTING_SERVICE_PORT"),
		pageNum,
		pageSize,
	)

	resp,err := http.Get(link)
	if err != nil{
		return c.Status(404).JSON(fiber.Map{
			"result":false,
			"message":"user not found",
		})
	}
	defer resp.Body.Close()

	body,err := io.ReadAll(resp.Body)
	if err != nil{
		return c.Status(500).JSON(fiber.Map{
			"result":false,
			"message":"can't create user",
		})
	}

	var ListingRespond model.ListingRespond
	if err := json.Unmarshal(body,&ListingRespond); err!=nil{
		return c.Status(500).JSON(fiber.Map{
			"result":false,
			"message":"failed to unmarshal json",
		})
	}

	if user.Id == 0{
		var listingWithUser model.ListingWithUserRespond
	}
}

func PostListings(c *fiber.Ctx) error{
	err := godotenv.Load()
	if err!=nil{
		log.Fatalln("cant load .env variable")
	}

	var listing model.Listing
	
	err = c.BodyParser(&listing)
	if err!=nil{
		return c.Status(400).JSON(fiber.Map{
			"result":"false",
			"message":"invalide JSON data",
		})
	}

	if listing.ListingType != "rent" && listing.ListingType != "sale"{
		return c.Status(400).JSON(fiber.Map{
			"result":"false",
			"message":"parameter listingType must be either 'rent' or 'sale'",
		})
	}

	if listing.Price < 0 {
		return c.Status(400).JSON(fiber.Map{
			"result":"false",
			"message":"parameter price can't negative",
		})
	}

	link := fmt.Sprintf("%v:%v/listing",os.Getenv("LISTING_SERVICE_HOST"),os.Getenv("LISTING_SERVICE_PORT"))
	
	formData := url.Values{}
	formData.Set("listingType",listing.ListingType)
	formData.Set("userId",strconv.Itoa(listing.UserId))
	formData.Set("price",strconv.Itoa(listing.Price))

	resp,err := http.Post(link,"application/x-www-form-urlencoded",
		io.NopCloser(bytes.NewBufferString(formData.Encode())),
	)
	if err != nil{
		return c.Status(500).JSON(fiber.Map{
			"result":false,
			"message":"can't create user",
		})
	}

	defer resp.Body.Close()

	body,err := io.ReadAll(resp.Body)

	if err != nil{
		return c.Status(500).JSON(fiber.Map{
			"result":false,
			"message":"can't create user",
		})
	}

	var ListingRespond model.ListingRespond
	if err := json.Unmarshal(body,&ListingRespond); err!=nil{
		return c.Status(500).JSON(fiber.Map{
			"result":false,
			"message":"failed to unmarshal json",
		})
	}

	return c.Status(201).JSON(ListingRespond)
}

func PostUser(c *fiber.Ctx) error{
	err := godotenv.Load()
	if err!=nil{
		log.Fatalln("cant load .env variable")
	}

	var user model.User

	err = c.BodyParser(&user)
	if err!=nil{
		return c.Status(400).JSON(fiber.Map{
			"result":"false",
			"message":"invalide JSON data",
		})
	}

	if user.Name == "" || len(user.Name)==0 || byte(user.Name[0]) == ' '{// check for empty or just ' ' char
		return c.Status(400).JSON(fiber.Map{
			"result":"false",
			"message":"field name can't empty or start with space (' ') character",
		})
	}
	
	link := fmt.Sprintf("%v:%v/users",os.Getenv("USER_SERVICE_HOST"),os.Getenv("USER_SERVICE_PORT"))

	formData := url.Values{}
	formData.Set("name",user.Name)

	resp,err := http.Post(link,"application/x-www-form-urlencoded",
		io.NopCloser(bytes.NewBufferString(formData.Encode())),
	)
	if err != nil{
		return c.Status(500).JSON(fiber.Map{
			"result":false,
			"message":"can't create user",
		})
	}
	defer resp.Body.Close()

	body,err := io.ReadAll(resp.Body)

	if err != nil{
		return c.Status(500).JSON(fiber.Map{
			"result":false,
			"message":"can't create user",
		})
	}

	var userRespond model.UserRespond
	if err := json.Unmarshal(body,&userRespond); err!=nil{
		return c.Status(500).JSON(fiber.Map{
			"result":false,
			"message":"failed to unmarshal json",
		})
	}

	return c.Status(201).JSON(userRespond)
}