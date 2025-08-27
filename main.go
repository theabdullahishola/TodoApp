package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/theabdullahishola/to-do/db"
	"github.com/theabdullahishola/to-do/routes"
	"github.com/theabdullahishola/to-do/util"
)

func main() {
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal(err)
		}
	}

	util.InitGoogleOAuth()

	server := gin.Default()

	MONGO_URI := os.Getenv("MONGO_URI")
	if MONGO_URI == "" {
		log.Fatal("MONGO_URI is not set")
	}
	db.ConnectDB(MONGO_URI)
	defer db.Client.Disconnect(context.Background())

	PORT_NUMBER := os.Getenv("PORT")
	if PORT_NUMBER == "" {
		PORT_NUMBER = "8080"
	}
	
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("URL")}, // frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Register all API routes first
	routes.RegisterRoutes(server)

	if os.Getenv("ENV") == "production" {
		server.Static("/assets", "./client/dist/assets")
		server.NoRoute(func(c *gin.Context) {
			c.File("./client/dist/index.html")
		})
	}

	if err := server.Run(":" + PORT_NUMBER); err != nil {
		log.Fatal(err)
	}
}
