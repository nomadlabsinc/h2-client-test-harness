# Use an official Go runtime as a parent image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Install openssl which is required for generating self-signed certificates
RUN apk add --no-cache openssl

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the local package source code to the container
COPY . .

# Build the Go app
RUN go build -o /h2-client-test-harness

# Expose port 8080 to the outside world
EXPOSE 8080

# The command to run when the container starts.
# The test case will be passed as an argument to `docker run`.
ENTRYPOINT ["/h2-client-test-harness"]
CMD ["--help"]
