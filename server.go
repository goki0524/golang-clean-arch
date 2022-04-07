package main

import (
	"github.com/goki0524/clean-arch/infrastructure/database"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("Error env file could not be read: %v", err)
	}
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	// DB Connect
	db, err := database.Connect()
	if err != nil {
		logrus.Infof("Error connecting DB: %v", err)
		// Heroku用 アプリの起動に合わせてDBが起動できないことがあるので再接続を試みる
		db, _ = database.Connect()
	}

	defer db.Close()
	// output sql query
	db.LogMode(true)

	// Routes
	r := Initialize(e, db)
	r.Init()

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
