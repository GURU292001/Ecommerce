package models

import "time"

// Customer Details
type CustomerStruc struct {
	Customer_Id      string    `gorm:"column:customer_id"`
	Customer_name    string    `gorm:"column:customer_name"`
	Customer_email   string    `gorm:"column:customer_email"`
	Customer_address string    `gorm:"column:customer_address"`
	CreatedDate      time.Time `gorm:"column:createdDate;autoCreateTime"`
	UpdatedDate      time.Time `gorm:"column:UpdatedDate;autoUpdateTime"`
}

// Product Details
type ProductStruc struct {
	ProductId   string    `gorm:"column:product_id"`
	ProductName string    `gorm:"column:product_name"`
	Unit_price  float64   `gorm:"column:unit_price"`
	Category    string    `gorm:"column:product_category"`
	CreatedDate time.Time `gorm:"column:createdDate;autoCreateTime"`
	UpdatedDate time.Time `gorm:"column:UpdatedDate;autoUpdateTime"`
}

// Order Details
type OrderStruc struct {
	Order_Id       string  `gorm:"column:order_id"`
	Product_Id     string  `gorm:"column:product_id"`
	Customer_Id    string  `gorm:"column:customer_id"`
	Date_of_sale   string  `gorm:"column:date_of_sale"`
	Quantity_sold  float64 `gorm:"column:quantity_sold"`
	Unit_price     float64 `gorm:"-"`
	Discount       float64 `gorm:"column:discount"`
	Shipping_cost  float64 `gorm:"column:shipping_cost"`
	Payment_method string  `gorm:"column:payment_method"`
	Region_name    string  `gorm:"-"`
	Region_id      int     `gorm:"column:region_id"`
	// Category       string `gorm:"column:category"`
	CreatedDate time.Time `gorm:"column:createdDate;autoCreateTime"`
	UpdatedDate time.Time `gorm:"column:UpdatedDate;autoUpdateTime"`
}

type Region struct {
	Id          int       `gorm:"column:id"`
	Name        string    `gorm:"column:region_name"`
	CreatedDate time.Time `gorm:"column:createdDate;autoCreateTime"`
	UpdatedDate time.Time `gorm:"column:UpdatedDate;autoUpdateTime"`
}

type ResponseStruc struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

type Revenue_Req struct {
	From_date       string `json:"from_date"`
	To_date         string `json:"to_date"`
	Total_RevenueBy string `json:"total_revenueby"`
}

type Revenue_Resp struct {
	Status              string             `json:"status"`
	Msg                 string             `json:"msg"`
	Total_revenue       float64            `json:"total_revenue"`
	Revenue_by_product  []Product_revenue  `json:"revenue_by_product"`
	Revenue_by_category []Category_revenue `json:"revenue_by_category"`
	Revenue_by_region   []Region_revenue   `json:"revenue_by_region"`
}

type Product_revenue struct {
	Product_id    string  `gorm:"column:product_id"`
	Product_name  string  `gorm:"column:product_name"`
	Total_revenue float64 `gorm:"column:total_revenue"`
}

type Region_revenue struct {
	Region_name   string  `gorm:"column:region_name"`
	Total_revenue float64 `gorm:"column:total_revenue"`
}

type Category_revenue struct {
	Category      string  `gorm:"column:category"`
	Total_revenue float64 `gorm:"column:total_revenue"`
}
