# Build

FROM rust:1.41-stretch AS builder

# Work outside of $GOPATH, since we're using modules.
WORKDIR /app

# Copy source to image.
COPY . .

# muslc is required in order to build the rust image.
RUN apt-get update && apt-get -y install ca-certificates cmake musl-tools libssl-dev && rm -rf /var/lib/apt/lists/*

# Get deps, clean, and build.
RUN rustup target add x86_64-unknown-linux-musl
RUN cargo build --target x86_64-unknown-linux-musl --release

# Create export directory and copy binaries.
RUN mkdir /export && \
    cp -r target/* /export

# Package

FROM alpine:latest

# Create application directory.
RUN mkdir -p /opt/app/bin

# Get exported binaries.
COPY --from=builder /export /opt/app/bin

# Launch application.
ENTRYPOINT /opt/app/bin/x86_64-unknown-linux-musl/release/discordbot-transcoder
