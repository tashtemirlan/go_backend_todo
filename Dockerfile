# Use the official Golang image as the base
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o material_todo_go

# Expose the port your Go app runs on
EXPOSE 2525

# Command to run the executable
CMD ["./material_todo_go"]
