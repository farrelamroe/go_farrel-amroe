#!/bin/bash

# Function to send request to API
send_request() {
    local endpoint="$1"
    local method="$2"
    local body="$3"
    local content_type="$4"

    case $method in
        "GET")
            response=$(curl -s -X GET "$endpoint")
            ;;
        "POST")
            response=$(curl -s -X POST -H "Content-Type: $content_type" -d "$body" "$endpoint")
            ;;
        "PUT")
            response=$(curl -s -X PUT -H "Content-Type: $content_type" -d "$body" "$endpoint")
            ;;
        "DELETE")
            response=$(curl -s -X DELETE "$endpoint")
            ;;
        *)
            echo "Invalid HTTP method."
            exit 1
            ;;
    esac

    echo "$response"
}

# Main script

# Get user input for endpoint, HTTP method, request body, and request body type
read -rp "Enter API endpoint: " endpoint
read -rp "Enter HTTP method (GET/POST/PUT/DELETE): " method
read -rp "Enter request body: " body
read -rp "Enter request body type (e.g., application/json): " content_type

# Send request to API
response=$(send_request "$endpoint" "$method" "$body" "$content_type")

# Display response
echo "Response from API:"
echo "$response"

