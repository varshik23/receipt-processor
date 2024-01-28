package models

type Receipt struct {
	Retailer     string `json:"retailer" binding:"required"`
	PurchaseDate string `json:"purchaseDate" binding:"required"`
	PurchaseTime string `json:"purchaseTime" binding:"required"`
	Items        []Item `json:"items" binding:"required,dive"`
	Total        string `json:"total" binding:"required,numeric"`
}

type Item struct {
	ShortDescription string `json:"shortDescription" binding:"required"`
	Price            string `json:"price" binding:"required,numeric"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
