package main

import (
	"fmt"
	"log"
)

const (
	DATA_NOT_FOUND = "data not found"
)

// it creates STD errors
func createErrResp(msg string) *ErrorResp {
	log.Printf("Error Occured:  %v", fmt.Sprintf("Message: %v", msg))

	return &ErrorResp{
		ErrorMsg: msg,
	}
}

// create success reponses
func createSuccessResp(shortCode string) SuccessResp {
	return SuccessResp{
		ShortCode: shortCode,
	}
}
