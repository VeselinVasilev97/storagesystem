# Use the official Golang image as a build stage
FROM golang:1.21.6 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o storagesystem

# Use a minimal Docker image to run the application
FROM alpine:latest

# Install necessary packages for running the Go binary
RUN apk add --no-cache libc6-compat

# Set the working directory inside the container
WORKDIR /root/

# Copy the built Go binary from the builder stage
COPY --from=builder /app/storagesystem .

# Copy the configuration file from the builder stage
COPY --from=builder /app/configuration/config.json ./config.json

# Expose the port the application runs on
EXPOSE 8000

# Set environment variables (optional, but usually better passed at runtime)
ENV JWT_SECRET_KEY=xYz1Kbg7pLv5mD1qQ0udFJ0EdRXlNkD7ZphcBQ8KpR8=
# ENV DB_USER=default_user
# ENV DB_PASSWORD=default_password

# Command to run the application
CMD ["./storagesystem"]
