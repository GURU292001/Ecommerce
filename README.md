# Ecommerce  sales_analytics

This project processes large-scale sales data from CSV into a normalized MariaDB database. It supports both scheduled and on-demand data refreshes with validation, deduplication, and logging. A RESTful API provides endpoints to trigger data updates and calculate key revenue metrics by date, product, category, or region. Designed for scalability and reliable business insights.



## 🛠 Tech Stack

- **Go (Golang)** `v1.22.6` – Backend APIs and CSV processing  
- **GORM** – ORM for seamless database interactions  
- **MariaDB (SQL)** – Relational database for storing normalized sales data  
- **JSON API** – For communication between frontend and backend
## 📦 External Packages

- **TOML** – For configuration file parsing  
- **GORM** – ORM for efficient database interaction  
- **Mux** – HTTP router for building RESTful APIs

## 📥 Install Dependencies

```bash
go mod tidy 
```

## 📥 Run the Server

```bash
go run main.go
```

# 🌐 WebAuthn API Documentation

This project implements a **WebAuthn-based authentication system**. Below is a list of available API endpoints.

---

## 🔗 Base URL


http://localhost:29069/
> Use this base URL to access all authentication-related endpoints during local development.

---
### 📘 API Endpoints

| Route           | Method | Request Body                                                                                      | Description                                                                                                  |
|----------------|--------|---------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------|
| `/upload-file` | GET    | _No body_                                                                                         | Used to refresh the data. It reads the latest CSV file and loads the data into the MariaDB database.        |
| `/get-revenue` | POST   | ```json\n{ "from_date": "2023-03-22", "to_date": "2024-02-21","total_revenueby":"product" }```                                | Retrieves sales revenue data: total revenue, revenue by <ins> product</ins> ,  <ins> by category</ins> , and <ins> by region </ins>  within the date range. |



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

### total_revenueby `parameters `

`total_revenueby = 'product' ` 

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
