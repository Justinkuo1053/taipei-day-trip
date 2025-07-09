package models

import "time"

type Order struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	OrderNumber   string    `json:"order_number" gorm:"size:50;unique;not null"`
	UserID        uint      `json:"user_id" gorm:"not null"`
	BookingID     uint      `json:"booking_id" gorm:"not null"`
	Status        string    `json:"status" gorm:"size:20;not null"` // pending, paid, cancelled
	TotalAmount   int       `json:"total_amount" gorm:"not null"`
	ContactName   string    `json:"contact_name" gorm:"size:255"`
	ContactEmail  string    `json:"contact_email" gorm:"size:255"`
	ContactPhone  string    `json:"contact_phone" gorm:"size:20"`
	PaymentStatus string    `json:"payment_status" gorm:"size:20"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	User    User    `json:"user" gorm:"foreignKey:UserID"`
	Booking Booking `json:"booking" gorm:"foreignKey:BookingID"`
}

type AttractionInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Image   string `json:"image"`
}

type TripInfo struct {
	Attraction AttractionInfo `json:"attraction"`
	Date       string         `json:"date"`
	Time       string         `json:"time"`
}

type ContactInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type OrderInput struct {
	Prime string `json:"prime"`
	Order struct {
		Price   int         `json:"price"`
		Trip    TripInfo    `json:"trip"`
		Contact ContactInfo `json:"contact"`
	} `json:"order"`
}

type PaymentResult struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type OrderCreateResponse struct {
	Data *struct {
		Number  string        `json:"number"`
		Payment PaymentResult `json:"payment"`
	} `json:"data,omitempty"`
	Error   bool   `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

type OrderDetailResponse struct {
	Data *struct {
		Number  string      `json:"number"`
		Price   int         `json:"price"`
		Trip    TripInfo    `json:"trip"`
		Contact ContactInfo `json:"contact"`
		Status  int         `json:"status"`
	} `json:"data,omitempty"`
	Error   bool   `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

type PaymentInput struct {
	Prime  string `json:"prime"`
	Amount int    `json:"amount"`
}
