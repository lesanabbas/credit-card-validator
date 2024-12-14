# Stage 1: Build the Go binary
FROM golang:1.21 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire project and build the binary
COPY . ./
RUN go build -o main .

# Stage 2: Run the application
FROM golang:1.21 as runner

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main /app/main

# Copy the templates folder
COPY --from=builder /app/templates /app/templates

# Expose port 8080 (or your API port)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
