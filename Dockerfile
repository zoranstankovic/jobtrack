# Multi-stage build: the first stage has the whole Go toolchain and compiles
# the app; the second stage contains only the finished binary. The final image
# is a few MB instead of several hundred.

# ---- Stage 1: build ----
FROM golang:1.27-alpine AS build
WORKDIR /src

# Download dependencies first, in their own layer. Docker caches this layer and
# reuses it as long as go.mod and go.sum don't change, so code edits don't
# trigger a full re-download.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# templ generation is part of the build: *_templ.go files are not in git.
RUN go tool templ generate

# CGO_ENABLED=0 builds a fully static binary that runs without any C libraries.
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/jobtrack ./cmd/jobtrack

# ---- Stage 2: run ----
# distroless/static has no shell and no package manager, just enough to run a
# static binary. ":nonroot" runs it as an unprivileged user.
FROM gcr.io/distroless/static-debian13:nonroot
COPY --from=build /out/jobtrack /jobtrack
EXPOSE 8080
ENTRYPOINT ["/jobtrack"]
