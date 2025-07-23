#!/bin/bash

# Demo script to show what the Lambda function would process

echo "🧪 Lambda Kinesis Demo"
echo "====================="
echo ""
echo "This shows the decoded EventBridge event that would be processed by the Lambda function."
echo ""

# The base64 data from test-event.json decoded
echo "📨 EventBridge Event (decoded from Kinesis record):"
echo "---------------------------------------------------"

# Base64 decode the sample data
echo "ewogICJpZCI6ICIxMjM0NTY3OC0xMjM0LTEyMzQtMTIzNC0xMjM0NTY3ODkwMTIiLAogICJkZXRhaWwtdHlwZSI6ICJDYWxsU3VtbWFyeSIsCiAgInNvdXJjZSI6ICJhd3MucGFydG5lci9vcGVyYXRhLmNvbS9jdXN0b21lci1ldmVudHMiLAogICJhY2NvdW50IjogIjEyMzQ1Njc4OTAxMiIsCiAgInRpbWUiOiAiMjAyNS0wNy0yMlQxMDozMDo0NVoiLAogICJyZWdpb24iOiAidXMtZWFzdC0xIiwKICAiZGV0YWlsIjogewogICAgImNvbnRhY3QiOiB7CiAgICAgICJpZCI6IHsKICAgICAgICAiY3VycmVudCI6ICJjb250YWN0LTEyMzQ1IgogICAgICB9LAogICAgICAiZGlyZWN0aW9uIjogIklOQk9VTkQiLAogICAgICAiZW5kZWRCeSI6ICJBR0VOVCIsCiAgICAgICJxdWV1ZU5hbWUiOiAiU3VwcG9ydCIsCiAgICAgICJjYWxsZXJJRCI6ICIrMTIzNDU2Nzg5MCIKICAgIH0sCiAgICAic2VydmljZUFnZW50IjogewogICAgICAiZnJpZW5kbHlOYW1lIjogIkpvaG4gRG9lIiwKICAgICAgInVzZXJuYW1lIjogImpvaG4uZG9lIiwKICAgICAgImludGVyYWN0aW9uIjogewogICAgICAgICJ0b3RhbER1cmF0aW9uU2VjIjogMTgwLAogICAgICAgICJ0YWxraW5nRHVyYXRpb25TZWMiOiAxNTAsCiAgICAgICAgIm9uSG9sZER1cmF0aW9uU2VjIjogMAogICAgICB9CiAgICB9LAogICAgIndlYlJUQ1Nlc3Npb24iOiB7CiAgICAgICJtZXRyaWNzIjogewogICAgICAgICJpbmJvdW5kIjogewogICAgICAgICAgInBhY2tldHNMb3N0UGVyY2VudGFnZSI6IDAuMDUKICAgICAgICB9LAogICAgICAgICJvdXRib3VuZCI6IHsKICAgICAgICAgICJwYWNrZXRzTG9zdFBlcmNlbnRhZ2UiOiAwLjAzCiAgICAgICAgfSwKICAgICAgICAibW9zIjogewogICAgICAgICAgImF2ZyI6IDQuNQogICAgICAgIH0KICAgICAgfQogICAgfQogIH0KfQ==" | base64 -d | jq .

echo ""
echo "⚡ Lambda Processing Output:"
echo "---------------------------"
echo "When this event is processed by the Lambda function, it would output:"
echo ""
echo "==========================================
Operata Event Received
==========================================
Event ID: 12345678-1234-1234-1234-123456789012
Event Type: Call Summary
Source: aws.partner/operata.com/customer-events
Account: 123456789012
Region: us-east-1
Time: 2025-07-22 10:30:45 UTC
------------------------------------------
Call Summary Details:
  Direction: INBOUND
  Ended By: AGENT
  Queue: Support
  Caller ID: +1234567890
  Agent: John Doe (john.doe)
  Duration: 180 seconds (Medium)
  Talk Time: 150 seconds
  Audio Quality:
    Inbound Packet Loss: 0.05% (Minimal)
    Outbound Packet Loss: 0.03% (Minimal)
    MOS Score: 4.5 (Excellent)
  ✅ Overall Quality: Excellent
=========================================="

echo ""
echo "🚀 To deploy this Lambda function:"
echo "1. Run: ./build.sh"
echo "2. Upload lambda-kinesis.zip to AWS Lambda"
echo "3. Configure Kinesis event source mapping"
echo "4. Monitor CloudWatch Logs for event processing output"
