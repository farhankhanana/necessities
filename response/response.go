package response

import (
	"fmt"
	"strings"

	"github.com/gat/necessities/logger"
)

// Response parent object for response
// use this as all-purpose response object
type Response struct {
	ResponseData interface{}
	StatusCode   int
}

// JsonResponse defined as json object for response.
type JsonResponse struct {
	Data      interface{} `json:"data,omitempty"`
	TotalData *int64      `json:"total_data,omitempty"`
	Message   string      `json:"message,omitempty"`
	ErrorCode string      `json:"error_code,omitempty"`
	Success   bool        `json:"success"`
}

// GenerateJsonResponse generates json response object.
//
// Deprecated: this function no longer updated or supported, and will be deleted in version 3.x.x.
// Use normal declaration of `JsonResponse` struct instead.
func GenerateJsonResponse(success bool, message string, data interface{}, totalData *int64, errorCode string) JsonResponse {
	response := JsonResponse{}

	response.Success = success
	response.Message = message
	response.Data = data
	response.TotalData = totalData
	response.ErrorCode = errorCode

	return response
}

// GenerateJsonErrorResponse generates json that specifically for error response.
// In order to use this function properly, inside this function need Database client object and
// Redis client object. So, after Database and Redis client initiated, set `DBGorm` and `Redis` client object
// from this module. Otherwise, the function will generate empty response message.
// Priority is Database client over Redis client.
//
// If `message` is empty, the function will start get data from Database or Redis. If message is given, the given message
// will be used for the response.
//
// Parameter 'data' is used for pass data to the response.
//
// If `parseTemplateData` is nil, some error messages that needs parse/format data might not be parsed properly. So, if
// the error message needs data, pass needed data here.
//
// `errorCode` error code want to be searched.
//
// `locale` currently only support EN/ID. If empty is given, ID will be selected.
func GenerateJsonErrorResponse(message string, data interface{}, parseTemplateData interface{}, errorCode, locale string, setup ...Setup) JsonResponse {
	logger := logger.NewLogger("")
	response := JsonResponse{}
	var err error

	if len(locale) <= 0 {
		locale = strings.ToUpper("id")
	} else {
		locale = strings.ToUpper(locale)
	}

	if len(errorCode) < 1 {
		errorCode = RCUnknownError
	}

	response.Success = false
	if len(message) > 0 {
		response.Message = message
	} else {
		var errorData *Error
		if len(setup) > 0 {
			errorMessageSource := NewErrorMessageSource(setup...)
			errorData, err = errorMessageSource.GetError(errorCode, locale, parseTemplateData)
		} else {
			errorData, err = defaultErrorMessageSource.GetError(errorCode, locale, parseTemplateData)
		}
		if err != nil {
			logger.LogWarn("get error description", err)
			response.Message = ""
		} else {
			response.Message = fmt.Sprintf("%v", errorData.Descriptions[locale])
		}
	}
	response.Data = data
	response.TotalData = nil
	response.ErrorCode = errorCode

	return response
}
