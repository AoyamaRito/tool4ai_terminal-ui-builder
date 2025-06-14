#!/bin/bash

# Build script for Tool4AI Terminal UI Builder
# Builds cross-platform binaries

set -e

echo "🚀 Building Tool4AI Terminal UI Builder..."

# Create releases directory
mkdir -p releases

# Build for multiple platforms
echo "📦 Building cross-platform binaries..."

# Linux
echo "  Building for Linux AMD64..."
GOOS=linux GOARCH=amd64 go build -o releases/terminal-ui-builder-linux-amd64 terminal-ui-builder.go

echo "  Building for Linux ARM64..."
GOOS=linux GOARCH=arm64 go build -o releases/terminal-ui-builder-linux-arm64 terminal-ui-builder.go

# macOS
echo "  Building for macOS AMD64..."
GOOS=darwin GOARCH=amd64 go build -o releases/terminal-ui-builder-darwin-amd64 terminal-ui-builder.go

echo "  Building for macOS ARM64 (Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -o releases/terminal-ui-builder-darwin-arm64 terminal-ui-builder.go

# Windows
echo "  Building for Windows AMD64..."
GOOS=windows GOARCH=amd64 go build -o releases/terminal-ui-builder-windows-amd64.exe terminal-ui-builder.go

echo "  Building for Windows ARM64..."
GOOS=windows GOARCH=arm64 go build -o releases/terminal-ui-builder-windows-arm64.exe terminal-ui-builder.go

# Build local binary for current platform
echo "  Building for current platform..."
go build -o releases/terminal-ui-builder terminal-ui-builder.go

echo ""
echo "✅ Build complete! Binaries available in releases/"
echo ""
echo "📋 Available binaries:"
ls -la releases/

echo ""
echo "🧪 Testing local binary..."
./releases/terminal-ui-builder --help | head -3

echo ""
echo "🎉 All done! Ready for distribution."