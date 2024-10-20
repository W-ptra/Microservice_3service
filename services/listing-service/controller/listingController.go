package controller

import (
	"strconv"
	"github.com/W-ptra/microservice_3service/listing-service/model"
	"github.com/gofiber/fiber/v2"
)

func GetListings(c *fiber.Ctx) error{
	/* /listing?{pageNum}&{pageSize}&{userId}*/
	pageNumStr := c.Query("pageNum")
	pageSizeStr := c.Query("pageSize")
	userIdStr := c.Query("userId")

	pageNum,pageSize,userId := 1,10,0

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

	if userIdStr != ""{
		num,err := strconv.Atoi(userIdStr)

		if err != nil{
			return c.Status(400).JSON(fiber.Map{
				"result":false,
				"message":"invalid query string userId: "+userIdStr,
			})
		}
		userId = num
	}

	listingList,err := model.GetAllListing(pageNum,pageSize,userId)

	if err!=nil && err.Error() == "record not found" || len(listingList) == 0{
		return c.Status(404).JSON(fiber.Map{
			"result":false,
			"message":"listing record is empty",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result": true,
		"listing": listingList,
	})
}

func CreateListing(c *fiber.Ctx) error{
	var listing model.Listing

	err := c.BodyParser(&listing)
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

	listing,err = model.CreateListing(listing)

	if err!=nil{
		return c.Status(422).JSON(fiber.Map{
			"result":"false",
			"message":"can't create listing",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"result":true,
		"listing": listing,
	})
}