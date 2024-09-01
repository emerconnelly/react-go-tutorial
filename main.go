package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        int    `json:"_id" bson:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
} // empty default value: {id:0,completed:false,body:""}

var collection *mongo.Collection // imports mongo package and bson data type

func main() {
	// hey
	fmt.Println("hello world")

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// connect to MongoDB Atlas
	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background()) // disconnect if some fancy shit happens idk this doesn't make sense yet

	// test MongoDB Atlas connection even though we literally just connected to it wtf
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB Atlas")

	// load the entire collection from the database
	collection = client.Database("golang_db").Collection("todos")

	// app routers
	app := fiber.New()
	app.Get("/api/todos", getTodos)
	// app.Post("/api/todos", createTodo)
	// app.Patch("/api/todos/:id", updateTodo)
	// app.Delete("/api/todos/:id", deleteTodo)

	// port listener
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func getTodos(c *fiber.Ctx) error {
	var todos []Todo

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	defer cursor.Close(context.Background()) // another fancy thingy

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}

	return c.JSON(todos)
}

// func createTodo(c *fiber.Ctx) error {}
// func updateTodo(c *fiber.Ctx) error {}
// func deleteTodo(c *fiber.Ctx) error {}
