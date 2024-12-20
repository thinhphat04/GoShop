package routes

import (
	"API/internal/handlers"
	"API/internal/repositories"
	"API/internal/services"
	"API/pkg/database"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	// Kết nối tới cơ sở dữ liệu MySQL
	database.ConnectDatabase()

	// Tạo repository, service và handler
	productRepo := repositories.NewProductRepository()
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	router := gin.Default()

	// Cấu hình Trusted Proxies
	err := router.SetTrustedProxies([]string{"192.168.1.1", "127.0.0.1"})
	if err != nil {
		panic(err)
	}

	// Định nghĩa routes cho product
	api := router.Group("/api")
	{
		api.GET("/products", productHandler.GetProducts)    // Lấy danh sách sản phẩm
		api.POST("/products", productHandler.CreateProduct) // Tạo sản phẩm mới
	}

	return router
}
