# Start with a Go base image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go binary
RUN go build -o main .
ENV REDIS_ADDRESS=localhost:6379
ENV REDIS_PASSWORD=
ENV REDIS_DATABASE=0

ENV PG_HOST=localhost 
ENV PG_USER=postgres 
ENV PG_PASSWORD=password 
ENV PG_PORT=5432
ENV PG_DBNAME=test_saham_rakyat

# Expose port 5000
EXPOSE 5000

# Command to run the executable
CMD ["./main"]