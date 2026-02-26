#!/bin/bash

# Build script for the blog - generates tags and builds the site

set -e  # Exit on error

echo "Generating tag pages..."
go run gen-tag-pages.go

echo "Building site with Zola..."
zola build

echo "✓ Build complete!"
