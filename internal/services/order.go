package services

import (
	"fmt"
	"math/rand"
	"taipei-day-trip-go-go/internal/interfaces"
	"taipei-day-trip-go-go/internal/models"
	"time"
)

type OrderServiceImpl struct {
	orderRepo interfaces.OrderRepository
}

func NewOrderServiceImpl(orderRepo interfaces.OrderRepository) interfaces.OrderService {
	return &OrderServiceImpl{orderRepo: orderRepo}
}

func (s *OrderServiceImpl) GetOrder(orderNumber string) (*models.Order, error) {
	// 呼叫 repository 查詢訂單（含景點與圖片）
	return s.orderRepo.GetOrderByNumber(orderNumber)
}

func (s *OrderServiceImpl) CreateOrder(order models.OrderInput, userID uint) (string, error) {
	// userID 現在由 handler 傳入，確保是登入者的 id
	// 這樣可以避免寫死帳號，讓每個登入者都能正確建立屬於自己的訂單
	// bookingID 目前暫時寫死，未來建議根據實際 booking 流程取得
	bookingID := uint(1) // TODO: 之後根據實際 booking 流程取得

	// 允許同一 user 多次下單，不檢查重複
	// 【本次更動理由】
	// 原本會檢查同一 user/booking 是否已經有訂單，避免重複下訂。
	// 但如果需求是「同一帳號可以多次下單」，就要移除這個檢查，
	// 讓每次下單都能產生新訂單編號。

	// 付款驗證（串接 TapPay）
	// 這裡會將 prime 及付款資訊送到 TapPay，確認付款成功才建立訂單
	//ok, payMsg := verifyTapPay(order.Prime, order.Order.Contact.Name, order.Order.Contact.Email, order.Order.Contact.Phone, order.Order.Price)
	ok, payMsg := verifyTapPay(order.Prime)
	if !ok {
		// 若付款失敗，直接回傳錯誤，不建立訂單
		return "", fmt.Errorf("付款失敗: %v", payMsg)
	}

	// 產生唯一訂單編號（每次下單都不同）
	// 用時間戳+亂數，確保訂單編號唯一且可追蹤
	rand.Seed(time.Now().UnixNano())
	orderNumber := fmt.Sprintf("%s%06d", time.Now().Format("20060102150405"), rand.Intn(1000000))

	// 組成訂單資料，userID/bookingID/attractionID/聯絡人等都正確寫入
	orderModel := models.Order{
		OrderNumber:  orderNumber,
		UserID:       userID,
		BookingID:    bookingID,
		AttractionID: uint(order.Order.Trip.Attraction.ID),
		Price:        order.Order.Price,
		TripDate:     order.Order.Trip.Date,
		TripTime:     order.Order.Trip.Time,
		ContactName:  order.Order.Contact.Name,
		ContactEmail: order.Order.Contact.Email,
		ContactPhone: order.Order.Contact.Phone,
		Status:       1, // 付款成功
	}
	// 寫入資料庫
	if err := s.orderRepo.CreateOrder(&orderModel); err != nil {
		return "", fmt.Errorf("資料庫寫入失敗: %v", err)
	}
	// 回傳訂單編號
	return orderNumber, nil
}

// verifyTapPay 串接 TapPay API，回傳付款是否成功與訊息
// 這個 function 會將前端送來的 prime 及付款資訊送到 TapPay，
// 並根據回傳結果決定是否建立訂單。
//
// ====== TapPay 真實串接版本（已註解） ======
/*
func verifyTapPay(prime, name, email, phone string, amount int) (bool, string) {
	// 準備 TapPay API 請求資料
	reqBody := models.TapPayRequest{
		Prime:      prime,
		PartnerKey: "partner_GT622rIWjVMyQ0k0dbONEC9jh5k4YctfXxh4mgS1onSxk6dEUsgFdEQE", // 你的 sandbox PartnerKey，請勿外洩
		MerchantID: "tppf_justinkuo719_GP_POS_INS_2",                                   // 你的 sandbox MerchantID
		Amount:     amount,
		Details:    "台北一日遊訂單",
		Cardholder: models.Cardholder{
			PhoneNumber: phone,
			Name:        name,
			Email:       email,
		},
	}
	// 將資料轉為 JSON
	jsonBody, _ := json.Marshal(reqBody)
	// 建立 HTTP POST 請求
	req, _ := http.NewRequest("POST", "https://sandbox.tappaysdk.com/tpc/payment/pay-by-prime", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	// x-api-key 必須帶 PartnerKey，這是後端的機密金鑰
	req.Header.Set("x-api-key", reqBody.PartnerKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err.Error()
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// 解析 TapPay 回傳結果
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	if status, ok := result["status"].(float64); ok && status == 0 {
		// status=0 代表付款成功
		return true, "付款成功"
	}
	// 其他狀態皆為失敗，回傳錯誤訊息
	return false, fmt.Sprintf("TapPay 錯誤: %v", result)
}
*/

// ====== Mock 版本：收到 prime 就假裝付款成功 ======
// 只保留 prime 參數，移除未使用參數避免 lint 錯誤
func verifyTapPay(prime string) (bool, string) {
	// 只要有 prime 就直接回傳付款成功（side project/mock 測試用）
	if prime != "" {
		return true, "[Mock] 付款成功 (未實際串接 TapPay)"
	}
	return false, "[Mock] 付款失敗：未收到 prime"
}

func (s *OrderServiceImpl) ProcessPayment(orderNumber string, paymentData models.PaymentInput) error {
	// TODO: implement payment logic if needed
	return nil
}

// 你需要在 handler 呼叫時，將 userID 傳進來，例如：
// orderNumber, err := orderService.CreateOrder(input, userID)
