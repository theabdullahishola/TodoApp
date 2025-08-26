package model

import (
	"context"
	"errors"

	"github.com/theabdullahishola/to-do/db"
	"github.com/theabdullahishola/to-do/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email    string             `json:"email" binding:"required"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Provider string             `json:"provider" bson:"provider"` // "local" or "google"
	GoogleID string             `json:"google_id,omitempty" bson:"google_id,omitempty"`
}

type Todo struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"userid" bson:"userID"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

func (u *User) NewUser() error {
	collection := db.GetCollection("golang_db", "users")

	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword
	u.Provider = "local"
	result, err := collection.InsertOne(context.Background(), u)
	if err != nil {
		return err
	}

	u.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func CreateOrGetGoogleUser(email, googleID string) (User, error) {
	collection := db.GetCollection("golang_db", "users")

	// Check if user already exists
	var user User
	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err == nil {
		if user.Provider != "google" {
			return User{}, errors.New("email already registered with password, please login normally")
	}
		return user, nil // already exists
	}
	
	
	// New user
	user = User{
		Email:    email,
		Provider: "google",
		GoogleID: googleID,
	}

	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return User{}, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func GetUserbyID(id string) (User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return User{}, err
	}
	collection := db.GetCollection("golang_db", "users")
	cursor := collection.FindOne(context.Background(), bson.M{"_id": objectID})
	var user User
	err = cursor.Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *User) ValidateCredentials() error {
	collection := db.GetCollection("golang_db", "users")
	cursor := collection.FindOne(context.Background(), bson.M{"email": u.Email})

	var user User
	err := cursor.Decode(&user)
	if err != nil {
		return err
	}

	u.ID = user.ID
	err = util.VerifyPassword(u.Password, user.Password)
	if err != nil {
		return errors.New("invalid Password")
	}
	return nil
}

func GetTodosByUser(userID primitive.ObjectID) ([]Todo, error) {
	collection := db.GetCollection("golang_db", "todo")
	var todos []Todo
	cursor, err := collection.Find(context.Background(), bson.M{"userID": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (t *Todo) CreateTodo() error {
	collection := db.GetCollection("golang_db", "todo")

	_, err := collection.InsertOne(context.Background(), t)
	if err != nil {
		return err
	}

	return nil
}
func GetTodobyID(ID primitive.ObjectID) (Todo, error) {
	collection := db.GetCollection("golang_db", "todo")
	cursor := collection.FindOne(context.Background(), bson.M{"_id": ID})
	var todo Todo
	err := cursor.Decode(&todo)
	if err != nil {
		return Todo{}, err
	}
	return todo, nil

}
func (t *Todo) UpdateTodo() error {

	collection := db.GetCollection("golang_db", "todo")

	filter := bson.M{"_id": t.ID}
	update := bson.M{"$set": bson.M{"completed": t.Completed}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todo) DeleteTodo() error {
	collection := db.GetCollection("golang_db", "todo")

	filter := bson.M{"_id": t.ID}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
