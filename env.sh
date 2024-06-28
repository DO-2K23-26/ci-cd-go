#!/bin/bash

# Define environment variables
export POSTGRES_USER="gocity"
export POSTGRES_PASSWORD="gocity-pwd"
export POSTGRES_DB="gocity"
export CITY_API_ADDR="localhost"
export CITY_API_PORT="8080"
export CITY_API_DB_URL="localhost"
export CITY_API_DB_USER="gocity"
export CITY_API_DB_PWD="gocity-pwd"
export CITY_API_DB_NAME="gocity"

# Print a message indicating the environment variables have been set
echo "Environment variables set:"
echo "POSTGRES_USER=${POSTGRES_USER}"
echo "POSTGRES_PASSWORD=${POSTGRES_PASSWORD}"
echo "POSTGRES_DB=${POSTGRES_DB}"
echo "CITY_API_ADDR=${CITY_API_ADDR}"
echo "CITY_API_PORT=${CITY_API_PORT}"
echo "CITY_API_DB_URL=${CITY_API_DB_URL}"
echo "CITY_API_DB_USER=${CITY_API_DB_USER}"
echo "CITY_API_DB_PWD=${CITY_API_DB_PWD}"
echo "CITY_API_DB_NAME=${CITY_API_DB_NAME}"
