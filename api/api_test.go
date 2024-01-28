package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test to check empty response for get receipts
func TestGetReceiptsEmpty(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/receipt", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
}

// Test to check for empty data
func TestInvalidData(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/receipt/1/points", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
}

// Test to check the response for valid data
func TestValidData(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	jsonData := `{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [
			{
			"shortDescription": "Mountain Dew 12PK",
			"price": "6.49"
			},{
			"shortDescription": "Emils Cheese Pizza",
			"price": "12.25"
			},{
			"shortDescription": "Knorr Creamy Chicken",
			"price": "1.26"
			},{
			"shortDescription": "Doritos Nacho Cheese",
			"price": "3.35"
			},{
			"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
			"price": "12.00"
			}
		],
		"total": "35.35"
	}`

	req, _ := http.NewRequest("POST", "/receipt", bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}

// Test case to check if all the receipts are returned
func TestGetReceipts(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/receipt", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

// Test case for empty retailer
func TestInvalidData1(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	jsonData := `{
		"retailer": "",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [
			{
			"shortDescription": "Mountain Dew 12PK",
			"price": "6.49"
			},{
			"shortDescription": "Emils Cheese Pizza",
			"price": "12.25"
			},{
			"shortDescription": "Knorr Creamy Chicken",
			"price": "1.26"
			},{
			"shortDescription": "Doritos Nacho Cheese",
			"price": "3.35"
			},{
			"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
			"price": "12.00"
			}
		],
		"total": "35.35"
	}`

	req, _ := http.NewRequest("POST", "/receipt", bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

// Test case for empty shortDescription
func TestInvalidData2(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	jsonData := `{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [
			{
			"shortDescription": "",
			"price": "6.49"
			},{
			"shortDescription": "Emils Cheese Pizza",
			"price": "12.25"
			},{
			"shortDescription": "Knorr Creamy Chicken",
			"price": "1.26"
			},{
			"shortDescription": "Doritos Nacho Cheese",
			"price": "3.35"
			},{
			"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
			"price": "12.00"
			}
		],
		"total": "35.35"
	}`

	req, _ := http.NewRequest("POST", "/receipt", bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

// Test case for string input for Total
func TestInvalidData3(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	jsonData := `{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [
			{
			"shortDescription": "Mountain Dew 12PK",
			"price": "6.49"
			},{
			"shortDescription": "Emils Cheese Pizza",
			"price": "12.25"
			},{
			"shortDescription": "Knorr Creamy Chicken",
			"price": "1.26"
			},{
			"shortDescription": "Doritos Nacho Cheese",
			"price": "3.35"
			},{
			"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
			"price": "12.00"
			}
		],
		"total": "ABC"
	}`

	req, _ := http.NewRequest("POST", "/receipt", bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

// Test case for string input for Price
func TestInvalidData4(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	jsonData := `{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [
			{
			"shortDescription": "Mountain Dew 12PK",
			"price": "ABC"
			},{
			"shortDescription": "Emils Cheese Pizza",
			"price": "12.25"
			},{
			"shortDescription": "Knorr Creamy Chicken",
			"price": "1.26"
			},{
			"shortDescription": "Doritos Nacho Cheese",
			"price": "3.35"
			},{
			"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
			"price": "12.00"
			}
		],
		"total": "35.35"
	}`

	req, _ := http.NewRequest("POST", "/receipt", bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

// Test case for invalid format for PurchaseDate
func TestInvalidData5(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	jsonData := `{
		"retailer": "Target",
		"purchaseDate": "2022/01/01",
		"purchaseTime": "13:01",
		"items": [
			{
			"shortDescription": "Mountain Dew 12PK",
			"price": "6.49"
			},{
			"shortDescription": "Emils Cheese Pizza",
			"price": "12.25"
			},{
			"shortDescription": "Knorr Creamy Chicken",
			"price": "1.26"
			},{
			"shortDescription": "Doritos Nacho Cheese",
			"price": "3.35"
			},{
			"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
			"price": "12.00"
			}
		],
		"total": "35.35"
	}`

	req, _ := http.NewRequest("POST", "/receipt", bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

// Test case for invalid format for PurchaseTime
func TestInvalidData6(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	jsonData := `{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13-01",
		"items": [
			{
			"shortDescription": "Mountain Dew 12PK",
			"price": "6.49"
			},{
			"shortDescription": "Emils Cheese Pizza",
			"price": "12.25"
			},{
			"shortDescription": "Knorr Creamy Chicken",
			"price": "1.26"
			},{
			"shortDescription": "Doritos Nacho Cheese",
			"price": "3.35"
			},{
			"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
			"price": "12.00"
			}
		],
		"total": "35.35"
	}`

	req, _ := http.NewRequest("POST", "/receipt", bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

// Test case to check response in case of empty items
func TestInvalidData7(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	jsonData := `{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [],
		"total": "35.35"
	}`

	req, _ := http.NewRequest("POST", "/receipt", bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

// Test case to check if id is returned in response
func TestID(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	jsonData := `{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [
			{
			"shortDescription": "Mountain Dew 12PK",
			"price": "6.49"
			},{
			"shortDescription": "Emils Cheese Pizza",
			"price": "12.25"
			},{
			"shortDescription": "Knorr Creamy Chicken",
			"price": "1.26"
			},{
			"shortDescription": "Doritos Nacho Cheese",
			"price": "3.35"
			},{
			"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
			"price": "12.00"
			}
		],
		"total": "35.35"
	}`

	req, _ := http.NewRequest("POST", "/receipt", bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
	// get the id from the response
	var response map[string]interface{}

	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal("Error decoding response body:", err)
	}

	//  Verify that the ID is returned properly
	assert.Contains(t, response, "id", "ID not found in the response")
}

// Test case to check if the points is returned in response
func TestPoints(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	jsonData := `{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [
			{
			"shortDescription": "Mountain Dew 12PK",
			"price": "6.49"
			},{
			"shortDescription": "Emils Cheese Pizza",
			"price": "12.25"
			},{
			"shortDescription": "Knorr Creamy Chicken",
			"price": "1.26"
			},{
			"shortDescription": "Doritos Nacho Cheese",
			"price": "3.35"
			},{
			"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
			"price": "12.00"
			}
		],
		"total": "35.35"
	}`

	req, _ := http.NewRequest("POST", "/receipt", bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
	// get the id from the response
	var response map[string]interface{}

	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal("Error decoding response body:", err)
	}

	// Extract the ID from the response
	id, ok := response["id"].(string)
	if !ok {
		t.Fatal("ID not found in the response")
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/receipt/"+id+"/points", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	//  Verify that the points is returned properly
	assert.Contains(t, w.Body.String(), "points", "Points not found in the response")
}

// Test case to check if the points are calculated correctly
func TestValidData2(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	jsonData := `{
		"retailer": "M&M Corner Market",
		"purchaseDate": "2022-03-20",
		"purchaseTime": "14:33",
		"items": [
		  {
			"shortDescription": "Gatorade",
			"price": "2.25"
		  },{
			"shortDescription": "Gatorade",
			"price": "2.25"
		  },{
			"shortDescription": "Gatorade",
			"price": "2.25"
		  },{
			"shortDescription": "Gatorade",
			"price": "2.25"
		  }
		],
		"total": "9.00"
	}`
	req, _ := http.NewRequest("POST", "/receipt", bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
	// get the id from the response
	var response map[string]interface{}

	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal("Error decoding response body:", err)
	}

	// Extract the ID from the response
	id, ok := response["id"].(string)
	if !ok {
		t.Fatal("ID not found in the response")
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/receipt/"+id+"/points", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\n    \"points\": 109\n}", w.Body.String())
}

// Test case to check if the receipt is returned for a valid id
func TestGetReceiptById(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	jsonData := `{
		"retailer": "Best Buy",
		"purchaseDate": "2023-11-10",
		"purchaseTime": "16:20",
		"items": [
		  {
			"shortDescription": "Sony Wireless Earbuds",
			"price": "69.99"
		  },{
			"shortDescription": "Logitech Gaming Mouse",
			"price": "49.99"
		  },{
			"shortDescription": "Samsung 4K Smart TV",
			"price": "599.99"
		  }
		],
		"total": "719.97"
	  }`
	req, _ := http.NewRequest("POST", "/receipt", bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
	// get the id from the response
	var response map[string]interface{}

	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal("Error decoding response body:", err)
	}

	// Extract the ID from the response
	fmt.Println(response, "response in test")
	id, ok := response["id"].(string)
	if !ok {
		t.Fatal("ID not found in the response")
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/receipt/"+id+"/points", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

// Test case to update the receipt
func TestUpdateReceipt(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	jsonData := `{
		"retailer": "Best Buy",
		"purchaseDate": "2023-11-10",
		"purchaseTime": "16:20",
		"items": [
		  {
			"shortDescription": "Sony Wireless Earbuds",
			"price": "69.99"
		  },{
			"shortDescription": "Logitech Gaming Mouse",
			"price": "49.99"
		  },{
			"shortDescription": "Samsung 4K Smart TV",
			"price": "599.99"
		  }
		],
		"total": "719.97"
	  }`
	req, _ := http.NewRequest("POST", "/receipt", bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
	// get the id from the response
	var response map[string]interface{}

	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal("Error decoding response body:", err)
	}

	// Extract the ID from the response
	fmt.Println(response, "response in test")
	id, ok := response["id"].(string)
	if !ok {
		t.Fatal("ID not found in the response")
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/receipt/"+id, bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

// Test case to handle updating invalid id
func TestUpdateReceiptInvalidID(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	jsonData := `{
		"retailer": "Best Buy",
		"purchaseDate": "2023-11-10",
		"purchaseTime": "16:20",
		"items": [
		  {
			"shortDescription": "Sony Wireless Earbuds",
			"price": "69.99"
		  },{
			"shortDescription": "Logitech Gaming Mouse",
			"price": "49.99"
		  },{
			"shortDescription": "Samsung 4K Smart TV",
			"price": "599.99"
		  }
		],
		"total": "719.97"
	  }`
	id := "123"
	req, _ := http.NewRequest("PUT", "/receipt/"+id, bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
}

// Test case to delete the receipt
func TestDeleteReceipt(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	jsonData := `{
		"retailer": "Best Buy",
		"purchaseDate": "2023-11-10",
		"purchaseTime": "16:20",
		"items": [
		  {
			"shortDescription": "Sony Wireless Earbuds",
			"price": "69.99"
		  },{
			"shortDescription": "Logitech Gaming Mouse",
			"price": "49.99"
		  },{
			"shortDescription": "Samsung 4K Smart TV",
			"price": "599.99"
		  }
		],
		"total": "719.97"
	  }`
	req, _ := http.NewRequest("POST", "/receipt", bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
	// get the id from the response
	var response map[string]interface{}

	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal("Error decoding response body:", err)
	}

	// Extract the ID from the response
	fmt.Println(response, "response in test")
	id, ok := response["id"].(string)
	if !ok {
		t.Fatal("ID not found in the response")
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/receipt/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

// Test case to handle deleting invalid id
func TestDeleteInvalidId(t *testing.T) {
	router := gin.Default()
	SetupRoutes(router)
	w := httptest.NewRecorder()
	id := "123"
	req, _ := http.NewRequest("DELETE", "/receipt/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
}
