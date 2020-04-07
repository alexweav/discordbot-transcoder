# Build

FROM rust:1.41-alpine AS builder

# Work outside of $GOPATH, since we're using modules.
WORKDIR /app

# Copy source to image.
COPY . .

# Get deps, clean, and build.
RUN cargo install --path .

# Create export directory and copy binaries.
RUN mkdir /export && \
    cp target/* /export

# Package

FROM alpine:latest

# Create application directory.
RUN mkdir -p /opt/app/bin

# Get exported binaries.
COPY --from=builder /export /opt/app/bin

# Launch application.
ENTRYPOINT /opt/app/bin/release/discordbot-transcoder
