package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/theabdullahishola/to-do/db"
	"github.com/theabdullahishola/to-do/routes"
	"github.com/theabdullahishola/to-do/util"
)

func main() {

	if os.Getenv("ENV") !="production" {
		err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
		return
	}
	}
	
	util.InitGoogleOAuth()
	server := gin.Default()
	MONGO_URI := os.Getenv("MONGO_URI")
	if MONGO_URI == "" {
		log.Fatal("MONGO_URI is not set in .env file")
	}
	db.ConnectDB(MONGO_URI)
	PORT_NUMBER := os.Getenv("PORT")

	if PORT_NUMBER == "" {
		PORT_NUMBER = "8080"
	}
	
	if os.Getenv("ENV") == "production" {
		// Serve static assets (e.g. JS, CSS)
		server.Static("/assets", "./client/dist/assets")

		// Serve index.html for all other non-API routes (React fallback)
		server.NoRoute(func(c *gin.Context) {
			c.File("./client/dist/index.html")
		})
	}
	// server.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:5173"}, // frontend origin
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true, // allow cookies!
	// 	MaxAge:           12 * time.Hour,
	// }))
	routes.RegisterRoutes(server)
	err := server.Run(":" + PORT_NUMBER)

	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Client.Disconnect(context.Background())
}
