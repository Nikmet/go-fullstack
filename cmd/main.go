package main

import (
	"go-fullstack/config"
	"go-fullstack/internal/home"
	"go-fullstack/internal/vacancy"
	"go-fullstack/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Congigs
	config.Init()
	// dbConf := config.NewDBConfig()

	logConf := config.NewLogConfig()
	customLogger := logger.NewLogger(logConf)

	// App
	app := fiber.New()

	//MiddleWare
	app.Use(recover.New(), fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))

	//Static
	app.Static("/public", "./public")

	// Handlers
	home.NewHandler(app, customLogger)
	vacancy.NewHanndler(app, customLogger)

	app.Listen(":3000")
}
