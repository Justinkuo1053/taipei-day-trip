package main

import (
	"log"
	"os"
	"taipei-day-trip-go-go/internal/handlers"
	"taipei-day-trip-go-go/internal/repositories"
	"taipei-day-trip-go-go/internal/services"
	"taipei-day-trip-go-go/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 載入 .env
	if err := godotenv.Load(); err != nil {
		log.Println(".env 載入失敗，將使用系統環境變數")
	}

	// 初始化資料庫
	if err := utils.InitDB(); err != nil {
		log.Fatalf("資料庫初始化失敗: %v", err)
	}

	// 初始化 Repository
	attractionRepo := repositories.NewAttractionRepository(utils.Database)
	bookingRepo := repositories.BookingRepository{DB: utils.Database}
	userRepo := repositories.NewUserRepository(utils.Database)

	// 初始化 Service
	attractionService := services.NewAttractionService(attractionRepo)
	bookingService := services.NewBookingService(&bookingRepo)
	jwtSecret := os.Getenv("JWT_SECRET")
	userService := services.NewUserService(userRepo, jwtSecret)

	// 初始化 Handler
	attractionHandler := handlers.NewAttractionHandler(attractionService)
	bookingHandler := handlers.NewBookingHandler(bookingService)
	userHandler := handlers.NewUserHandler(userService)

	// 初始化路由
	r := gin.Default()
	// 設定路由
	handlers.RegisterRoutes(r, attractionHandler, bookingHandler, userHandler, userService)
	// r.GET("/attractions/:id", attractionHandler.GetAttractionByID) // 暫時註解
	// r.POST("/orders", orderHandler.CreateOrder)                    // 暫時註解

	// 印出環境變數方便 debug
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	caCertPath := os.Getenv("DB_SSL_CA")
	log.Printf("DB_HOST=%s, DB_PORT=%s, DB_USER=%s, DB_NAME=%s, CA=%s", dbHost, dbPort, dbUser, dbName, caCertPath)

	// 設定 Gin 靜態檔案服務，讓 /static/* 路徑可以正確回傳前端靜態資源
	r.Static("/static", "./static") // 讓 /static/xxx 會對應到 static 資料夾

	// 設定前端主要頁面路由，讓使用者直接輸入網址也能正確顯示頁面
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})
	r.GET("/attraction/:id", func(c *gin.Context) {
		c.File("./static/attraction.html")
	})
	r.GET("/booking", func(c *gin.Context) {
		c.File("./static/booking.html")
	})
	r.GET("/thankyou", func(c *gin.Context) {
		c.File("./static/thankyou.html")
	})

	// 啟動伺服器
	r.Run(":8080")
}
