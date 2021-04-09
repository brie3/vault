# Initial stage: download modules
FROM golang:1.16 as modules

ADD go.mod go.sum /m/
RUN cd /m && go mod download

# Intermediate stage: Build the binary
FROM golang:1.16 as builder

COPY --from=modules /go/pkg /go/pkg

# Add a non-privileged user
RUN useradd -u 10001 app_user

RUN mkdir -p /app
ADD . /app
WORKDIR /app

# Build the binary with go build
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build -o ./bin/app .

# Final stage: Run the binary
FROM scratch

# Password from previous stage
COPY --from=builder /etc/passwd /etc/passwd
USER app_user

# Build the binary
COPY --from=builder /app/bin/app /app

CMD ["/app"]