FROM golang:1.14-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh gcc libc-dev

RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/githubnemo/CompileDaemon

# Set the Current Working Directory inside the container
WORKDIR /go/src/projetoapi

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
#CMD ["./main"]
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./main.go" -command="./main"