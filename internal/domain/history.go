package domain

import "time"

const (
	// StatusSuccess is a string that represents a success status
	StatusSuccess = "success"
	// StatusProcessing is a string that represents a processing status
	StatusProcessing = "processing"
	// StatusError is a string that represents an error status
	StatusError = "error"
)

type History struct {
	Shippings []ShippingHistory `json:"shippings"`
}

type ShippingHistory struct {
	ID         string     `gorm:"primaryKey" json:"id"`
	FileName   string     `json:"fileName"`
	Status     string     `gorm:"index" json:"status"`
	Returnpath string     `json:"returnpath"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}
