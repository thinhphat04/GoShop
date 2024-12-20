package main

import (
	"log"

	"API/config" // Đường dẫn gói config của bạn
	"API/routes" // Import file routes bạn vừa định nghĩa
)

func main() {
	// Load cấu hình từ file .env
	config.Load()

	// Khởi tạo router Gin
	router := routes.SetupRoutes()

	// Chạy server
	log.Println("Server is running on port 8080...")
	router.Run(":8080")
}
