FROM golang:1.9.4

# Copy the local package files to the container's workspace.
ADD . /go/src/deeplink

RUN mkdir -p /go/src/deeplink
WORKDIR /go/src/deeplink

RUN go get -u github.com/kardianos/govendor
RUN govendor init
RUN govendor fetch github.com/jinzhu/gorm

# Build the deeplink command inside the container.
# (You may fetch or manage dependencies here, either manually or with a tool like "godep".)
RUN go install

# Run the deeplink command by default when the container starts.
ENTRYPOINT /go/bin/deeplink

# Document that the service listens on port 8080.
EXPOSE 8080
