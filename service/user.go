package service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	IComplated bool `json:"isComplated"`
}


var users []User


func Create(c *fiber.Ctx) error {
	user := &User{}

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if user.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "name is required"})
	}


	user.ID = len(users) + 1
	users = append(users, *user)

	return c.Status(200).JSON(users)
}


func Update(c *fiber.Ctx) error {
	user := &User{}
	id := c.Params("id")
	err := c.BodyParser(user)
	fmt.Print(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if user.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "name is required"})
	}

	for i, u := range users {
		if fmt.Sprint(u.ID) == id {
			users[i].Name = string(user.Name)
			users[i].IComplated = user.IComplated

			return c.Status(200).JSON(users)
		}
	}

	return c.Status(400).JSON(fiber.Map{"error": "user not found"})
}


func Delete (c *fiber.Ctx)error{
	id := c.Params("id")
	for i, u := range users {
		if fmt.Sprint(u.ID) == id {
			users = append(users[:i], users[i+1:]...)
			return c.Status(200).JSON(users)
		}
	}
	return c.Status(400).JSON(fiber.Map{"error": "user not found"})
}

func Get(c *fiber.Ctx) error {
	return c.Status(200).JSON(users)
}