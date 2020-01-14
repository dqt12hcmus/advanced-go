FROM golang:alpine AS build
RUN apk update && apk add --no-cache git && apk add --no-cache gcc
# Create appuser.
RUN adduser -D -g '' appuser
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/micro-client
############################
# STEP 2 build a small image
############################
FROM scratch
# Import the user and group files from the builder.
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd
# Copy the executable
COPY --from=build /go/bin/micro-client /go/bin/micro-client
USER appuser