# Go-ToDoList

Go-ToDoList là một ứng dụng quản lý công việc (ToDo List) được xây dựng bằng **Go** với **Fiber** cho backend và **Vite + TypeScript** cho frontend, sử dụng **Chakra UI** và **TanStack Query** để quản lý dữ liệu.

## Nội dung

- [Cài đặt](#cài-đặt)
- [Chạy ứng dụng](#chạy-ứng-dụng)
- [Cấu trúc thư mục](#cấu-trúc-thư-mục)
- [API](#api)
- [Frontend](#frontend)
- [Cảm ơn](#cảm-ơn)

## Cài đặt

### Backend (Go)

1. Cài đặt Go (nếu chưa cài): [Hướng dẫn cài đặt Go](https://golang.org/doc/install).
2. Clone repository này về máy:

   ```bash
   git clone https://github.com/DaiNef163/Go-ToDoList.git
   cd Go-ToDoList
Tạo file .env trong thư mục gốc của dự án và thêm các biến môi trường cần thiết:

bash
Copy
Edit
MONGODB_URL=mongodb://<your_mongo_url>
PORT=3000
Cài đặt các dependencies:

bash
Copy
Edit
go mod init github.com/DaiNef163/Go-ToDoList
go get github.com/gofiber/fiber/v2
go install github.com/air-verse/air@latest
go get github.com/joho/godotenv
Chạy ứng dụng Go (backend):

bash
Copy
Edit
go run cmd/main.go
Frontend (Vite + TypeScript + Chakra UI)
Di chuyển đến thư mục frontend:

bash
Copy
Edit
cd frontend
Cài đặt các dependencies:

bash
Copy
Edit
npm install
Chạy ứng dụng frontend:

bash
Copy
Edit
npm run dev
Mở trình duyệt và truy cập vào http://localhost:5173.
Fiber: Framework Go cho việc xây dựng ứng dụng web hiệu suất cao.
MongoDB: Cơ sở dữ liệu NoSQL được sử dụng để lưu trữ dữ liệu.
Chakra UI: Bộ công cụ UI mạnh mẽ cho React.
TanStack Query: Thư viện giúp quản lý dữ liệu và trạng thái trong frontend.
