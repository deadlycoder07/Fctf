package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/deadlycoder07/Fctf/pkg/configs"
	"github.com/deadlycoder07/Fctf/pkg/middleware"
	"github.com/deadlycoder07/Fctf/pkg/routes"
	"github.com/deadlycoder07/Fctf/pkg/utils"
	"github.com/go-pg/pg/v10"

	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.AdminRoutes(app)
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// connection for database
	url, _ := utils.ConnectionURLBuilder("postgres")
	fmt.Print(url)
	start := time.Now()
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "admin",
		Database: "fctf",
		Addr:     url,
	})
	err := db.Ping(context.Background())
	elapsed := time.Since(start)
	fmt.Print("time taken to connect" + elapsed.String())
	if err != nil {
		fmt.Print(err)
	}

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}

func ConnectionURLBuilder(s string) {
	panic("unimplemented")
}
