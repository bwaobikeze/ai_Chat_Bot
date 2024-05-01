# Use an official Go runtime as a parent image
FROM golang:latest

# Set env variables
ARG CONNECTION_URL=neo4j+s://65320fe1.databases.neo4j.io
ARG USERNAME=neo4j
ARG PASSWORD=jsEZ9L2QhjIC8hKAMsLt3jIUhn5FpVKnuLjCKqZAVVM

# Set the working directory in the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Download and install any required dependencies
RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose the port on which the application will run
EXPOSE 8082

# Run the application
CMD ["./main"]
