FROM golang:1.18.3-alpine as builder
WORKDIR /S6-RecipeWebsite
COPY go.* ./
RUN go mod download
COPY . ./
RUN go build -o main .
ENV port=8080
EXPOSE 8080
CMD ["go", "run", "main.go"]