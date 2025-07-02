// filepath: taipei-day-trip-go-go/internal/utils/db.go
package utils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os" // <-- 【修正 1】: 引入 os 套件

	mysqlDriver "github.com/go-sql-driver/mysql" // 給底層驅動一個明確的別名，避免與 gorm 的驅動 Open 函數混淆
	// _ "github.com/go-sql-driver/mysql" // 這裡可以選擇性保留或移除，因為您明確呼叫了 mysqlDriver.RegisterTLSConfig

	gormMysql "gorm.io/driver/mysql" // GORM 的 MySQL 驅動
	"gorm.io/gorm"

	"taipei-day-trip-go-go/internal/models"
)

var Database *gorm.DB

func InitDB() error {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	caCertPath := os.Getenv("DB_SSL_CA")

	log.Printf("DB_HOST=%s, DB_PORT=%s, DB_USER=%s, DB_NAME=%s, CA=%s", dbHost, dbPort, dbUser, dbName, caCertPath)

	// 設定 CA 憑證
	// 建議將憑證路徑也放在環境變數中，或使用相對路徑，以提高移植性
	// caCertPath := os.Getenv("CA_CERT_PATH")
	// 如果不希望用環境變數，考慮使用相對路徑： "./ca.pem"
	// caCertPath := "/Users/ray/CSproject/台北旅遊一日遊網站(go language)/ca.pem" // 替換為您的 CA 憑證路徑

	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		return fmt.Errorf("無法讀取 CA 憑證 (%s): %w", caCertPath, err) // 加上路徑方便除錯
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		return fmt.Errorf("無法從 PEM 附加憑證到 RootCAs") // 錯誤處理更精確
	}

	// 註冊自定義 TLS 設定
	// 使用明確的別名 mysqlDriver
	err = mysqlDriver.RegisterTLSConfig("custom", &tls.Config{
		RootCAs: rootCertPool,
	})
	if err != nil {
		return fmt.Errorf("無法註冊 TLS 設定: %w", err)
	}
	log.Println("自定義 TLS 設定 'custom' 已註冊。") // 增加日誌確認

	// 建立資料庫連線字串
	// 確保 os.Getenv 獲取到的變數是正確設定的
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=custom",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	// 初始化資料庫連線
	// 【修正 2】: 使用 gormMysql.Open 而不是 mysql.Open
	Database, err = gorm.Open(gormMysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("資料庫連線失敗: %w", err)
	}

	log.Println("資料庫連線成功")

	// 自動建表
	models.Migrate(Database)

	return nil
}
