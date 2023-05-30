package main

import (
	"log"
	"net/http"

	"github.com/fadilahonespot/email-service/service"
)

func main() {

    err := service.LoadSMTPConfig("config.yaml")
    if err != nil {
        log.Fatalf("Failed to load SMTP configuration: %v", err)
    }

    // Define the endpoint for sending the email
    http.HandleFunc("/api/email", service.SendEmailHandler)

    // Start the HTTP server
    log.Println("Server listening on http://localhost:8011")
    log.Fatal(http.ListenAndServe(":8011", nil))
}