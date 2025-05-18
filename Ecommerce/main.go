package main

import (
	"ecommerce/Sales_Analysis/handlers"
	"ecommerce/dbconnection"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func AutoFetch() {

}

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

	fmt.Println("Server started---(+)")

	lRouter := mux.NewRouter()

	// EndPoints
	lRouter.HandleFunc("/upload-file", handlers.GetUploadFile).Methods(http.MethodGet)
	lRouter.HandleFunc("/get-revenue", handlers.GetRevenue).Methods(http.MethodPost)

	lSrv := &http.Server{
		Handler: lRouter,
		Addr:    ":29069",
	}

	log.Fatal(lSrv.ListenAndServe())

}
