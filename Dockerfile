# Create a stage for building the application.
ARG GO_VERSION=1.23.2
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
#LABEL org.opencontainers.image.source=https://github.com/milosobral/PoolPlanner
WORKDIR /src

# Download dependencies as a separate step to take advantage of Docker's caching.
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

# This is the architecture you're building for, which is passed in by the builder.
# Placing it here allows the previous steps to be cached across architectures.
ARG TARGETARCH

ARG TARGETBIN

# Build the application.
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/${TARGETBIN} cmd/${TARGETBIN}/main.go

################################################################################
FROM alpine:latest AS runtime

#LABEL org.opencontainers.image.source=https://github.com/milosobral/PoolPlanner

ARG TARGETBIN

# Install any runtime dependencies that are needed to run your application.
RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add \
        ca-certificates \
        tzdata \
        && \
        update-ca-certificates

# Create a non-privileged user that the app will run under.
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
   --home "/nonexistent" \
   --shell "/sbin/nologin" \
   --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

# Copy the executable from the "build" stage.
COPY --from=build /bin/${TARGETBIN} /bin/${TARGETBIN}
COPY ./templates ./templates

# Set the entry point to the application
ENV TARGETBINENTRY=/bin/${TARGETBIN}

# What the container should run when it is started.
ENTRYPOINT ${TARGETBINENTRY}
