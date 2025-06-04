FROM golang:1.24-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

COPY src /app/src

# Build the Go app
RUN go build -o main /app/src

# Expose port specified in .env file
ENV DB_HOST=host.docker.internal
EXPOSE ${SERVER_PORT}

# Command to run the executable
CMD ["./main"]

