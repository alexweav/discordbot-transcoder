# Build

FROM golang:1.13-alpine AS builder

# Work outside of $GOPATH, since we're using modules.
WORKDIR /app

# Copy source to image.
COPY . .

# Get deps, clean, and build.
RUN go mod download && \
    go clean && \
    go install

# Create export directory and copy binaries.
RUN mkdir /export && \
    cp $GOPATH/bin/* /export

# Package

FROM alpine:latest

# Create application directory.
RUN mkdir -p /opt/app/bin

# Get exported binaries.
COPY --from=builder /export /opt/app/bin

# Launch application.
ENTRYPOINT /opt/app/bin/discordbot-transcoder
