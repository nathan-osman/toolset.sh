# Build the application in a separate container
FROM golang:latest
ENV CGO_ENABLED=0
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /usr/local/bin/toolset.sh

# Create the final container with only the binary
FROM scratch
COPY --from=0 /usr/local/bin/toolset.sh /usr/local/bin/
ENTRYPOINT ["toolset.sh"]
