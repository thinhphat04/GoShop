package routes

import (
	"API/internal/handlers"
	"API/internal/repositories"
	"API/internal/services"
	"API/pkg/database"
	"API/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	// Kết nối tới cơ sở dữ liệu MySQL
	database.ConnectDatabase()
	// Tạo repository, service và handler cho product
	productRepo := repositories.NewProductRepository()
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)
	// Tạo repository, service và handler cho user
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router := gin.Default()
	// thêm middleware toàn cục
	router.Use(middlewares.LoggerMiddleware())

	// Cấu hình Trusted Proxies
	err := router.SetTrustedProxies([]string{"192.168.1.1", "127.0.0.1"})
	if err != nil {
		panic(err)
	}

	// Định nghĩa routes cho product
	api := router.Group("/api")
	{
		// product routes
		api.GET("/products", productHandler.GetProducts)    // Lấy danh sách sản phẩm
		api.POST("/products", productHandler.CreateProduct) // Tạo sản phẩm mới
		// user routes
		api.POST("/users", userHandler.CreateUser) // Tạo user mới
		api.GET("/users", userHandler.GetUsers)    // Lấy danh sách user
	}
	return router
}
