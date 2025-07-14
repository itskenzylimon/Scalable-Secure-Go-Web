package handlers

import (
	"Scalable-Secure-Go-Web/internal/config"
	"Scalable-Secure-Go-Web/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAllBrands godoc
// @Summary Get all brands
// @Description Retrieve all brands
// @Tags Brands
// @Accept json
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /brands [get]
func GetAllBrands(c *fiber.Ctx) error {
	var brands []models.Brand

	if err := config.DB.Find(&brands).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 500,
			Data:       nil,
			Message:    "Failed to fetch brands",
		})
	}

	return c.JSON(models.APIResponse{
		Status:     "success",
		StatusCode: 200,
		Data:       brands,
		Message:    "Brands retrieved successfully",
	})
}

// GetBrandByID godoc
// @Summary Get brand by ID
// @Description Retrieve a single brand by ID
// @Tags Brands
// @Accept json
// @Produce json
// @Param id path int true "Brand ID"
// @Success 200 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Router /brands/{id} [get]
func GetBrandByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var brand models.Brand

	if err := config.DB.First(&brand, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(models.APIResponse{
				Status:     "error",
				StatusCode: 404,
				Data:       nil,
				Message:    "Brand not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 500,
			Data:       nil,
			Message:    "Error retrieving brand",
		})
	}

	return c.JSON(models.APIResponse{
		Status:     "success",
		StatusCode: 200,
		Data:       brand,
		Message:    "Brand retrieved successfully",
	})
}

// CreateBrand godoc
// @Summary Create a new brand
// @Description Create a brand entry
// @Tags Brands
// @Accept json
// @Produce json
// @Param brand body models.Brand true "Brand JSON"
// @Success 201 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /brands [post]
func CreateBrand(c *fiber.Ctx) error {
	var brand models.Brand

	// Parse body
	if err := c.BodyParser(&brand); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 400,
			Data:       nil,
			Message:    "Invalid request body",
		})
	}

	// Validate input
	if err := validateBrand.Struct(&brand); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 400,
			Data:       nil,
			Message:    err.Error(),
		})
	}

	// Insert into DB
	if err := config.DB.Create(&brand).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 500,
			Data:       nil,
			Message:    "Failed to create brand",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(models.APIResponse{
		Status:     "success",
		StatusCode: 201,
		Data:       brand,
		Message:    "Brand created successfully",
	})
}

// UpdateBrand godoc
// @Summary Update a brand by ID
// @Description Update an existing brand's data
// @Tags Brands
// @Accept json
// @Produce json
// @Param id path int true "Brand ID"
// @Param brand body models.Brand true "Brand JSON"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Router /brands/{id} [put]
func UpdateBrand(c *fiber.Ctx) error {
	id := c.Params("id")
	var existing models.Brand

	// Check existence
	if err := config.DB.First(&existing, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 404,
			Data:       nil,
			Message:    "Brand not found",
		})
	}

	var input models.Brand
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 400,
			Data:       nil,
			Message:    "Invalid input",
		})
	}

	if err := validateBrand.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 400,
			Data:       nil,
			Message:    err.Error(),
		})
	}

	// Update fields
	existing.Name = input.Name
	existing.CoverImage = input.CoverImage

	if err := config.DB.Save(&existing).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 500,
			Data:       nil,
			Message:    "Failed to update brand",
		})
	}

	return c.JSON(models.APIResponse{
		Status:     "success",
		StatusCode: 200,
		Data:       existing,
		Message:    "Brand updated successfully",
	})
}

// DeleteBrand godoc
// @Summary Delete a brand
// @Description Remove a brand by ID
// @Tags Brands
// @Accept json
// @Produce json
// @Param id path int true "Brand ID"
// @Success 204
// @Failure 404 {object} models.APIResponse
// @Router /brands/{id} [delete]
func DeleteBrand(c *fiber.Ctx) error {
	id := c.Params("id")
	var brand models.Brand

	if err := config.DB.First(&brand, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 404,
			Data:       nil,
			Message:    "Brand not found",
		})
	}

	if err := config.DB.Delete(&brand).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Status:     "error",
			StatusCode: 500,
			Data:       nil,
			Message:    "Failed to delete brand",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
