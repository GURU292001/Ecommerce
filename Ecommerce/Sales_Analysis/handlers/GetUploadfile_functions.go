package handlers

import (
	"ecommerce/Sales_Analysis/models"
	"ecommerce/dbconnection"
	"ecommerce/helper"
	"ecommerce/toml"
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"

	"gorm.io/gorm/clause"
)

// ---------------------------------------------GetUploadfile.go functions (+)------------------------------------------------

// Even the million of data in csv construct and bulk insert will happen
func Set_csv_Datas() error {
	lCustomers, lRegions, lProducts, lOrders, lErr := Csv_FileReader()
	if lErr != nil {
		helper.LogError(lErr)
	}
	// Region
	lRegions, lErr = Get_Region_Data(lRegions)
	if lErr != nil {
		helper.LogError(lErr)
	}
	// convert the arraystruc to map
	RegionMap := make(map[string]int)
	for _, value := range lRegions {
		RegionMap[value.Name] = value.Id
	}

	for idx, value := range lOrders {
		if match, exists := RegionMap[value.Region_name]; exists {
			// Update fields as needed
			log.Println("lOrders[idx].Region:", lOrders[idx].Region_name)
			log.Println("strconv.Itoa(match)", strconv.Itoa(match))
			lOrders[idx].Region_id = match
		}
	}
	//customer
	lErr = Upload_Customer_data(lCustomers)
	if lErr != nil {
		helper.LogError(lErr)
		return lErr
	}
	// Products
	lErr = Upload_Product_data(lProducts)
	if lErr != nil {
		helper.LogError(lErr)
		return lErr
	}
	// Orders
	lErr = Upload_Orders_data(lOrders)
	if lErr != nil {
		helper.LogError(lErr)
		return lErr
	}

	return nil
}

