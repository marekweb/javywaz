#!/bin/bash
# filepath: /Users/marek/dev/javywazero/build.sh
# Script to download Javy if needed and build WASM modules

set -e

# Determine platform
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
VERSION="v5.0.2"  # Latest version available

# Map OS names to Javy naming conventions
if [[ "$OS" == "darwin"* ]]; then
  OS="macos"
fi

# Map architecture names to Javy naming conventions
if [[ "$ARCH" == "aarch64" ]]; then
  ARCH="arm"
fi

JAVY_BINARY="./javy-${ARCH}-${OS}-${VERSION}"
JAVY_GZ="${JAVY_BINARY}.gz"
DOWNLOAD_URL="https://github.com/bytecodealliance/javy/releases/download/${VERSION}/javy-${ARCH}-${OS}-${VERSION}.gz"

# Only download if we don't already have the binary
if [ ! -f "$JAVY_BINARY" ]; then
  echo "Javy binary not found. Downloading from ${DOWNLOAD_URL}..."
  
  if ! curl -L -o "${JAVY_GZ}" "${DOWNLOAD_URL}"; then
    echo "Failed to download Javy binary"
    exit 1
  fi
  
  gzip -d "${JAVY_GZ}"
  chmod +x "${JAVY_BINARY}"
  echo "Javy binary downloaded successfully"
else
  echo "Using existing Javy binary"
fi

# Build the WASM module
echo "Building example.wasm with Javy..."
"$JAVY_BINARY" build example.js -o example.wasm
echo "Build complete: example.wasm"
