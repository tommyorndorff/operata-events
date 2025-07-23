# Lambda Kinesis Example

This example demonstrates how to create an AWS Lambda function that processes Operata events from a Kinesis stream.

## Overview

The Lambda function:
1. Receives Kinesis events containing Operata EventBridge events
2. Parses each record to extract the EventBridge event data
3. Validates that events are from Operata
4. Uses the Operata events library to parse specific event types
5. Outputs detailed information about each event to stdout (CloudWatch Logs)

## Event Processing

The function handles all Operata event types:
- **CallSummary**: Detailed call metrics including quality, duration, and agent information
- **InsightsSummary**: AI-generated insights about potential issues
- **AgentReportedIssue**: Issues reported by agents with system context
- **HeadsetSummary**: Headset usage and audio quality metrics

## Building and Deployment

### Local Build
```bash
go build -o bootstrap main.go
zip lambda-kinesis.zip bootstrap
```

### Using AWS SAM
```bash
sam build
sam deploy --guided
```

### Direct Upload
Upload the `lambda-kinesis.zip` file to your Lambda function.

## Environment Setup

### IAM Permissions
Your Lambda execution role needs:
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "logs:CreateLogGroup",
                "logs:CreateLogStream",
                "logs:PutLogEvents"
            ],
            "Resource": "arn:aws:logs:*:*:*"
        },
        {
            "Effect": "Allow",
            "Action": [
                "kinesis:DescribeStream",
                "kinesis:DescribeStreamSummary",
                "kinesis:GetRecords",
                "kinesis:GetShardIterator",
                "kinesis:ListShards",
                "kinesis:ListStreams"
            ],
            "Resource": "arn:aws:kinesis:*:*:stream/your-operata-events-stream"
        }
    ]
}
```

### Event Source Mapping
Create an event source mapping to connect your Kinesis stream to the Lambda function:

```bash
aws lambda create-event-source-mapping \
    --function-name operata-events-processor \
    --event-source-arn arn:aws:kinesis:us-east-1:123456789012:stream/operata-events \
    --starting-position LATEST \
    --batch-size 10
```

## Configuration

### Lambda Settings
- **Runtime**: `provided.al2` (for Go binary)
- **Handler**: `bootstrap`
- **Memory**: 128-256 MB (depending on batch size)
- **Timeout**: 30-60 seconds
- **Batch Size**: 10-100 records (adjust based on processing needs)

### Environment Variables
No environment variables are required for basic operation.

## Output Format

The function outputs structured information for each event type:

### CallSummary Example Output
```
==========================================
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
==========================================
```

## Monitoring

### CloudWatch Metrics
Monitor your Lambda function with these key metrics:
- Duration
- Error Rate
- Throttles
- Iterator Age (for Kinesis)

### Alarms
Set up CloudWatch alarms for:
- High error rate (> 1%)
- Long duration (> timeout - 10s)
- Iterator age (> 60 seconds)

### Logs
All event processing output goes to CloudWatch Logs. Search for:
- Event types: "Call Summary", "Insights Summary", etc.
- Quality issues: "Issues detected", "⚠️"
- Errors: "Error parsing", "Error processing"

## Scaling Considerations

### Batch Size
- Smaller batches (10-25): Lower latency, higher Lambda invocation cost
- Larger batches (50-100): Higher latency, lower cost, risk of timeout

### Parallelization
- Each Kinesis shard processes events sequentially
- Add more shards to increase parallelism
- Consider using Kinesis Data Firehose for high-volume scenarios

### Error Handling
- Failed records are automatically retried
- Consider implementing a dead letter queue for persistent failures
- Monitor iterator age to detect processing issues

## Testing

### Local Testing
```bash
go run main.go
# Use AWS SAM for local testing with sample events
sam local invoke -e test-event.json
```

### Sample Test Event
Create a `test-event.json` file with a sample Kinesis event structure.

## Troubleshooting

### Common Issues
1. **Import errors**: Ensure `go mod tidy` was run
2. **Parsing errors**: Check that Kinesis records contain valid JSON
3. **High iterator age**: Check Lambda concurrency limits and processing time
4. **Memory issues**: Increase Lambda memory allocation

### Debug Mode
Add debug logging by setting log level:
```go
log.SetLevel(log.DebugLevel)
```
