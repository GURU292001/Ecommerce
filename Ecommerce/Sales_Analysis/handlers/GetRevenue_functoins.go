package handlers

import (
	"ecommerce/Sales_Analysis/models"
	"ecommerce/dbconnection"
	"ecommerce/helper"
	"log"
)

// Business Logic
func Revenue_Calculation(pData models.Revenue_Req) (Revenue models.Revenue_Resp, Err error) {
	log.Println("Revenue_Calculation(+)")
	switch pData.Total_RevenueBy {
	case "product":
		Revenue, Err = Revenue_by_product(pData)
		if Err != nil {
			helper.LogError(Err)
		}
		log.Println("Total_RevenueBy by ", pData.Total_RevenueBy)
	case "category":
		Revenue, Err = Revenue_by_category(pData)
		if Err != nil {
			helper.LogError(Err)
		}
		log.Println("Total_RevenueBy by ", pData.Total_RevenueBy)
	case "region":

		Revenue, Err = Revenue_by_region(pData)
		if Err != nil {
			helper.LogError(Err)
		}
		log.Println("Total_RevenueBy by ", pData.Total_RevenueBy)

		log.Println("Total_RevenueBy by ", pData.Total_RevenueBy)
	case "overall":
		Revenue, Err = Total_revenue(pData)
		if Err != nil {
			helper.LogError(Err)
		}
		log.Println("Total_RevenueBy by overAll", pData.Total_RevenueBy)
	default:
		Revenue, Err = Total_revenue(pData)
		if Err != nil {
			helper.LogError(Err)
		}
		log.Println("Total_RevenueBy by overAll", pData.Total_RevenueBy)
	}

	log.Println("Revenue_Calculation(-)")
	return Revenue, nil
}

// Revenue by product
func Revenue_by_product(pInput models.Revenue_Req) (Revenue models.Revenue_Resp, Err error) {
	log.Println("Revenue_by_product(+)")

	var lProductRevenue []models.Product_revenue

	lErr := dbconnection.Gdb_instance.Gormdb.Table("Orders o").Select("p.product_id ,nvl(p.product_name,'') product_name,(sum(o.quantity_sold)*p.unit_price)-sum(discount )+sum(o.shipping_cost ) total_revenue").Joins("JOIN products p ON o.product_id = p.product_id").Where("date_of_sale between ? and ?", pInput.From_date, pInput.To_date).Group("p.product_id ,p.unit_price").Scan(&lProductRevenue).Error
	if lErr != nil {
		helper.LogError(lErr)
		return Revenue, lErr
	}
	Revenue.Revenue_by_product = lProductRevenue
	log.Println("Revenue_by_product(-)")
	return Revenue, nil
}

// Revenue_by_category
func Revenue_by_category(pInput models.Revenue_Req) (Revenue models.Revenue_Resp, Err error) {
	log.Println("Revenue_by_category(+)")

	var lCategoryRevenue []models.Category_revenue

	subQuery := dbconnection.Gdb_instance.Gormdb.
		Table("Orders o").
		Select(`COALESCE(p.product_category, '') AS category,
                SUM(o.quantity_sold) * p.unit_price - SUM(o.discount) + SUM(o.shipping_cost) AS total_revenue`).
		Joins("JOIN products p ON o.product_id = p.product_id").
		Where("o.date_of_sale BETWEEN ? AND ?", pInput.From_date, pInput.To_date).
		Group("p.product_category, p.product_id")

	lErr := dbconnection.Gdb_instance.Gormdb.Table("(?) as c", subQuery).
		Select("c.category, SUM(c.total_revenue) as total_revenue").
		Group("c.category").Scan(&lCategoryRevenue).Error

	if lErr != nil {
		helper.LogError(lErr)
		return Revenue, lErr
	}
	Revenue.Revenue_by_category = lCategoryRevenue
	log.Println("Revenue_by_category(-)")
	return Revenue, nil
}

// REvenue by region
func Revenue_by_region(pInput models.Revenue_Req) (Revenue models.Revenue_Resp, Err error) {
	log.Println("Revenue_by_region(+)")

	var lRegionRevenue []models.Region_revenue

	lErr := dbconnection.Gdb_instance.Gormdb.Table("Orders o").
		Select(`r.region_name,
            SUM((o.quantity_sold * p.unit_price) - o.discount + o.shipping_cost) AS total_revenue`).
		Joins("JOIN region r ON o.region_id = r.id").
		Joins("JOIN products p ON o.product_id = p.product_id").
		Group("o.region_id, r.region_name").Scan(&lRegionRevenue).Error
	if lErr != nil {
		helper.LogError(lErr)
		return Revenue, lErr
	}
	Revenue.Revenue_by_region = lRegionRevenue
	log.Println("Revenue_by_region(-)")
	return Revenue, nil
}

// total Revenue
func Total_revenue(pInput models.Revenue_Req) (Revenue models.Revenue_Resp, Err error) {
	log.Println("Total_revenue(+)")

	var lTotalRevenue float64

	lErr := dbconnection.Gdb_instance.Gormdb.Table("Orders o").Select("sum((o.quantity_sold * p.unit_price)- o.discount + o.shipping_cost ) total_revenue").Joins("JOIN products p ON o.product_id = p.product_id").Where("date_of_sale between ? and ?", pInput.From_date, pInput.To_date).Scan(&lTotalRevenue).Error
	if lErr != nil {
		helper.LogError(lErr)
		return Revenue, lErr
	}
	Revenue.Total_revenue = lTotalRevenue
	log.Println("Total_revenue(-)")
	return Revenue, nil
}
