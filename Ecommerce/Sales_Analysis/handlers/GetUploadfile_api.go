package handlers

import (
	globalvar "ecommerce/GlobalVar"
	"ecommerce/Sales_Analysis/models"
	"ecommerce/helper"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
GetUploadFile is an HTTP GET API used to refresh data from a CSV file.

- Request:
  - Method: GET
  - No request body is needed.
  - This API is meant to be called from a frontend or browser.

- Headers (set in the response):
  - Access-Control-Allow-Origin: Allows requests from any origin (CORS).
  - Access-Control-Allow-Credentials: Allows cookies/credentials.
  - Access-Control-Allow-Methods: Only allows GET and OPTIONS methods.
  - Access-Control-Allow-Headers: Allows common request headers.

- Internal Logic:
  - If the method is GET:
    - Calls Set_csv_Datas() to load fresh data from a CSV into the database.
    - If successful, returns a JSON response with status and message.
    - If an error occurs during data update, returns an error with code GUF01.
    - If an error occurs while sending the response, returns error GUF02.
  - If the method is not GET, returns error GUF03.

- Response (on success):
  {
    "status": "success",
    "msg": "Successfully Data Updated"
  }

- Response (on failure):
  {
    "errorCode": "GUF01" / "GUF02" / "GUF03",
    "errorMessage": "Meaningful error message"
  }
- Author: Gurumurugan L
- Date: 2025-05-18
*/

func GetUploadFile(w http.ResponseWriter, r *http.Request) {
	log.Println("GetUploadFile(+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", " Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, credentials")
	if r.Method == http.MethodGet {
		var lResponse models.ResponseStruc
		lResponse.Status = globalvar.Success
		lResponse.Msg = "Successfully Data Updated"
		// lErr := Csv_FileReader()
		lErr := Set_csv_Datas()
		if lErr != nil {
			fmt.Fprint(w, helper.GetErrorString("GUF01", lErr.Error()))
			return
		}
		lErr = json.NewEncoder(w).Encode(&lResponse)
		if lErr != nil {
			fmt.Fprint(w, helper.GetErrorString("GUF02", lErr.Error()))
		}

	} else {
		fmt.Fprint(w, helper.GetErrorString("GUF03", "invalid Method"))
	}

	log.Println("GetUploadFile(-)")

}
