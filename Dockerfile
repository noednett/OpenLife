##
## Build
##

# Create build stage based on buster image
FROM golang:1.16-buster AS build
# Create working directory under /app
WORKDIR /app
# Copy over all go config (go.mod, go.sum etc.)
COPY . ./
# Install any required modules
RUN go mod download

#install required github packages
RUN go get github.com/gin-gonic/gin

# Run the Go build and output binary. This also forces go to use all go implementations of net libraries. this is required, since alpine does not feature them. alternatively, if only net was used, we can use CGO_ENABLED=0 to achieve the same, where we just disable CGO.
RUN go build -o /app-backend -tags netgo -a -v

##
## Deploy
##

FROM alpine:latest

WORKDIR /

COPY --from=build /app-backend /app-backend

# Run the app binary when we run the container
ENTRYPOINT ["/app-backend"]
