package utilities

import (
	"math/rand"
	"net/http"
	"pa-back/jsonmodel"
	"pa-back/resources/consts"

	"encoding/json"

	"github.com/speps/go-hashids"
)

// FillHTTPResponse func
func FillHTTPResponse(responseWriter http.ResponseWriter, status int, hasError bool, message string, data interface{}) {
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(status)
	apiResponse := jsonmodel.APIResponse{HasError: hasError, Message: message, Data: data}
	json.NewEncoder(responseWriter).Encode(apiResponse)
}

func GenerateNewOTP() string {
	number := rand.Intn(100000)
	hd := hashids.NewData()
	hd.Salt = consts.OTPSalt
	hd.MinLength = consts.OTPMinLength
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{number})
	return e
}
