package routes

import (
	"github.com/diiegoburiti/go-book/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Post("/register-book", controllers.RegisterBook)
	app.Delete("/delete-book/:id", controllers.DeleteBook)

}
