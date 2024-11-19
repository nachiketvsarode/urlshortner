package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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

func createURL(originalURL string) {
	shortURL := generateShortURL(originalURL)
	id := shortURL //Use the short URL as the ID for the simplicity
	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	//return shortURL

}

func main() {
	fmt.Println("Starting URL shortner......")
	OriginalURL := "https://github.com/nachiketvsarode"
	generateShortURL(OriginalURL)

}
