package handlers

import (
	"API/pkg/redis"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"API/internal/models"
	"API/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// GetProducts godoc
// @Summary      Retrieve a list of products
// @Description  Get all products from the database
// @Tags         Products
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Product
// @Failure      500  {object}  map[string]interface{}
// @Router       /products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	ctx := context.Background()
	cacheKey := "products_cache"

	// Lấy dữ liệu từ Redis
	start := time.Now()
	cachedProducts, err := redis.Client.Get(ctx, cacheKey).Result()
	if err == nil {
		// Deserialize dữ liệu Redis
		var products []models.Product
		json.Unmarshal([]byte(cachedProducts), &products)

		// Trả về dữ liệu từ Redis
		duration := time.Since(start)
		log.Printf("Redis cache hit, duration: %v\n", duration)
		c.JSON(http.StatusOK, gin.H{"data": products, "source": "cache"})
		return
	}

	// Nếu Redis cache miss, lấy từ cơ sở dữ liệu
	start = time.Now()
	products, err := h.productService.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
		return
	}

	// Lưu dữ liệu vào Redis
	productJSON, _ := json.Marshal(products)
	redis.Client.Set(ctx, cacheKey, productJSON, 10*time.Minute)

	// Trả về dữ liệu từ database
	duration := time.Since(start)
	log.Printf("Database query duration: %v\n", duration)
	c.JSON(http.StatusOK, gin.H{"data": products, "source": "database"})
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Add a new product to the database
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        product  body      models.Product  true  "Product details"
// @Success      201      {object}  models.Product
// @Failure      400      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := h.productService.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}
