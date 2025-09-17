# Use official Go image as base
FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod ./
# RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go app
RUN go build -o server .

# Expose port 80
EXPOSE 80

# Run the app
CMD ["./server"]

