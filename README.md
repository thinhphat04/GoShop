#  GoShop E-Commerce API
GoShop là một API RESTful được xây dựng bằng ngôn ngữ Go và sử dụng MySQL làm cơ sở dữ liệu chính. API này cung cấp các chức năng cơ bản cần thiết để xây dựng một hệ thống e-commerce, bao gồm các chức năng như quản lý sản phẩm, người dùng, đơn hàng, và giỏ hàng.
## Features
- **Quản lý người dùng:** Endpoint cho đăng ký, đăng nhập và thông tin người dùng.
- **Quản lý sản phẩm:** CRUD sản phẩm (tạo, đọc, cập nhật, xóa).
- **Giỏ hàng & Đơn hàng:** Thêm sản phẩm vào giỏ hàng, tạo đơn hàng, và quản lý đơn hàng.
- **Authentication:** Sử dụng JWT để xác thực.
- **Môi trường phát triển:** Hỗ trợ Docker để triển khai môi trường nhanh chóng.
## Requirements
### **Phần mềm yêu cầu**
1. **Go (>= v1.23):** Để biên dịch và chạy ứng dụng.
2. **MySQL (>= v8.0):** Cơ sở dữ liệu chính cho ứng dụng.
3. **Docker + Docker Compose:** Để dễ dàng triển khai.
## Getting Started
Dưới đây là hướng dẫn để cài đặt và khởi chạy API.
### **1. Clone dự án**
``` bash
git clone https://github.com/your-repository-url/goshop.git
cd goshop
```
### **2. Cấu hình môi trường**
Tạo file `.env` hoặc chỉnh sửa các biến môi trường trong Docker Compose để khớp với nhu cầu của bạn.
Mẫu file `.env`:
``` env
MYSQL_DSN=user:password@tcp(mysql:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local
JWT_SECRET=your_jwt_secret_key
APP_PORT=8080
```
### **3. Sử dụng Docker để khởi chạy ứng dụng**
Trong file `docker-compose.yml`, môi trường đã được định cấu hình với:
- **API Service:** Ứng dụng Go được gọi qua cổng 8080.
- **MySQL Service:** Chạy MySQL với cấu hình sẵn.

Khởi chạy với lệnh:
``` bash
docker-compose up --build
```
### **4. Truy cập ứng dụng**
Sau khi khởi chạy thành công:
- **API Endpoint:** [http://localhost:8080]()
