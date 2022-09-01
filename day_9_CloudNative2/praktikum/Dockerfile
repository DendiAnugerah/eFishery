FROM golang:alpine

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main
ENTRYPOINT ["./app/main"]

# WORKDIR /app
# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download
    
# COPY . ./

# RUN go build -o main main.go

# EXPOSE 8080
# CMD ["./app/main"]