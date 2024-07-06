package controller

import (
	"go-language/config"
	"go-language/model"

	"github.com/gofiber/fiber/v2"
)

// CreateAuthor handles POST /api/authors
func CreateAuthor(c *fiber.Ctx) error {
	author := new(model.Author)
	if err := c.BodyParser(author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	if err := config.DB.Create(&author).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"author":  author,
	})
}

// GetAllAuthors handles GET /api/authors
func GetAllAuthors(c *fiber.Ctx) error {
	var authors []model.Author
	if err := config.DB.Where("del_flag = ?", false).Find(&authors).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"authors": authors,
	})
}

// GetAuthorByID handles GET /api/authors/:id
func GetAuthorByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var author model.Author
	if err := config.DB.Where("id = ? AND del_flag = ?", id, false).First(&author).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "Author not found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"author":  author,
	})
}

// UpdateAuthorByID handles PUT /api/authors/:id
func UpdateAuthorByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var author model.Author
	if err := config.DB.Where("id = ? AND del_flag = ?", id, false).First(&author).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "Author not found",
		})
	}

	// Store the old email address
	oldEmail := author.Email

	if err := c.BodyParser(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	if err := config.DB.Save(&author).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	// Check if the email has changed and prepare response message
	var message string
	if author.Email != oldEmail {
		message = "Email notification: Author's email has been changed"
	}

	return c.JSON(fiber.Map{
		"success": true,
		"author":  author,
		"message": message,
	})
}

// DeleteAuthorByID handles DELETE /api/authors/:id
func DeleteAuthorByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var author model.Author

	if err := config.DB.First(&author, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "Author not found",
		})
	}

	// Update the del_flag to true
	author.DelFlag = true
	if err := config.DB.Save(&author).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Author marked as deleted",
	})
}

