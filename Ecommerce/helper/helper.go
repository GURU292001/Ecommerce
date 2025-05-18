package helper

import (
	globalvar "ecommerce/GlobalVar"
	"encoding/json"
	"log"

	"runtime"
)

type Error_Response struct {
	Status       string `json:"status"`
	ErrorCode    string `json:"statusCode"`
	ErrorMessage string `json:"msg"`
}

func GetErrorString(pErrTitle string, pDescription string) string {

	var lErr_Response Error_Response
	lErr_Response.Status = globalvar.Success
	lErr_Response.ErrorCode = pErrTitle
	lErr_Response.ErrorMessage = pDescription

	lData, lErr := json.Marshal(lErr_Response)
	if lErr != nil {
		log.Fatal(lErr.Error())
	}

	return string(lData)
}

func LogError(pErr error) {
	if pErr == nil {
		return
	}

	lPc, _, lLine, lOk := runtime.Caller(1)
	lDetails := runtime.FuncForPC(lPc)
	lFuncName := "unknown"
	if lOk && lDetails != nil {
		lFuncName = lDetails.Name()
	}

	log.Printf("[ERROR] Func: %s | Line: %d | Error: %v\n", lFuncName, lLine, pErr)
}
