package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}



func main(){
	fmt.Print("Hello World")

	app := fiber.New()

	todos := []Todo{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("Im OK 👋!")
	})

	log.Fatal(app.Listen(":3000"))

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(&todo); err != nil {
			return err
		}
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)
		return c.JSON(todo)
	})

	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprintf("%d", todo.ID) == id {
				todos[i].Completed = true
				return c.JSON(todos[i])
			}
		}
		return c.SendStatus(fiber.StatusNotFound)
	})

	
}
