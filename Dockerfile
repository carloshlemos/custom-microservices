FROM golang:1.14-alpine
RUN go version
# install git
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
# install glide
RUN go get github.com/Masterminds/glide
# create a working directory
WORKDIR /go/src/app
# add glide.yaml and glide.lock
ADD glide.yaml glide.yaml
ADD glide.lock glide.lock
# update glide
RUN glide --debug update
# install packages
RUN glide install
# add source code
ADD custom-microservice src
# build main.go
RUN go build /custom-microservice/main.go
# run the binary
CMD ["./main"]