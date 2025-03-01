# Start from golang base image
FROM golang:alpine as builder

# Install different timezone
RUN apk add --no-cache tzdata

# Install git.
# Git is required for fetching the dependencies and pull the project details from docker.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /accessApp

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN go mod tidy

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest

RUN apk add --no-cache ca-certificates

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /accessApp/main .

#for time zone
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

#if the above timezone setting will not work then uncomment below
#COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
#ENV ZONEINFO=/zoneinfo.zip
# Expose http port 9001 
EXPOSE 8085

#Command to run the executable
CMD ["./main"]

