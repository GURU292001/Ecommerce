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
GetRevenue handles POST requests to calculate and return revenue details.

- Request:
  - Method: POST
  - Body: JSON matching the Revenue_Req model.

- Internal Logic:
  - Decode JSON request into Revenue_Req struct.
  - Call Revenue_Calculation with the input data.
  - If successful, set status and send Revenue_Resp as JSON.
  - On error, return appropriate error code and message.

- Response (success):
  JSON with status and revenue data.

- Response (failure):
  JSON with errorCode and errorMessage.

Author: Gurumurugan L
Date: 2025-05-18
*/

func GetRevenue(w http.ResponseWriter, r *http.Request) {
	log.Println("GetUploadFile(+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", " Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, credentials")
	if r.Method == http.MethodPost {
		var lRequestDetails models.Revenue_Req
		var lResponse models.Revenue_Resp
		lErr := json.NewDecoder(r.Body).Decode(&lRequestDetails)
		if lErr != nil {
			fmt.Fprint(w, helper.GetErrorString("GR01", lErr.Error()))
			return
		}
		log.Println("--------------REvencue called-------------")
		lResponse, lErr = Revenue_Calculation(lRequestDetails)
		if lErr != nil {
			fmt.Fprint(w, helper.GetErrorString("GR02", lErr.Error()))
			return
		}
		lResponse.Status = globalvar.Success
		lErr = json.NewEncoder(w).Encode(lResponse)
		if lErr != nil {
			fmt.Fprint(w, helper.GetErrorString("GR01", lErr.Error()))
			return
		}

	}
}
