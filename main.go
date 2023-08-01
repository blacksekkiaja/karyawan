package main

import (
	"karyawan/internal/app/controllers"
	"karyawan/internal/app/models"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "karyawan/docs"
)

// @title Data Karyawan API
// @version 1.0
// @description API untuk CRUD data karyawan
// @host localhost:8080
// @BasePath /
func main() {
	loadEnv()
	e := echo.New()

	// Middleware
	e.Use(middleware.CORS()) //untuk menghendle swaager di local agar bisa berjalan
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize database
	models.InitDB()
	models.AutoMigrate()

	// Routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/karyawan", controllers.CreateKaryawan)
	e.GET("/karyawan", controllers.GetKaryawanList)
	e.GET("/karyawan/:id", controllers.GetKaryawanByID)
	e.PUT("/karyawan/:id", controllers.UpdateKaryawan)
	e.DELETE("/karyawan/:id", controllers.DeleteKaryawan)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
