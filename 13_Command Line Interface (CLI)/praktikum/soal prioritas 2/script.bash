#!/bin/bash

validate_project_name() {
    if [[ ! "$1" =~ ^[a-zA-Z0-9_-]+$ ]]; then
        echo "Project name can only contain alphanumeric characters, hyphens, and underscores."
        exit 1
    fi
}

create_simple_project() {
    mkdir "$project_name"
    cd "$project_name" || exit

    go mod init "$project_name"

    touch main.go

    echo "Simple project '$project_name' created successfully."
}

create_api_project() {
    mkdir "$project_name"
    cd "$project_name" || exit

    go mod init "$project_name"

    mkdir routes models repositories services configs controllers

    touch main.go

    echo "API project '$project_name' created successfully."
}

read -rp "Enter project name: " project_name
validate_project_name "$project_name"

echo "Choose project type:"
echo "1. Simple"
echo "2. API"
read -rp "Enter project type (1 or 2): " project_type

case $project_type in
    1) create_simple_project ;;
    2) create_api_project ;;
    *) echo "Invalid project type." ;;
esac

