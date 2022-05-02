package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

type Post struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Comments    []Comment `json:"comments" gorm:"-" default:"[]"`
}

func main() {
	db, err := gorm.Open(postgres.Open("host=localhost user=omar dbname=posts password=ramo sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Post{})

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/api/posts", func(c *fiber.Ctx) error {
		var posts []Post

		db.Find(&posts)

		for i, post := range posts {
			res, err := http.Get(fmt.Sprintf("http://localhost:8001/api/posts/%d/comments", post.Id))
			if err != nil {
				return err
			}
			defer res.Body.Close()
			posts[i].Comments = []Comment{}
			if err := json.NewDecoder(res.Body).Decode(&posts[i].Comments); err != nil {
				return err
			}
		}

		return c.JSON(posts)
	})

	app.Post("/api/posts", func(c *fiber.Ctx) error {
		var post Post

		if err := c.BodyParser(&post); err != nil {
			return err
		}
		db.Create(&post)

		return c.JSON(post)
	})

	app.Listen(":8000")
}
