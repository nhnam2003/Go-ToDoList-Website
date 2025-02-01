package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Hàm kết nối MongoDB
func MongoDB() {
	MONGODB_URL := os.Getenv("MONGODB_URL")
	if MONGODB_URL == "" {
		log.Fatal("❌ MONGODB_URL chưa được thiết lập")
	}

	// Cấu hình MongoDB client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MONGODB_URL).SetServerAPIOptions(serverAPI)

	// Kết nối đến MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal("❌ Lỗi kết nối MongoDB:", err)
	}

	// Kiểm tra kết nối
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		log.Fatal("❌ Không thể ping MongoDB:", err)
	}

	fmt.Println("✅ Đã kết nối thành công với MongoDB!")
}

// Hàm lấy collection cho cơ sở dữ liệu và collection cụ thể
func GetCollection(collectionName string) *mongo.Collection {
	var DBName = os.Getenv("DBName")

	if client == nil {
		log.Fatal("❌ MongoDB client chưa được khởi tạo!")
	}

	
	log.Printf("✅ Đang lấy collection %s từ database %s",DBName, collectionName)
	return client.Database(DBName).Collection(collectionName)
}

// Hàm đóng kết nối MongoDB khi ứng dụng dừng
func CloseDB() {
	if client != nil {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal("❌ Lỗi khi đóng kết nối MongoDB:", err)
		}
		fmt.Println("✅ Đã đóng kết nối MongoDB!")
	}
}
