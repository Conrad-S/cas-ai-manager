# Use the official Golang image as the base
FROM golang:1.21.1

# Set the working directory inside the container
WORKDIR /app

# Copy the Go app files to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080 for the application
EXPOSE 8080

# Command to run the application
CMD ["./main"]
