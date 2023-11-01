# Use an official Golang runtime as a parent image
FROM golang:latest
RUN apk add --no-cache git


# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY ["go.mod","go.sum","./"]
RUN go mod dowliad

COPY . .

RUN go build -o backend -v .

# Expose the port the application runs on
EXPOSE 3000

# Run the application
CMD ["./backend"]
