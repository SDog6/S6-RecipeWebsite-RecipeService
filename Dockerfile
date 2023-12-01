FROM golang:1.20.11-alpine as builder
WORKDIR /S6-RecipeWebsite
COPY go.* ./
RUN go mod download
COPY . ./
RUN go build -o main .
ENV port=9000
EXPOSE 9000
CMD ["go", "run", "main.go"]