// Read the Csv completely, and insert into the table at minimum db_Connection
func Csv_FileReader() (lCustomer_List []models.CustomerStruc, lRegion_List []models.Region, lProduct_List []models.ProductStruc, lOrder_List []models.OrderStruc, lErr error) {
	log.Println("Csv_FileReader(+)")

	ltomlItems, lErr := toml.ReadTomlFile("./toml/filereadconfig.toml")
	if lErr != nil {
		helper.LogError(lErr)
		return lCustomer_List, lRegion_List, lProduct_List, lOrder_List, lErr
	}
	lFilePath := toml.GetKeyVal(ltomlItems, "FileReadyPath")
	// Open the CSV file
	file, lErr := os.Open(lFilePath)
	if lErr != nil {
		helper.LogError(lErr)
		return lCustomer_List, lRegion_List, lProduct_List, lOrder_List, lErr
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records
	records, lErr := reader.ReadAll()
	if lErr != nil {
		log.Println("Failed to read CSV: ", lErr)
		helper.LogError(lErr)
		return lCustomer_List, lRegion_List, lProduct_List, lOrder_List, lErr
	}

	for _, col := range records[1:] {
		var lCustomer models.CustomerStruc
		var lProduct models.ProductStruc
		var lRegion models.Region
		var lOrder models.OrderStruc
		// Read each record and assign to struct
		lOrder.Order_Id = col[0]
		lProduct.ProductId = col[1]
		lCustomer.Customer_Id = col[2]
		lProduct.ProductName = col[3]
		lProduct.Category = col[4]
		lRegion.Name = col[5]
		lOrder.Region_name = col[5]
		lOrder.Date_of_sale = col[6]
		lOrder.Quantity_sold, _ = strconv.ParseFloat(col[7], 64)
		lOrder.Unit_price, _ = strconv.ParseFloat(col[8], 64)
		lProduct.Unit_price = lOrder.Unit_price
		log.Println("lErr:", lErr)
		log.Println("col[8] unit Price:", col[8], lOrder.Unit_price)
		lOrder.Discount, _ = strconv.ParseFloat(col[9], 64)
		lOrder.Shipping_cost, _ = strconv.ParseFloat(col[10], 64)
		log.Println("col[10] shipprig cost:", col[10], lOrder.Shipping_cost, ",lErr:", lErr)
		lOrder.Payment_method = col[11]
		lCustomer.Customer_name = col[12]
		lCustomer.Customer_email = col[13]
		lCustomer.Customer_address = col[14]

		lOrder.Product_Id = col[1]
		lOrder.Customer_Id = col[2]

		log.Println("Customer: ", lCustomer)
		log.Println("Product: ", lProduct)
		log.Println("Order: ", lOrder)
		log.Println("region: ", lRegion)

		lCustomer_List = append(lCustomer_List, lCustomer)
		lProduct_List = append(lProduct_List, lProduct)
		lOrder_List = append(lOrder_List, lOrder)
		lRegion_List = append(lRegion_List, lRegion)

	}

	log.Println("Csv_FileReader(-)")
	return lCustomer_List, lRegion_List, lProduct_List, lOrder_List, nil
}

func Upload_Customer_data(pCustomer []models.CustomerStruc) error {
	log.Println("Upload_Customer_data(+)")

	lErr := dbconnection.Gdb_instance.Gormdb.
		Table("customers").
		Clauses(clause.Insert{Modifier: "IGNORE"}).CreateInBatches(pCustomer, 1000).Error
	if lErr != nil {
		helper.LogError(lErr)
		return lErr
	}
	log.Println("Upload_Customer_data(-)")
	return nil
}

func Upload_Product_data(pProduct []models.ProductStruc) error {
	log.Println("Upload_Product_data(+)")

	lErr := dbconnection.Gdb_instance.Gormdb.
		Table("products").
		Clauses(clause.Insert{Modifier: "IGNORE"}).CreateInBatches(pProduct, 1000).Error
	if lErr != nil {
		helper.LogError(lErr)
		return lErr
	}
	log.Println("Upload_Product_data(-)")
	return nil
}

func Upload_Orders_data(pOrder []models.OrderStruc) error {
	log.Println("Upload_Orders_data(+)")

	lErr := dbconnection.Gdb_instance.Gormdb.
		Table("Orders").
		Clauses(clause.Insert{Modifier: "IGNORE"}).CreateInBatches(pOrder, 1000).Error
	if lErr != nil {
		helper.LogError(lErr)
		return lErr
	}
	log.Println("Upload_Orders_data(-)")
	return nil
}

// filter, insert ,fetch the Region Data
func Get_Region_Data(pRegions []models.Region) (Regions []models.Region, err error) {

	// Remove Duplicates
	lRegions, lErr := Remove_Duplicate_Region(pRegions)
	if lErr != nil {
		helper.LogError(lErr)
		return Regions, lErr
	}

	// insert into the table
	lErr = Upload_Region_data(lRegions)
	if lErr != nil {
		helper.LogError(lErr)
		return Regions, lErr
	}

	Regions, lErr = GetRegionData()
	if lErr != nil {
		helper.LogError(lErr)
		return Regions, lErr
	}

	log.Println("Regions--------:", Regions)

	return Regions, nil
}

// Insert the region data
func Upload_Region_data(pRegion []models.Region) error {
	log.Println("Upload_Region_data(+)")

	lErr := dbconnection.Gdb_instance.Gormdb.
		Table("region").
		Clauses(clause.Insert{Modifier: "IGNORE"}).
		CreateInBatches(pRegion, 1000).Error

	if lErr != nil {
		helper.LogError(lErr)
		return lErr
	}
	log.Println("Upload_Region_data(-)")
	return nil
}

// Get the region data
func GetRegionData() ([]models.Region, error) {
	log.Println("GetRegionData(+)")
	lRegions := []models.Region{}
	lErr := dbconnection.Gdb_instance.Gormdb.Table("region").Find(&lRegions).Error
	if lErr != nil {
		helper.LogError(lErr)
		return lRegions, lErr
	}

	log.Println("GetRegionData(-)")
	return lRegions, nil
}

// Remove the duplicate region list
func Remove_Duplicate_Region(input []models.Region) ([]models.Region, error) {
	log.Println("en(input):", input)
	if len(input) < 2 {
		return input, errors.New("empty or single data is present")
	}

	seen := make(map[string]bool) // map[ID]bool
	unique := []models.Region{}

	for _, r := range input {
		if !seen[r.Name] {
			seen[r.Name] = true
			unique = append(unique, r)
		}
	}

	return unique, nil

}

// ---------------------------------------------GetUploadfile.go functions (-)------------------------------------------------
