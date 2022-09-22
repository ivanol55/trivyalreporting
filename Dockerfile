FROM golang:1.19.1-alpine3.16 AS builder

# Copy the repository into the working directory
WORKDIR /trivyalreporting
COPY . .

# Compile the helper executable
RUN go build helper.go

# Remove .git files to avoid repository disclosure
RUN rm -rf .git .gitignore README.md docker-compose.yaml Dockerfile

# Switch to Alpine image as a final base
FROM alpine:3.16.2 AS helper

# Install AWS CLI to access cloud infrastructure
RUN apk add --no-cache aws-cli

# Install trivy to use as a backend
RUN apk add --no-cache curl && \
	curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.32.0 && \
	apk del curl

# Copy built artifacts from builder container
WORKDIR /trivyalreporting
COPY --from=builder /trivyalreporting ./

ENTRYPOINT ["./helper"]
