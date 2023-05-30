package service

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Global SMTP server configuration variable
var smtpConfig SMTPConfig

// Load SMTP server configuration from the YAML file
func LoadSMTPConfig(filename string) error {
    data, err := os.ReadFile(filename)
    if err != nil {
        return err
    }

    err = yaml.Unmarshal(data, &smtpConfig)
    if err != nil {
        return err
    }

    return nil
}