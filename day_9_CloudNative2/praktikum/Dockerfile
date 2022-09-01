FROM golang:alpine as builder
WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main .

CMD ["./main"]




# WORKDIR /app
# COPY . .
# RUN go mod tidy
# RUN go build -o main
# ENTRYPOINT ["./app/main"]

# WORKDIR /app
# RUN go mod download
    
# COPY . ./

# RUN go build -o main main.go

# EXPOSE 8080