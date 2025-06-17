package main

import (
	"log"
	"qr-backend/database"
	"qr-backend/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		log.Fatal("データベース接続に失敗しました:", err)
	}
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	qrHandler := handlers.NewQRHandler(db)

	api := e.Group("/api")
	api.GET("/health", qrHandler.HealthCheck)
	api.POST("/qr-code/generate", qrHandler.GenerateQRCode)

	log.Println("サーバーを開始します: http://localhost:8080")
	log.Fatal(e.Start(":8080"))
}
