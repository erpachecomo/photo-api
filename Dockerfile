FROM golang:latest

# Set the working directory inside the container
WORKDIR /app
# Copy the Go module files first to leverage Docker's caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Set the working directory to the api-server main package
WORKDIR /app/cmd/api-server

RUN go build -o /app/main .

CMD [ "/app/main" ]