package main

import (
	"ecommerce/Sales_Analysis/handlers"
	"ecommerce/dbconnection"
	"ecommerce/helper"
	"ecommerce/toml"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	lFile, lErr := os.OpenFile("./log/logfile"+time.Now().Format("02012006.15.04.05.000000000")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if lErr != nil {
		log.Fatalf("error opening file: %v", lErr)
	}
	defer lFile.Close()
	log.SetOutput(lFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Database connection process
	lErr = dbconnection.BuildConnection()
	if lErr != nil {
		log.Fatal(lErr)
	}
	defer dbconnection.Gdb_instance.Mysql_sqldb.Close()

	fmt.Println("Server started(+)")

	lRouter := mux.NewRouter()
	// -------Data Refresh Mechanism (+)--------

	lDbConfig, lErr := toml.ReadTomlFile("./toml/serviceconfig.toml")
	if lErr != nil {
		log.Println("Error (DBDR01) ", lErr.Error())
	}
	lAutoRun := toml.GetKeyVal(lDbConfig, "AutoRun")
	log.Println("lAutoRun", lAutoRun)
	if lAutoRun == "Y" {
		go scheduleAt()
	}
	// -------Data Refresh Mechanism (-)--------

	// EndPoints
	lRouter.HandleFunc("/upload-file", handlers.GetUploadFile).Methods(http.MethodGet)
	lRouter.HandleFunc("/get-revenue", handlers.GetRevenue).Methods(http.MethodPost)

	lSrv := &http.Server{
		Handler: lRouter,
		Addr:    ":29069",
	}

	log.Fatal(lSrv.ListenAndServe())

}

/*
scheduleAt schedules and runs a daily data refresh task at a configured time.

- Purpose: Triggers a database refresh at a specific time set in the TOML config.
- Config File: ./toml/serviceconfig.toml
- Config Keys: "hour", "minute"
- Behavior:
  - Polls the current time every 30 seconds.
  - When the time matches the configured hour and minute, it calls Set_csv_Datas().
  - Sleeps for 61 seconds after execution to prevent duplicate runs within the same minute.
- Inputs: None
- Outputs: None (performs database operations as a side effect)
- Author: Gurumurugan L
- Date: 2025-05-18
*/

func scheduleAt() {

	for {

		lNow := time.Now()

		lDbConfig, lErr := toml.ReadTomlFile("./toml/serviceconfig.toml")
		if lErr != nil {
			log.Println("Error (DBDR01) ", lErr.Error())
		}
		lhourStr := toml.GetKeyVal(lDbConfig, "hour")
		lminutesStr := toml.GetKeyVal(lDbConfig, "minute")

		lhour, err := strconv.Atoi(lhourStr)
		if err != nil {
			log.Printf("Invalid hour value: %v", err)
		}
		lminutes, err := strconv.Atoi(lminutesStr)
		if err != nil {
			log.Printf("Invalid minute value: %v", err)
		}
		log.Println("lhour", lhour, ",minute:", lminutes)
		if lNow.Hour() == lhour && lNow.Minute() == lminutes {
			lErr = handlers.Set_csv_Datas()
			if lErr != nil {
				helper.LogError(lErr)
			}
			time.Sleep(61 * time.Second)

		} else {
			time.Sleep(30 * time.Second)
		}

	}

}

/* func Data_Refresh_Mechanism() {
	for {
		lDbConfig, lErr := toml.ReadTomlFile("./toml/serviceconfig.toml")
		if lErr != nil {
			log.Println("Error (DBDR01) ", lErr.Error())
		}
		lhourStr := toml.GetKeyVal(lDbConfig, "hour")
		lminutesStr := toml.GetKeyVal(lDbConfig, "minute")

		lhour, err := strconv.Atoi(lhourStr)
		if err != nil {
			log.Printf("Invalid hour value: %v", err)
			continue
		}
		lminutes, err := strconv.Atoi(lminutesStr)
		if err != nil {
			log.Printf("Invalid minute value: %v", err)
			continue
		}
		log.Println("hour:", lhour, ",minute:", lminutes)
		now := time.Now()

		if now.Hour() == lhour && now.Minute() == lminutes {

			lErr = handlers.Set_csv_Datas()
			if lErr != nil {
				helper.LogError(lErr)
			}

			time.Sleep(61 * time.Second)

		} else {
			time.Sleep(30 * time.Second)
		}

	}
} */
