package controller

import (
	"strconv"
	"github.com/W-ptra/microservice_3service/user-service/model"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error{
	/* /users?{pageNum}&{pageSize} */
	pageNumStr := c.Query("pageNum")
	pageSizeStr := c.Query("pageSize")

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

	userList,err := model.GetAllUsers(pageNum,pageSize)

	if err!=nil && err.Error() == "record not found" || len(userList) == 0{
		return c.Status(404).JSON(fiber.Map{
			"result":false,
			"message":"users record is empty",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result": true,
		"users": userList,
	})
}

func GetUserById(c *fiber.Ctx) error{
	/* /users/{id} */
	idStr := c.Params("id")
	var id int

	if idStr == "" {
		return c.Status(400).JSON(fiber.Map{
			"result":  false,
			"message": "missing parameter id",
		})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"result":  false,
			"message": "invalid parameter id: " + idStr,
		})
	}

	user,err := model.GetUserById(id)
	if err !=nil || user.Id == 0{
		return c.Status(404).JSON(fiber.Map{
			"result":false,
			"message":"user not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result":true,
		"user":user,
	})
}

func CreateUser(c *fiber.Ctx) error{
	var user model.User

	err := c.BodyParser(&user)
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

	user,err = model.CreateUser(user)

	if err!=nil{
		return c.Status(422).JSON(fiber.Map{
			"result":"false",
			"message":"can't create user",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"result":true,
		"user": user,
	})
}