package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// middleware
	// you can control a path just add path before func
	app.Use(func(c *fiber.Ctx) error {
		//before
		fmt.Println("before")
		c.Next() // call every path
		//after
		fmt.Println("after")
		return nil
		//please, get error or return
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Get Hello fiber")
	})
	app.Post("", func(c *fiber.Ctx) error {
		return c.SendString("Post hello fiver")
	})

	//Parameter
	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString("name : " + name)
	})
	//parameter optionnal
	//use ? before surname ?surname, you can pass paramter or not pass
	app.Get("/hello/:name/:surname", func(c *fiber.Ctx) error {
		name := c.Params("name")
		surname := c.Params("surname")
		return c.SendString("name : " + name + " suername : " + surname)
	})
	//parameter int

	app.Get("/number/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.ErrBadRequest
		}
		return c.SendString(fmt.Sprintf("ID = %d", id))
	})
	//query string
	//example
	//case : localhost:8080/query?name=koe || case : localhost/query?name=koe&surname=srithong
	app.Get("/query", func(c *fiber.Ctx) error {
		name := c.Query("name")
		surname := c.Query("surname")
		return c.SendString("name : " + name + " surname : " + surname)
	})

	//query json struct
	//example
	// case : localhost:8080/query2?id=10&name=phongphat or
	// json {
	//			"id":2
	//			"name":"phongphat"
	//		}
	app.Get("/query2", func(c *fiber.Ctx) error {
		person := Person{}
		c.QueryParser(&person)
		return c.JSON(person)
	})
	//wildcards
	//example case : localhost:8080/wildcards/product/computer
	//return : product/computer
	app.Get("/wildcards/*", func(c *fiber.Ctx) error {
		wildcard := c.Params("*")
		return c.SendString(wildcard)
	})
	//staticfile
	//openbrowser to localhost:8080/staticpage to show http page from file index.html in wwwroot folder
	app.Static("/staticpage", "./wwwroot", fiber.Static{
		// you can config static
	})
	app.Get("/error", func(c *fiber.Ctx) error {
		// error return  status example status not found 404
		return fiber.NewError(fiber.StatusNotFound, "content not found")
	})

	app.Listen("localhost:8080")
}

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
