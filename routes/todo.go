package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theabdullahishola/to-do/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getTodos(c *gin.Context) {

	userId := c.GetString("userID")
	objectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cant get data"})
		return
	}
	todos, err := model.GetTodosByUser(objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cant get data"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	userId := c.GetString("userID")
	objectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cant get data"})
		return
	}
	var todo model.Todo
	err = c.ShouldBind(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cant parse data"})
		return
	}
	todo.UserID = objectID
	err = todo.CreateTodo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cant post data"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "success"})

}

func updateTodo(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cant parse id"})
		return
	}

	userId := c.GetString("userID")
	objectUserID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cant get data"})
		return
	}
	todo, err := model.GetTodobyID(objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cant get todo"})
		return
	}

	if todo.UserID != objectUserID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	var updatedTodo model.Todo

	err = c.ShouldBindJSON(&updatedTodo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cant parse data"})
		return
	}
	updatedTodo.ID = objectID
	err = updatedTodo.UpdateTodo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cant update data"})
		return
	}
	update, err := model.GetTodobyID(objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cant get todo"})
		return
	}
	c.JSON(http.StatusOK, update)

}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cant parse id"})
		return
	}

	userId := c.GetString("userID")
	objectUserID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cant get data"})
		return
	}
	todo, err := model.GetTodobyID(objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cant get todo"})
		return
	}

	if todo.UserID != objectUserID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	err = todo.DeleteTodo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cant delete data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
