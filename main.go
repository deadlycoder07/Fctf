package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/deadlycoder07/Fctf/app/models"
	"github.com/deadlycoder07/Fctf/pkg/configs"
	"github.com/deadlycoder07/Fctf/pkg/middleware"
	"github.com/deadlycoder07/Fctf/pkg/routes"
	"github.com/deadlycoder07/Fctf/pkg/utils"

	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func TestFunc() {
	// connection for database
	start := time.Now()
	db := utils.GetPostgresConnection()
	data := models.Challenges{
		Name:        "Test",
		Description: "test description",
		MaxAttempts: 2,
		Value:       120,
		Category:    "pwn",
		Type:        "visible",
		State:       false,
		Files: []*models.Files{
			&models.Files{
				Type:     "crypto",
				Location: "dubai",
			},
		},
	}
	if err := db.Ping(context.Background()); err != nil {
		fmt.Print("\n Failed to connect to database")
	}
	_, err := db.Model(&data).Insert()
	for _, file := range data.Files {
		file.ChallengesId = 9
		_, e := db.Model(file).Insert()
		if e != nil {
			fmt.Print("\n", e)
		}
	}
	// if e := db.Model(&models.Challenges{}).DropTable(&orm.DropTableOptions{}); e != nil {
	// 	fmt.Print("\nfailed ")
	// }
	// if e := db.Model(&models.Files{}).DropTable(&orm.DropTableOptions{}); e != nil {
	// 	fmt.Print("\nfailed ")
	// }

	// err1 := db.Model(&models.Challenges{}).CreateTable(&orm.CreateTableOptions{Varchar: 90})
	// err2 := db.Model(&models.Files{}).CreateTable(&orm.CreateTableOptions{Varchar: 90})
	// if err1 != nil || err2 != nil {
	// 	fmt.Print("failed to create")
	// }
	chall := new(models.Challenges)
	err = db.Model(chall).Relation("Files").Where("id = ?", 9).Select()
	if err == nil {
		fmt.Print(chall.Files[1])
	}
	elapsed := time.Since(start)
	fmt.Print("\ntime taken to connect" + elapsed.String())
	if err != nil {
		fmt.Print(err)
	}
}

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

	TestFunc()
	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
