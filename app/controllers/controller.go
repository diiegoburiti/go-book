package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/diiegoburiti/go-book/app/configs"
	"github.com/diiegoburiti/go-book/app/models"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

var validate = validator.New()
var bookCollection *mongo.Collection = configs.GetCollection(configs.Db, "books")

func RegisterBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var book models.Book

	defer cancel()

	if err := c.BodyParser(&book); err != nil {
		c.Status(http.StatusBadRequest).JSON(BookResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationError := validate.Struct(&book); validationError != nil {
		return c.Status(http.StatusBadRequest).JSON(BookResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": validationError.Error}})
	}

	newBook := models.Book{
		ID:    primitive.NewObjectID(),
		Isbn:  book.Isbn,
		Title: book.Title,
	}

	result, err := bookCollection.InsertOne(ctx, newBook)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(BookResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": err.Error}})
	}

	return c.Status(http.StatusCreated).JSON(BookResponse{Status: http.StatusCreated, Message: "Success", Data: &fiber.Map{"data": result}})
}

func DeleteBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var bookId = c.Params("bookId")
	objectId, _ := primitive.ObjectIDFromHex(bookId)

	defer cancel()

	result, err := bookCollection.DeleteOne(ctx, bson.M{"id": objectId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(BookResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": err.Error}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(BookResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "Book not found"}})
	}

	return c.Status(http.StatusOK).JSON(BookResponse{Status: http.StatusCreated, Message: "Success", Data: &fiber.Map{"data": "Book deleted"}})
}

func FindBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var bookId = c.Params("bookId")
	objectId, _ := primitive.ObjectIDFromHex(bookId)
	var book models.Book

	defer cancel()

	err := bookCollection.FindOne(ctx, bson.M{"id": objectId}).Decode(&book)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(BookResponse{Status: http.StatusBadRequest, Message: "Error", Data: &fiber.Map{"data": err.Error}})
	}

	return c.Status(http.StatusOK).JSON(BookResponse{Status: http.StatusCreated, Message: "Success", Data: &fiber.Map{"data": book}})
}
