package controllers

import (
	"strconv"

	database "github.com/codewithmujab/gocrud/config"
	"github.com/codewithmujab/gocrud/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// hello
func Hello(c *fiber.Ctx) error {
	return c.SendString("welcome fiber restapi mujab")
}

// addbook
func AddBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	//validation
	var validate = validator.New()
	if err := validate.Struct(book); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// end validate

	database.DBConn.Create(&book)
	return c.Status(200).JSON(book)
}

// getbookbyid
func GetBook(c *fiber.Ctx) error {
	books := []models.Book{}

	database.DBConn.First(&books, c.Params("id"))
	return c.Status(200).JSON(books)
}

// allbook
func AllBooks(c *fiber.Ctx) error {
	books := []models.Book{}

	database.DBConn.Find(&books)
	return c.Status(200).JSON(books)
}

// update
func Update(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		//400 bad request
		return c.Status(400).JSON(err.Error())
	}

	id, _ := strconv.Atoi(c.Params("id"))
	//ketika field lain tidak diupdate maka harusnya tidak empty

	// Update attributes dengan `map`
	database.DBConn.Model(&models.Book{}).Where("id = ?", id).Updates(map[string]interface{}{"title": book.Title, "author": book.Author})
	return c.Status(200).JSON("updated successfully")
}

// delete
func Destroy(c *fiber.Ctx) error {
	book := new(models.Book)

	id, _ := strconv.Atoi(c.Params("id"))

	database.DBConn.Where("id = ?", id).Delete(&book)
	return c.Status(200).JSON("deleted successfully")
}
