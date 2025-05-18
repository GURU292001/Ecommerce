package dbconnection

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Dbconnection() (*gorm.DB, *sql.DB, error) {
	log.Println("Dbconnection (+) ")

	// var lConstring string
	var lDb *sql.DB
	var lErr error
	var lGormDb *gorm.DB

	lDbDetailsRec := DB_Details()
	lConpoolconfig := connectionpoolConfig()
	// `root:root@123@tcp(localhost:3306)/?charset=utf8mb4&parseTime=True&loc=Local`
	lcoreStr := fmt.Sprintf(`%s:%s@tcp(%s:%d)/dev?charset=utf8mb4&parseTime=True&loc=Local`, lDbDetailsRec.Mysql.User, lDbDetailsRec.Mysql.Password, lDbDetailsRec.Mysql.Server, lDbDetailsRec.Mysql.Port)

	lGormDb, lErr = gorm.Open(mysql.Open(lcoreStr), &gorm.Config{})
	if lErr != nil {
		// helperpkg.LogError(lErr)
		return lGormDb, lDb, lErr
	}

	lDb, lErr = lGormDb.DB()
	if lErr != nil {
		// helperpkg.LogError(lErr)
		return lGormDb, lDb, lErr
	}

	lDb.SetMaxIdleConns(lConpoolconfig.OpenConnCount)

	lDb.SetMaxOpenConns(lConpoolconfig.IdleConnCount)

	lDb.SetConnMaxIdleTime(time.Second * time.Duration(lConpoolconfig.MaxIdleConnCount))

	log.Println("Dbconnection (-) ")
	return lGormDb, lDb, lErr
}
