package main

import (
	"blog-sp-kernelpanic/controller"
	"blog-sp-kernelpanic/database"
	"blog-sp-kernelpanic/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	database.InitDatabaseConnection(
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PORT"))
	if err != nil {
		panic(err)
	}
	database.DatabaseConnection.AutoMigrate(
		&model.Post{},
		&model.Section{},
		&model.HtmlSection{},
		&model.Tag{})

	app := fiber.New()

	app.Get("/", controller.DefaultController)
	app.Get("/post/create", controller.CreatePostController)
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics Page"}))

	err = app.Listen(fmt.Sprintf("%s:%s", "0.0.0.0", os.Getenv("APPLICATION_PORT")))
	if err != nil {
		panic(err)
	}
}
