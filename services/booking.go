package services

import (
	"app/db"
	"app/models"
	"app/utils"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-resty/resty/v2"
)

// Cached token and expiry
var cachedToken string
var tokenExpiry time.Time

// Replace with your actual Reloadly credentials
var clientID = "ZYL4r3nC3l8PnHTaH0fpqP1kFWWcnP4e"
var clientSecret = "KEqByd56HN-ktGnHb9pAYxOOQCVzQf-iNuwNfUbK8U1dsoLkaNn4KpKqCw5g7YI"

// GetAccessToken retrieves and caches a Reloadly access token
func GetAccessToken() (string, error) {
	//Check if token is still valid
	if cachedToken != "" && time.Now().Before(tokenExpiry) {
		return cachedToken, nil
	}

	client := resty.New()
	url := "https://auth.reloadly.com/oauth/token"

	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"client_id":     clientID,
			"client_secret": clientSecret,
			"grant_type":    "client_credentials",
			"audience":      "https://giftcards-sandbox.reloadly.com",
		}).
		Post(url)

	if err != nil {
		return "", err
	}

	if resp.StatusCode() != 200 {
		return "", errors.New("failed to fetch access token")

	}
	// Debug: print raw body
	//fmt.Println("Raw response body:", resp.String())

	var result struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return "", err
	}

	// Cache the token and expiry
	cachedToken = result.AccessToken
	tokenExpiry = time.Now().Add(time.Duration(result.ExpiresIn-60) * time.Second) // buffer

	//fmt.Println("cachedToken:", cachedToken)
	return cachedToken, nil
}

func Order(data *models.EventBooking) (interface{}, error) {
	if data.ClientName == "" {
		return "", errors.New("client name is required")
	}
	if data.Date == "" {
		return "", errors.New("date is required")
	}
	if data.EventType == "" {
		return "", errors.New("event type is required")
	}
	if data.MusicianType == "" {
		return "", errors.New("musician is required")
	}

	booking_id, err := utils.GenerateBookingNumber(7)
	if err != nil {
		return "", errors.New("Unabale to create booking number")
	}

	//Checking if order number already exists
	_, err = db.IdExists(booking_id)
	if err != nil {
		return "", err
	}

	data.OrderNumber = booking_id

	//Adding order or booking to a database
	err = db.Order(data)
	if err != nil {
		return "", err
	}
	return data, nil
}
