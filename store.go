package main

import (
	"errors"
	"log"
)

// GlobalStore
var Store *URLSTORE

// init store
func NewStore() {
	Store = &URLSTORE{
		Store: make(URLStore),
	}
}

// adds urls in in-memory data structure by keeping shortcode as key
func (store *URLSTORE) AddURLToStore(url, shortCode string) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.Store[shortCode] = url
	log.Printf("Stored URL %v with ShortCode %v", url, shortCode)

}

// retrives values from im-memory datastrcture based on shortcode key
func (store *URLSTORE) FetchURLFromStore(shortCode string) (url string, err error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	url, exist := store.Store[shortCode]
	if !exist {
		return url, errors.New(DATA_NOT_FOUND)
	}

	log.Printf("Fetched URL %v for ShortCode %v", url, shortCode)

	return url, nil
}
