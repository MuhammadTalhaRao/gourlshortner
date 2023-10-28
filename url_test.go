package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var router *gin.Engine

func TestCreateShortURL(t *testing.T) {
	request := createMockRequestForCreateURL()
	writer := makeRequest(t, http.MethodPost, CREATE_SHORT_URL, request, false)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestFetchURLSUCCESS(t *testing.T) {
	url := fmt.Sprintf(TEST_BSAE_URL, CORRECT_CODE)
	writer := makeRequest(t, http.MethodGet, url, nil, false)
	assert.Equal(t, http.StatusMovedPermanently, writer.Code)
}

func TestFetchURLFAILURE(t *testing.T) {
	url := fmt.Sprintf(TEST_BSAE_URL, WRONG_CODE)
	writer := makeRequest(t, http.MethodGet, url, nil, false)
	assert.Equal(t, http.StatusNotFound, writer.Code)
}

/*
it sends a request to endpoint we want to test after creating a router if it already exists it moves on with other operations
*/
func makeRequest(t *testing.T, method, url string, body interface{}, isAuthenticatedRequest bool) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	writer := httptest.NewRecorder()

	if router == nil {
		router = setupTestEnv()
	}

	router.ServeHTTP(writer, request)
	return writer
}

func setupTestEnv() *gin.Engine {
	return setup()
}

func createMockRequestForCreateURL() CreateShotURLRequest {
	return CreateShotURLRequest{
		URL: "https://medium.com/@raotalha302.rt",
	}
}

// MOCK DATA
const (
	TEST_BSAE_URL = "http://localhost:7088/%v"
	CORRECT_CODE  = "AK_zpi6y_BNkjffTYaw6W3y9Vxcg_Mv8ZR_eXP_Ush8="
	WRONG_CODE    = "2131"
)
