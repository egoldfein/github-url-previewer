package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// LinkPreviewResponse is type returned by linkpreview API
type LinkPreviewResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	URL         string `json:"url"`
}

// Client is the http client used to call the linkpreview API
type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

// GetPreview generates link preview link
func GetPreview(linkPreviewAccessKey string, url string) string {
	linkPreviewURL := fmt.Sprintf("http://api.linkpreview.net/?key=%v&q=%v", linkPreviewAccessKey, url)
	req, err := http.NewRequest(http.MethodGet, linkPreviewURL, nil)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	var response LinkPreviewResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Fatalf("Error decoding json: %v", err)
	}

	return fmt.Sprintf("<b>%s</b>: %s (%s)", response.Title, response.Description, response.URL)
}
