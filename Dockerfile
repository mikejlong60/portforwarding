# Use a minimal base image (Alpine is small and efficient)
FROM alpine:latest

# Install iptables and dependencies
RUN apk add --no-cache iptables curl

# Download kubectl binary
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" \
    && chmod +x kubectl \
    && mv kubectl /usr/local/bin/

# Set default command to keep the container running
CMD ["sh", "-c", "while true; do sleep 3600; done"]
