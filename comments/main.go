package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Comment struct {
	Id     uint   `json:"id"`
	PostId uint   `json:"post_id"`
	Text   string `json:"text"`
}

func main() {
	db, err := gorm.Open(postgres.Open("host=localhost user=omar dbname=comments password=ramo sslmode=disable"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Comment{})

	app := fiber.New()

	app.Use(cors.New())

	app.Post("/api/comments", func(c *fiber.Ctx) error {
		var comment Comment

		if err := c.BodyParser(&comment); err != nil {
			return err
		}

		db.Create(&comment)

		return c.JSON(comment)
	})

	app.Get("/api/posts/:id/comments", func(c *fiber.Ctx) error {
		var comments []Comment

		if err := db.Where("post_id =?", c.Params("id")).Find(&comments).Error; err != nil {
			return err
		}

		return c.JSON(comments)
	})

	app.Listen(":8001")
}
