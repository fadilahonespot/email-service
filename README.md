# email-service

#Request
```
curl --location 'http://localhost:8011/api/email' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "admin@example.com",
    "phone": "0887830000000",
    "name": "Hildan"
}'
```