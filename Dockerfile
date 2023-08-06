FROM golang:1.18

## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /2hf

## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /2hf

COPY go.mod /2hf
COPY go.sum /2hf
RUN go mod download

## We copy everything in the root directory
## into our /app directory
ADD . /app

## we run go build to compile the binary
## executable of our Go program
RUN go build -o main .

## Our start command which kicks off
## our newly created binary executable
CMD ["/2hf/main"]

