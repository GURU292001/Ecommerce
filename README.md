# Ecommerce  sales_analytics

This project processes a large CSV file containing historical sales data, potentially consisting of millions of rows.

# Technology Stack

#  go version 1.22.6
- Go (Golang): Backend APIs and CSV processing

- GORM: ORM for database interactions
- MariaDB(SQL): Database to store sales records
- JSON API: For frontend/backend communication

## External Package

- Toml 
- GORM 
- MUX

## Install Dependencies

go mod tidy

## Run

go run main.go

## LOG file 
The application creates logs while running to help track the execution flow and debug issues.
Log Location: log/logfile18052025.14.09.06.977875964.txt

# WebAuthn API Documentation

This project implements a WebAuthn-based authentication system. Below is a list of available API endpoints.

---
# Base URl

http://localhost:29069/
---
### 📘 API Endpoints

| Route                   | Method | Request Body                                                                                  |
|------------------------|--------|-----------------------------------------------------------------------------------------------|
| `/upload-file` | GET | ``` no Body```          |
| `/get-revenue`          | POST   | ```json\n{ "username": "johndoe", "displayName": "John Doe" }```                             |                                                                       |



### ✅ Success Response Format for `/upload-file`

```json
{"status":"S","msg":"Successfully Data Updated"}
```
### ❌ Failure Response Format (Validation or User Error)

```json
{
    "status": "S",
    "statusCode": "GUF01",
    "msg": "Error 1054 (42S22): Unknown column 'unit_price' in 'INSERT INTO'"
}
```







### ✅ Success Response Format `/get-revenue` 

`Request Body`


```json
{
    "from_date": "2023-03-10",
    "to_date": "2024-04-22",
    "total_revenueby": "category"
}
```

`total_revenueby parameters `

` total_revenueby = 'product'` 
---
```json
{
    "status": "S",
    "msg": "",
    "total_revenue": 0,
    "revenue_by_product": [
        {
            "Product_id": "P123",
            "Product_name": "UltraBoost Running Shoes",
            "Total_revenue": 1115.8
        },
        {
            "Product_id": "P234",
            "Product_name": "Sony WH-1000XM5 Headphones",
            "Total_revenue": 723.68
        },
        {
            "Product_id": "P456",
            "Product_name": "iPhone 15 Pro",
            "Total_revenue": 2628
        },
        {
            "Product_id": "P789",
            "Product_name": "Levi's 501 Jeans",
            "Total_revenue": 369.54
        }
    ],
    "revenue_by_category": null,
    "revenue_by_region": null
}
```


` Revenue By product = 'overall' or ''` 
---
```json
{
    "status": "S",
    "msg": "",
    "total_revenue": 4837.02,
    "revenue_by_product": null,
    "revenue_by_category": null,
    "revenue_by_region": null
}
```

` Revenue By Category = 'category'` 
---
```json
{
    "status": "S",
    "msg": "",
    "total_revenue": 0,
    "revenue_by_product": null,
    "revenue_by_category": [
        {
            "Category": "Clothing",
            "Total_revenue": 369.54
        },
        {
            "Category": "Electronics",
            "Total_revenue": 3351.68
        },
        {
            "Category": "Shoes",
            "Total_revenue": 1115.8
        }
    ],
    "revenue_by_region": null
}
```


` Revenue By region = 'region'` 

---
```json
{
    "status": "S",
    "msg": "",
    "total_revenue": 0,
    "revenue_by_product": null,
    "revenue_by_category": null,
    "revenue_by_region": [
        {
            "Region_name": "North America",
            "Total_revenue": 1463.48
        },
        {
            "Region_name": "Europe",
            "Total_revenue": 2628
        },
        {
            "Region_name": "Asia",
            "Total_revenue": 5605.44
        },
        {
            "Region_name": "South America",
            "Total_revenue": 376
        }
    ]
}
```



### ❌ Failure Response Format (Validation or User Error)

```json
{
    "status": "E",
    "msg": "",
    "total_revenue": 0,
    "revenue_by_product": null,
    "revenue_by_category": null,
    "revenue_by_region": null
}
```

###  🔁 Refresh Mechanism Implemented
- A daily routine runs at a configurable time specified in the TOML file.
- At the scheduled time, it clears the database and reloads fresh data.
- The refresh mechanism can be enabled or disabled via the TOML configuration.
- The execution time and related settings are fully customizable in the TOML file.


### 📝 Notes

- All endpoints return JSON.
- please change the toml file according to you requirement
