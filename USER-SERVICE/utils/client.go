package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
    "taqsir/model"
)

func SendNotification(notification model.Cation) {
	url := "http://localhost:9001/order/notify"
	jsonValue, _ := json.Marshal(notification)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Printf("Failed to send notification: %v", err)
		return
	}
	defer resp.Body.Close()
	log.Println("Notification sent successfully")
}
