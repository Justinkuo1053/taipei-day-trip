package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

type AttractionRaw struct {
	ID          int64   `json:"_id"`
	Name        string  `json:"name"`
	Category    string  `json:"CAT"`
	Description string  `json:"description"`
	Address     string  `json:"address"`
	Transport   string  `json:"direction"`
	MRT         string  `json:"MRT"`
	Latitude    float64 `json:"latitude,string"`
	Longitude   float64 `json:"longitude,string"`
	File        string  `json:"file"`
}

type AttractionsResult struct {
	Results []AttractionRaw `json:"results"`
}

type AttractionsData struct {
	Result AttractionsResult `json:"result"`
}

func splitImageURLs(urls string) []string {
	var result []string
	re := regexp.MustCompile(`https.*?\.(jpg|JPG|png|PNG)`)
	matches := re.FindAllString(urls, -1)
	for _, m := range matches {
		result = append(result, m)
	}
	return result
}

func main() {
	// 1. 讀取環境變數
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// 組合 host:port
	hostPort := dbHost
	if dbPort != "" {
		hostPort = fmt.Sprintf("%s:%s", dbHost, dbPort)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, hostPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("資料庫連線失敗:", err)
	}
	defer db.Close()

	// 2. 讀取 JSON 檔案
	jsonFile, err := ioutil.ReadFile("taipei-attractions.json")
	if err != nil {
		log.Fatal("讀取 JSON 失敗:", err)
	}

	var data AttractionsData
	if err := json.Unmarshal(jsonFile, &data); err != nil {
		log.Fatal("解析 JSON 失敗:", err)
	}

	// 3. 批次寫入（可用 transaction）
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("啟動 transaction 失敗:", err)
	}

	attrStmt, err := tx.Prepare(`INSERT INTO attractions (id, name, category, description, address, transport, mrt, lat, lng) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal("預備 attraction SQL 失敗:", err)
	}
	defer attrStmt.Close()

	imgStmt, err := tx.Prepare(`INSERT INTO images (url, attraction_id) VALUES (?, ?)`)
	if err != nil {
		log.Fatal("預備 image SQL 失敗:", err)
	}
	defer imgStmt.Close()

	for _, attr := range data.Result.Results {
		_, err := attrStmt.Exec(attr.ID, attr.Name, attr.Category, attr.Description, attr.Address, attr.Transport, attr.MRT, attr.Latitude, attr.Longitude)
		if err != nil {
			log.Printf("寫入 attraction 失敗: %v\n", err)
			tx.Rollback()
			return
		}
		urls := splitImageURLs(attr.File)
		for _, url := range urls {
			_, err := imgStmt.Exec(url, attr.ID)
			if err != nil {
				log.Printf("寫入 image 失敗: %v\n", err)
				tx.Rollback()
				return
			}
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal("commit 失敗:", err)
	}

	fmt.Println("資料匯入完成！")
}
