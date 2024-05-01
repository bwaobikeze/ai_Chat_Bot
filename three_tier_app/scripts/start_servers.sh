#!/bin/bash

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo "Docker is not installed. Please install Docker before running this script."
    exit 1
fi

# Check if Docker daemon is running
if ! docker info &> /dev/null; then
    echo "Docker daemon is not running. Start the Docker daemon before running this script."
    exit 1
fi

# Run Docker Compose
docker-compose up