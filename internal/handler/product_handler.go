package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ruanv123/api-go-crud/internal/model"
	"github.com/ruanv123/api-go-crud/internal/service"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// CreateProduct cria um novo produto
func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product model.Product

	// Parse do corpo da requisição para um modelo Produto
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Chama o serviço para salvar o produto
	if err := h.productService.CreateProduct(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create product"})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

// GetAllProducts retorna todos os produtos
func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	name := c.Query("name", "")      // Filtro de nome
	page := c.QueryInt("page", 1)    // Número da página (valor padrão: 1)
	limit := c.QueryInt("limit", 10) // Limite de itens por página (valor padrão: 10)

	products, total, err := h.productService.GetAllProducts(name, page, limit)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products"})
	}

	// Calcula o total de páginas
	totalPages := (total + int64(limit) - 1) / int64(limit)

	// Retorna os produtos com informações de paginação
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"products":    products,
		"total":       total,
		"totalPages":  totalPages,
		"currentPage": page,
	})
}

// GetProductByID retorna um produto pelo ID
func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

// UpdateProduct atualiza um produto existente
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Set the ID for the product to update
	product.ID = uint(id)

	if err := h.productService.UpdateProduct(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update product"})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

// DeleteProduct exclui um produto pelo ID
func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	if err := h.productService.DeleteProduct(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete product"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Product deleted"})
}
