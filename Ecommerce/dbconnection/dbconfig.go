package dbconnection

import (
	"database/sql"
	"ecommerce/helper"
	"ecommerce/toml"
	"log"

	"strconv"

	"gorm.io/gorm"
)

type DatabaseDetails struct {
	User     string
	Port     int
	Server   string
	Password string
	Database string
	DBType   string
	DB       string
}

type Db_Details struct {
	Mysql DatabaseDetails
}

type Db_Instance struct {
	Mysql_sqldb *sql.DB
	Gormdb      *gorm.DB
}

// Variable holds the instances of the database connection.
var Gdb_instance Db_Instance

// Struct for to hold the connection pool configuration
type connectionpoolconfig struct {
	OpenConnCount    int
	IdleConnCount    int
	MaxIdleConnCount int
}

/*
Read the database detail from the toml
*/
func DB_Details() Db_Details {
	log.Println("DB_Details (+) ")

	var lDbDetailsRec Db_Details

	lDbConfig, lErr := toml.ReadTomlFile("./toml/dbconfig.toml")
	if lErr != nil {
		log.Println("Error (DBDR01) ", lErr.Error())
	}
	lDbDetailsRec.Mysql.User = toml.GetKeyVal(lDbConfig, "Db_User")
	lDbDetailsRec.Mysql.Port, _ = strconv.Atoi(toml.GetKeyVal(lDbConfig, "Db_Port"))
	lDbDetailsRec.Mysql.Server = toml.GetKeyVal(lDbConfig, "Db_Server")
	lDbDetailsRec.Mysql.Password = toml.GetKeyVal(lDbConfig, "Db_Password")
	lDbDetailsRec.Mysql.Database = toml.GetKeyVal(lDbConfig, "Db_Database")
	lDbDetailsRec.Mysql.DBType = toml.GetKeyVal(lDbConfig, "Db_Type")
	lDbDetailsRec.Mysql.DB = toml.GetKeyVal(lDbConfig, "Db_Name")

	log.Println("DB_Details (-) ")
	return lDbDetailsRec
}

func connectionpoolConfig() connectionpoolconfig {
	log.Println("connectionpoolConfig (+) ")

	var lConpoolconfig connectionpoolconfig
	var lErr error

	// reading a connection pool details from the toml
	lDbConnectionpool, lErr := toml.ReadTomlFile("./toml/dbconfig.toml")
	if lErr != nil {
		log.Println("Error (DBDR01) ", lErr.Error())
	}
	lSetMaxOpenConns := toml.GetKeyVal(lDbConnectionpool, "SetMaxOpenConnsdb")
	lSetMaxIdleConnsdb := toml.GetKeyVal(lDbConnectionpool, "SetMaxIdleConnsdb")
	lSetConnMaxIdleTime := toml.GetKeyVal(lDbConnectionpool, "SetConnMaxIdleTimedb")

	// If the details not properly readen from the toml file this will handle the issue
	if lSetMaxOpenConns == "" {
		lSetMaxOpenConns = "10"
	}

	if lSetMaxIdleConnsdb == "" {
		lSetMaxIdleConnsdb = "5"
	}

	if lSetConnMaxIdleTime == "" {
		lSetConnMaxIdleTime = "5"
	}

	lConpoolconfig.OpenConnCount, lErr = strconv.Atoi(lSetMaxOpenConns)
	if lErr != nil {
		helper.LogError(lErr)
		return lConpoolconfig
	}

	lConpoolconfig.IdleConnCount, lErr = strconv.Atoi(lSetMaxIdleConnsdb)
	if lErr != nil {
		helper.LogError(lErr)
		return lConpoolconfig
	}

	lConpoolconfig.IdleConnCount, lErr = strconv.Atoi(lSetConnMaxIdleTime)
	if lErr != nil {
		helper.LogError(lErr)
		return lConpoolconfig
	}

	log.Println("connectionpoolConfig (-) ")
	return lConpoolconfig
}
