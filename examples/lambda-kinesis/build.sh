#!/bin/bash

# Build script for AWS Lambda deployment

set -e

echo "Building Lambda function for AWS deployment..."

# Build for Linux (Lambda runtime)
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bootstrap main.go

# Create deployment package
zip -r lambda-kinesis.zip bootstrap

echo "✅ Built successfully!"
echo "📦 Deployment package: lambda-kinesis.zip"
echo ""
echo "To deploy:"
echo "1. Upload lambda-kinesis.zip to your Lambda function"
echo "2. Set runtime to 'provided.al2'"
echo "3. Set handler to 'bootstrap'"
echo "4. Configure Kinesis event source mapping"
