package handlers

import (
	"Scalable-Secure-Go-Web/internal/config"
	"Scalable-Secure-Go-Web/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
)

// GetAllProducts godoc
// @Summary Get all products with pagination
// @Description Retrieve a list of products with pagination and relations
// @Tags Products
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /products [get]
func GetAllProducts(c *fiber.Ctx) error {
	// Parse query parameters
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	var products []models.Product

	// Query products with related Category and Brand using GORM Preload
	err := config.DB.
		Preload("Category").
		Preload("Brand").
		Limit(limit).
		Offset(offset).
		Find(&products).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 500,
			Data:       nil,
			Message:    "Failed to fetch products",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.APIResponse{
		Status:     "success",
		StatusCode: 200,
		Data:       products,
		Message:    "Products fetched successfully",
	})
}

// GetProductByID godoc
// @Summary Get a single product by ID
// @Description Retrieve a product with its Category and Brand by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Router /products/{id} [get]
func GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	// Fetch product and preload Category and Brand
	if err := config.DB.Preload("Category").Preload("Brand").First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.APIResponse{
				Status:     "error",
				StatusCode: 404,
				Data:       nil,
				Message:    "Product not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 500,
			Data:       nil,
			Message:    "Error retrieving product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.APIResponse{
		Status:     "success",
		StatusCode: 200,
		Data:       product,
		Message:    "Product fetched successfully",
	})
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a product with Category and Brand references
// @Tags Products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product JSON"
// @Success 201 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /products [post]
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	// Parse JSON input
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 400,
			Data:       nil,
			Message:    "Invalid request body",
		})
	}

	// Validate input using validator package
	if err := validateProduct.Struct(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 400,
			Data:       nil,
			Message:    err.Error(),
		})
	}

	// Validate foreign keys: CategoryID and BrandID must exist
	if err := config.DB.First(&models.Category{}, product.CategoryID).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 400,
			Data:       nil,
			Message:    "Invalid CategoryID",
		})
	}
	if err := config.DB.First(&models.Brand{}, product.BrandID).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 400,
			Data:       nil,
			Message:    "Invalid BrandID",
		})
	}

	// Create product
	if err := config.DB.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 500,
			Data:       nil,
			Message:    "Failed to create product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(models.APIResponse{
		Status:     "success",
		StatusCode: 201,
		Data:       product,
		Message:    "Product created successfully",
	})
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Update a product by ID with new details
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Product JSON"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Router /products/{id} [put]
func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var existing models.Product

	// Fetch the product
	if err := config.DB.First(&existing, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 404,
			Data:       nil,
			Message:    "Product not found",
		})
	}

	var input models.Product
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 400,
			Data:       nil,
			Message:    "Invalid input",
		})
	}

	// Validate input
	if err := validateProduct.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 400,
			Data:       nil,
			Message:    err.Error(),
		})
	}

	// Check if referenced Category and Brand exist
	if err := config.DB.First(&models.Category{}, input.CategoryID).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 400,
			Data:       nil,
			Message:    "Invalid CategoryID",
		})
	}
	if err := config.DB.First(&models.Brand{}, input.BrandID).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 400,
			Data:       nil,
			Message:    "Invalid BrandID",
		})
	}

	// Update fields
	existing.Name = input.Name
	existing.Description = input.Description
	existing.Price = input.Price
	existing.CoverImage = input.CoverImage
	existing.CategoryID = input.CategoryID
	existing.BrandID = input.BrandID

	if err := config.DB.Save(&existing).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 500,
			Data:       nil,
			Message:    "Failed to update product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.APIResponse{
		Status:     "success",
		StatusCode: 200,
		Data:       existing,
		Message:    "Product updated successfully",
	})
}

// DeleteProduct godoc
// @Summary Delete a product by ID
// @Description Delete a product from the catalog by its ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 204 {object} nil
// @Failure 404 {object} models.APIResponse
// @Router /products/{id} [delete]
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	// Check if the product exists
	if err := config.DB.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 404,
			Data:       nil,
			Message:    "Product not found",
		})
	}

	// Delete product
	if err := config.DB.Delete(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 500,
			Data:       nil,
			Message:    "Failed to delete product",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
