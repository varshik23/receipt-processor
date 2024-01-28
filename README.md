# receipt-processor

"Receipt Processor" is a web service implemented in the Go programming language, designed to process receipts and compute the corresponding points based on specific rules. The application leverages the Gin framework for web development and is encapsulated within a Docker container for streamlined deployment and portability.

## Table of Contents
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
- [Usage](#usage)
    - [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [Rules](#rules)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Prerequisites
- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [Git](https://git-scm.com/downloads)

### Installation
1. Clone the repository
```sh
git clone https://github.com/varshik23/receipt-processor.git
```
2. Change directory to the repository
```sh
cd receipt-processor
```
3. Build the Docker image
```sh
docker build --tag receipt-processor:latest .
```
4. Running the Docker container
- Run the Docker container in the background
```sh
docker run -d -p 8080:8080 receipt-processor:latest
```
- Run the Docker container in the foreground
```sh
docker run -t -p 8080:8080 receipt-processor:latest
```
5. Access the application at http://localhost:8080

## Usage
### API Endpoints
- **GET** `/receipts`
    - Returns a list of all receipts
    - Response
        - Status : 200 OK
        ```json
        [
            {
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
            },
            {
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
            }
        ]
        ```
- **GET** `/receipts/:id`
    - Returns a receipt with the given ID
    - Response
        - Status: 200 OK
        ```json 
        { "id" : "d7da31e6-a9b0-42b1-af56-4eab5079e96d" }
        ```
- **POST** `/receipts`
    - Creates a new receipt
    - Request
    ```json
    {
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
    }
    ```
    - Response
        - Status: 201 Created
        ```json
        { "points" : 28 }
        ```
- **PUT** `/receipts/:id`
    - Updates a receipt with the given ID
    - Request
    ```json
    {
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
    }
    ```
    - Response
        - Status: 200 OK
        ```json
        { 
            "updated points": 109,
            "message": "Receipt with id 5190881e-1552-441c-972c-9f0be1f5778f updated"
        }
        ```
- **DELETE** `/receipts/:id`
    - Deletes a receipt with the given ID
    - Response
        - Status: 204
        ```json
        { "message": "Receipt with id e53dffab-b729-49b1-81ca-4f61ce0f1714 deleted" }
        ```

## Testing
- Run the tests
```sh
go test ./...
```

## Rules
These rules collectively define how many points should be awarded to a receipt.

1. One point for every alphanumeric character in the retailer name.
2. 50 points if the total is a round dollar amount with no cents.
3. 25 points if the total is a multiple of 0.25.
4. 5 points for every two items on the receipt.
5. If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
6. 6 points if the day in the purchase date is odd.
7. 10 points if the time of purchase is after 2:00pm and before 4:00pm.

## Examples
- **Example 1**
```json
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
```
```
Total Points: 28
Breakdown:
     6 points - retailer name has 6 characters
    10 points - 4 items (2 pairs @ 5 points each)
     3 Points - "Emils Cheese Pizza" is 18 characters (a multiple of 3)
                item price of 12.25 * 0.2 = 2.45, rounded up is 3 points
     3 Points - "Klarbrunn 12-PK 12 FL OZ" is 24 characters (a multiple of 3)
                item price of 12.00 * 0.2 = 2.4, rounded up is 3 points
     6 points - purchase day is odd
  + ---------
  = 28 points
```
- **Example 2**
```json
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
```
```
Total Points: 109
Breakdown:
    50 points - total is a round dollar amount
    25 points - total is a multiple of 0.25
    14 points - retailer name (M&M Corner Market) has 14 alphanumeric characters
                note: '&' is not alphanumeric
    10 points - 2:33pm is between 2:00pm and 4:00pm
    10 points - 4 items (2 pairs @ 5 points each)
  + ---------
  = 109 points
  ```

## Author
Varshik Sonti - [Linkedin](https://www.linkedin.com/in/varshik-sonti/) - [GitHub](https://github.com/varshik23) - [Portofolio](https://vsonti.tech)
