package routes

import (
	"go-language/controller"

	"github.com/gofiber/fiber/v2"
)

func Book(app *fiber.App) {
	app.Post("/api/books", controller.CreateBook)
	app.Get("/api/books", controller.GetAllBooks)
	app.Get("/api/books/:id", controller.GetBookByID)
	app.Put("/api/books/:id", controller.UpdateBookByID)
	app.Delete("/api/books/:id", controller.DeleteBookByID)
}

func Author(app *fiber.App) {
	app.Post("/api/authors", controller.CreateAuthor)
	app.Get("/api/authors", controller.GetAllAuthors)
	app.Get("/api/authors/:id", controller.GetAuthorByID)
	app.Put("/api/authors/:id", controller.UpdateAuthorByID)
	app.Delete("/api/authors/:id", controller.DeleteAuthorByID)
}
