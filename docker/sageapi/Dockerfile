# Base image.
FROM golang:1.15.6-alpine

# Create root dir.
WORKDIR /app

# Copies just gigamono-bakend and gigamono-ui/dist directories.
# Check .dockerignore for ignored folders and files.
COPY ./ ./

# Enter working directory.
WORKDIR /app/gigamono-backend

# Download dependencies.
# Build executable.
# Delete dependencies.
RUN go mod download && \
    go build cmd/app/app.go && \
    go clean -modcache

# Add and switch to non-root user.
# After the build because we get permissing error when go copies executables around.
RUN addgroup -S non-root && adduser -S non-root -G non-root
USER non-root

# Specify script to run when container is started.
ENTRYPOINT ["./app"]
