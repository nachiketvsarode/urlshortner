package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

/*
how the map would look like ?

	p3244424 ----> {
						ID : "p3244424"
						OriginalURL: "http://github.com/nachiketvsarode"
						ShortURL: "p3244424"
						CreationDate: time.Now()
					}

*/

func generateShortURL(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL)) // this will converts the originalURL string to a byte slice
	data := hasher.Sum(nil)
	fmt.Println("hasher data:", data)
	hash := hex.EncodeToString(data)
	fmt.Println("Encode to string :", hash)
	fmt.Println("Final string :", hash[:8])
	return hash[:8]
}

func createURL(originalURL string) string {
	shortURL := generateShortURL(originalURL)
	id := shortURL //Use the short URL as the ID for the simplicity
	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL

}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

func main() {
	fmt.Println("Starting URL shortner......")
	OriginalURL := "https://github.com/nachiketvsarode"
	generateShortURL(OriginalURL)

	// Start the HTTP Server on port 8080
	fmt.Println("Starting server on port 3000.....")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error on starting server")
	}

}
