# Use the official Go image
FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

# Install Tesseract OCR and development headers
RUN set -eux \
    && sed -i '/yarn/d' /etc/apt/sources.list /etc/apt/sources.list.d/* 2>/dev/null || true \
    && apt-get update \
    && apt-get install -y --no-install-recommends curl gnupg \
    && apt-get install -y --no-install-recommends \
        tesseract-ocr \
        libtesseract-dev \
        tesseract-ocr-eng \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /workspaces/gw2-translator
