# Taipei Day Trip

**台北一日遊**是一個電商型網站，提供台北景點搜尋、行程預定與線上付款功能。

## 特色功能

- **景點搜尋**：可依景點名稱、捷運站名稱或關鍵字查詢詳細資訊
- **無限滾動**：每頁顯示 12 筆資料，自動載入更多景點
- **會員系統**：註冊、登入、登出功能
- **行程預定**：可於景點頁預定行程
- **線上付款**：支援信用卡付款（TapPay，開發中）
- **圖片預載與動畫**：提升瀏覽體驗

## 技術架構

- **前端**：HTML、CSS、JavaScript
- **後端**：Go (Gin 框架)
- **資料庫**：MySQL（Aiven）
- **驗證**：JWT
- **API 設計**：RESTful、MVC 架構
- **第三方金流**：TapPay SDK（開發中）
- **版本控制**：Git、GitHub，遵循 Git Flow 流程

## 系統架構圖



## 資料庫設計

| Table      | 說明           | 狀態     |
|------------|----------------|----------|
| attractions| 景點資訊       | 已完成   |
| image      | 景點圖片       | 已完成   |
| user       | 使用者資訊     | 已完成   |
| booking    | 預定資訊(購物車)| 已完成   |
| purchase   | 訂單資訊       | 開發中   |
| payment    | 付款資訊       | 開發中   |



## 使用說明

1. 安裝 [Go 1.20+](https://golang.org/) 與 [MySQL 8+](https://www.mysql.com/)
2. 下載專案
   ```bash
   git clone https://github.com/xxx/taipei-day-trip.git
   cd taipei-day-trip
   ```
3. 複製並編輯環境變數
   ```bash
   cp .env.example .env
   # 編輯 .env 填入資料庫連線資訊
   ```
4. 啟動後端服務
   ```bash
   go run main.go
   ```
5. 前端請直接開啟 `index.html` 或依需求部署

## API 文件

- 主要 API 皆採 RESTful 設計，詳見 `/docs` 或 Swagger 文件（建議補充）

## 測試帳號

- 帳號：`test@example.com`
- 密碼：`test1234`

## 已知問題與待辦

- TapPay 金流、訂單與付款功能尚在開發中
- API 文件持續補充中

## 版本分支
- **main**：Go-Gin + MySQL (Aiven)