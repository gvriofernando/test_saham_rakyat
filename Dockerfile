# Start with a Go base image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go binary
RUN go build -o main .

# Expose port 5000
EXPOSE 5000

# Command to run the executable
CMD ["./main"]