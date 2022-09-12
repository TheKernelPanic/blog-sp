package main

import (
	"blog-sp-kernelpanic/controller"
	"blog-sp-kernelpanic/database"
	"blog-sp-kernelpanic/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/joho/godotenv"
	"os"
	"time"
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

	outputLoggerFile, err := os.OpenFile(
		fmt.Sprintf("%s/app.log", os.Getenv("LOGGER_OUTPUT_DIRECTORY")),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0664)
	if err != nil {
		panic(err)
	}
	defer outputLoggerFile.Close()

	app.Use(logger.New(logger.Config{
		Format:     "${time}: ${ip} - ${status} - ${method} ${path}\n",
		TimeFormat: time.RFC3339,
		TimeZone:   "UTC",
		Output:     outputLoggerFile,
	}))

	app.Get("/", controller.DefaultController)
	app.Post("/post/create", controller.CreatePostController)
	app.Get("/post/listing", controller.ListingGetController)
	app.Get("/post/:slug", controller.SlugGetController)
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics Page"}))

	err = app.Listen(fmt.Sprintf("%s:%s", os.Getenv("APPLICATION_HOST"), os.Getenv("APPLICATION_PORT")))
	if err != nil {
		panic(err)
	}
}
