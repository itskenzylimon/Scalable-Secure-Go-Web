package handlers

import (
	"Scalable-Secure-Go-Web/internal/config"
	"Scalable-Secure-Go-Web/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAllCategories godoc
// @Summary Get all categories
// @Description Retrieve a list of all product categories
// @Tags Categories
// @Accept json
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /categories [get]
func GetAllCategories(c *fiber.Ctx) error {
	var categories []models.Category

	if err := config.DB.Find(&categories).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 500,
			Data:       nil,
			Message:    "Failed to fetch categories",
		})
	}

	return c.JSON(models.APIResponse{
		Status:     "success",
		StatusCode: 200,
		Data:       categories,
		Message:    "Categories retrieved successfully",
	})
}

// GetCategoryByID godoc
// @Summary Get category by ID
// @Description Retrieve a single category by its ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Router /categories/{id} [get]
func GetCategoryByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category

	if err := config.DB.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.APIResponse{
				Status:     "error",
				StatusCode: 404,
				Data:       nil,
				Message:    "Category not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 500,
			Data:       nil,
			Message:    "Error retrieving category",
		})
	}

	return c.JSON(models.APIResponse{
		Status:     "success",
		StatusCode: 200,
		Data:       category,
		Message:    "Category retrieved successfully",
	})
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a product category
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body models.Category true "Category JSON"
// @Success 201 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /categories [post]
func CreateCategory(c *fiber.Ctx) error {
	var category models.Category

	// Parse JSON body
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 400,
			Data:       nil,
			Message:    "Invalid request body",
		})
	}

	// Validate input
	if err := validateCategory.Struct(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 400,
			Data:       nil,
			Message:    err.Error(),
		})
	}

	// Insert category into DB
	if err := config.DB.Create(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 500,
			Data:       nil,
			Message:    "Failed to create category",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(models.APIResponse{
		Status:     "success",
		StatusCode: 201,
		Data:       category,
		Message:    "Category created successfully",
	})
}

// UpdateCategory godoc
// @Summary Update a category by ID
// @Description Update an existing product category
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body models.Category true "Category JSON"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Router /categories/{id} [put]
func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var existing models.Category

	if err := config.DB.First(&existing, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 404,
			Data:       nil,
			Message:    "Category not found",
		})
	}

	var input models.Category
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 400,
			Data:       nil,
			Message:    "Invalid input",
		})
	}

	if err := validateCategory.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 400,
			Data:       nil,
			Message:    err.Error(),
		})
	}

	existing.Title = input.Title
	existing.CoverImage = input.CoverImage

	if err := config.DB.Save(&existing).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 500,
			Data:       nil,
			Message:    "Failed to update category",
		})
	}

	return c.JSON(models.APIResponse{
		Status:     "success",
		StatusCode: 200,
		Data:       existing,
		Message:    "Category updated successfully",
	})
}

// DeleteCategory godoc
// @Summary Delete a category by ID
// @Description Delete a product category
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 204
// @Failure 404 {object} models.APIResponse
// @Router /categories/{id} [delete]
func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category

	if err := config.DB.First(&category, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 404,
			Data:       nil,
			Message:    "Category not found",
		})
	}

	if err := config.DB.Delete(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 500,
			Data:       nil,
			Message:    "Failed to delete category",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
