FROM golang:alpine AS build
WORKDIR /app
COPY . .
COPY .env ./
RUN go mod download
RUN go build -o main .
EXPOSE 8081
EXPOSE 2112
EXPOSE 8080 
CMD ["./main"]
