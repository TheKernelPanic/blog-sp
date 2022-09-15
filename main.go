package main

import (
	"blog-sp-kernelpanic/controller"
	"blog-sp-kernelpanic/database"
	"blog-sp-kernelpanic/model"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	controller.JWTPrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	controller.UploadsDirectory = os.Getenv("UPLOADS_DIRECTORY")

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
		&model.TextSection{},
		&model.FileUploaded{},
		&model.FileSection{},
		&model.Category{})

	app := fiber.New()

	outputLoggerFile, err := os.OpenFile(
		fmt.Sprintf("%s/app.log", os.Getenv("LOGGER_OUTPUT_DIRECTORY")),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0664)
	if err != nil {
		panic(err)
	}
	defer outputLoggerFile.Close()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(logger.New(logger.Config{
		Format:     "${time}: ${ip} - ${status} - ${method} ${path}\n",
		TimeFormat: time.RFC3339,
		TimeZone:   "UTC",
		Output:     outputLoggerFile,
	}))

	app.Post("/auth", controller.Authentication)

	jwtMiddlware := jwtware.New(jwtware.Config{
		SigningMethod: "RS256",
		SigningKey:    controller.JWTPrivateKey.Public()})

	postGroup := app.Group("/post").Use(jwtMiddlware)
	postGroup.Post("/create", controller.CreatePostController)
	postGroup.Get("/listing", controller.ListingGetController)
	postGroup.Get("/:slug", controller.SlugGetController)

	filesGroup := app.Group("/files").Use(jwtMiddlware)
	filesGroup.Post("/image/upload", controller.UploadImagePostController)

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics Page"})).Use(jwtMiddlware)

	err = app.Listen(fmt.Sprintf("%s:%s", os.Getenv("APPLICATION_HOST"), os.Getenv("APPLICATION_PORT")))
	if err != nil {
		panic(err)
	}
}
