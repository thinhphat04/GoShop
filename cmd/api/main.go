package main

import (
	"log"

	"API/config" // Đường dẫn gói config của bạn
	_ "API/docs" // Import tài liệu Swagger đã sinh
	"API/routes" // Import file routes
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API Documentation
// @version         1.0
// @description     API for managing users, products, orders, and carts.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Support Team
// @contact.url    http://www.swagger.io/support
// @contact.email  support@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api
func main() {
	// Load cấu hình từ file .env
	config.Load()

	// Khởi tạo router Gin
	router := routes.SetupRoutes()

	// Thêm Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Chạy server
	log.Println("Server is running on port 8080...")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
