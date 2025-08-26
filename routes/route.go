package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/theabdullahishola/to-do/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	// Authentication routes
	server.POST("/api/auth/signup", signUp)       // create account
	server.POST("/api/auth/login", login)         // login and get tokens
	server.GET("/api/auth/refresh", refreshToken) // optional refresh route
	server.POST("/api/auth/logout", logout)
	server.POST("/api/auth/google", googleAuth)

	// Todo routes (protected)
	todoRoutes := server.Group("/api/todos")
	todoRoutes.Use(middleware.Authenticate)
	todoRoutes.GET("/", getTodos)         // fetch user’s todos
	todoRoutes.POST("/", createTodo)      // create new todo
	todoRoutes.PUT("/:id", updateTodo)    // update existing
	todoRoutes.DELETE("/:id", deleteTodo) // delete existing

}
