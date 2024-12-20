# Sử dụng base image cho Golang
FROM golang:1.20

# Thiết lập thư mục làm việc trong container
WORKDIR /app

# Copy toàn bộ project vào container
COPY . .

# Tải các module cần thiết
RUN go mod tidy

# Build ứng dụng
RUN go build -o main cmd/api/main.go

# Expose cổng API
EXPOSE 8080

# Command để chạy ứng dụng
CMD ["./main"]