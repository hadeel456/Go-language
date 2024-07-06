package controller

import (
	"go-language/config"
	"go-language/model"

	"github.com/gofiber/fiber/v2"
)

// CreateBook handles POST /api/books
func CreateBook(c *fiber.Ctx) error {
	book := new(model.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	if err := config.DB.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"book":    book,
	})
}

// GetAllBooks handles GET /api/books
func GetAllBooks(c *fiber.Ctx) error {
	var books []model.Book
	if err := config.DB.Where("del_flag = ?", false).Find(&books).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"books":   books,
	})
}

// GetBookByID handles GET /api/books/:id
func GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var book model.Book
	if err := config.DB.Where("id = ? AND del_flag = ?", id, false).First(&book).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "Book not found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"book":    book,
	})
}

// UpdateBookByID handles PUT /api/books/:id
func UpdateBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var book model.Book
	if err := config.DB.Where("id = ? AND del_flag = ?", id, false).First(&book).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "Book not found",
		})
	}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	if err := config.DB.Save(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"book":    book,
	})
}

// DeleteBookByID handles DELETE /api/books/:id
func DeleteBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var book model.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "Book not found",
		})
	}

	// Update the del_flag to true
	book.DelFlag = true
	if err := config.DB.Save(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Book marked as deleted",
	})
}
