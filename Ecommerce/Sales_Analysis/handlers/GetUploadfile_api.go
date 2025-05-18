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
