package main

import "sync"

type CreateShotURLRequest struct {
	URL string `json:"url" binding:"required"`
}

// Error
type ErrorResp struct {
	ErrorMsg string `json:"error"`
}

// Success
type SuccessResp struct {
	ShortCode string `json:"shortCode"`
}

// map[shortcode]{url}
type URLStore map[string]string

type URLSTORE struct {
	Store URLStore
	mutex sync.Mutex
}
