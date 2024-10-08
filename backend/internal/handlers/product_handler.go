package handlers

import (
	"backend/internal/database"
	"backend/internal/models"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product
	database.DB.Preload("Category").Find(&products)

	// Adjust the image path to be a full URL
	for i, product := range products {
		products[i].ImagePath = fmt.Sprintf("%s/uploads/%s", c.BaseURL(), product.ImagePath)
	}

	return c.JSON(products)
}

func CreateProduct(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse form"})
	}

	var product models.Product
	product.Name = form.Value["name"][0]
	product.Description = form.Value["description"][0]
	product.Price, err = strconv.ParseFloat(form.Value["price"][0], 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid price"})
	}
	product.ExpiryDate, err = time.Parse("2006-01-02", form.Value["expirationDate"][0])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid expiration date"})
	}

	categoryID, err := strconv.Atoi(form.Value["category"][0])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid category ID"})
	}
	product.CategoryID = uint(categoryID)

	if image, ok := form.File["image"]; ok && len(image) > 0 {
		imagePath := fmt.Sprintf("images/%s", image[0].Filename)
		if err := c.SaveFile(image[0], fmt.Sprintf("./uploads/%s", imagePath)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not save image"})
		}
		product.ImagePath = imagePath
	}

	if err := database.DB.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create product"})
	}

	product.ImagePath = fmt.Sprintf("%s/uploads/%s", c.BaseURL(), product.ImagePath)

	return c.Status(fiber.StatusCreated).JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := database.DB.Preload("Category").First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	product.ImagePath = fmt.Sprintf("%s/uploads/%s", c.BaseURL(), product.ImagePath)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse form"})
	}

	product.Name = form.Value["name"][0]
	product.Description = form.Value["description"][0]
	product.Price, err = strconv.ParseFloat(form.Value["price"][0], 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid price"})
	}
	product.ExpiryDate, err = time.Parse("2006-01-02", form.Value["expirationDate"][0])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid expiration date"})
	}

	categoryID, err := strconv.Atoi(form.Value["category"][0])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid category ID"})
	}
	product.CategoryID = uint(categoryID)

	if image, ok := form.File["image"]; ok && len(image) > 0 {
		imagePath := fmt.Sprintf("images/%s", image[0].Filename)
		if err := c.SaveFile(image[0], fmt.Sprintf("./uploads/%s", imagePath)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not save image"})
		}
		product.ImagePath = imagePath
	}

	if err := database.DB.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update product"})
	}

	product.ImagePath = fmt.Sprintf("%s/uploads/%s", c.BaseURL(), product.ImagePath)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete product"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
