# Operata Events Go Module

[![CI](https://github.com/tommyorndorff/operata-events/workflows/CI/badge.svg)](https://github.com/tommyorndorff/operata-events/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/tommyorndorff/operata-events)](https://goreportcard.com/report/github.com/tommyorndorff/operata-events)
[![GoDoc](https://godoc.org/github.com/tommyorndorff/operata-events?status.svg)](https://godoc.org/github.com/tommyorndorff/operata-events)
[![Release](https://img.shields.io/github/release/tommyorndorff/operata-events.svg)](https://github.com/tommyorndorff/operata-events/releases/latest)

This Go module provides type definitions for Operata's EventBridge event catalog. It includes structs for all event types published by Operata to Amazon EventBridge.

## Installation

```bash
go get github.com/tommyorndorff/operata-events
```

## Event Types

This module includes support for the following Operata event types:

### CallSummary Events
Delivered after call completion with comprehensive metrics including:
- Contact information and call flow
- WebRTC session metrics (packet loss, jitter, MOS scores)
- Service agent details (machine specs, network info, browser)
- Interaction durations and billing information

### InsightsSummary Events
Generated when insights are established for a call, containing:
- Account and contact information
- Insight tags and descriptions

### AgentReportedIssue Events
Created when agents report issues through the system, including:
- Issue context (category, cause, severity)
- System information (CPU, memory usage)
- Browser and softphone error details

### HeadsetSummary Events
Delivered when headset statistics collection is enabled, containing:
- Headset device information
- Speech metrics (crosstalk, silence, etc.)
- Audio quality measurements

### HeartbeatWorkflow Events
Generated from heartbeat test workflows with:
- Agent and call identification
- Quality scores (AX, CX, Network)
- Test status and timing

## Usage

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    
    "github.com/tommyorndorff/operata-events/events"
)

func main() {
    // Example: Parse a CallSummary event
    var callEvent events.CallSummaryEvent
    if err := json.Unmarshal(eventData, &callEvent); err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Call ended by: %s\n", callEvent.Detail.Contact.EndedBy)
    fmt.Printf("Duration: %d seconds\n", 
        callEvent.Detail.ServiceAgent.Interaction.TotalDurationSec)
    fmt.Printf("Packet loss: %.2f%%\n", 
        callEvent.Detail.WebRTCSession.Metrics.Inbound.PacketsLostPercentage)
    
    // Example: Parse an InsightsSummary event
    var insightEvent events.InsightsSummaryEvent
    if err := json.Unmarshal(insightData, &insightEvent); err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Insights count: %d\n", insightEvent.Detail.Insights.Count)
    for _, tag := range insightEvent.Detail.Insights.Tags {
        fmt.Printf("- %s\n", tag.Description)
    }
}
```

## Event Structure

All events follow the standard EventBridge event structure:

```go
type EventBridgeEvent struct {
    Version    string      `json:"version"`
    ID         string      `json:"id"`
    DetailType string      `json:"detail-type"`
    Source     string      `json:"source"`
    Account    string      `json:"account"`
    Time       time.Time   `json:"time"`
    Region     string      `json:"region"`
    Resources  []string    `json:"resources"`
    Detail     interface{} `json:"detail"`
}
```

Each event type has a specific detail payload structure. See the individual struct definitions for complete field documentation.

## Testing

Run the test suite to verify the structs work correctly with example data:

```bash
go test ./events
```

## Documentation

For more information about Operata's event catalog and field descriptions, see:
- [Operata Event Catalog Documentation](https://docs.operata.com/docs/event-catalog)

## Contributing

This module is automatically generated based on Operata's published event schemas. For updates or corrections, please refer to the official Operata documentation.

### Conventional Commits

This project uses [Conventional Commits](https://www.conventionalcommits.org/) for commit messages. This enables automatic semantic versioning and changelog generation.

#### Commit Format

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

#### Types

- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation only changes
- `style`: Changes that do not affect the meaning of the code
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `perf`: A code change that improves performance
- `test`: Adding missing tests or correcting existing tests
- `build`: Changes that affect the build system or external dependencies
- `ci`: Changes to our CI configuration files and scripts
- `chore`: Other changes that don't modify src or test files

#### Scopes

- `events`: Changes to event struct definitions
- `examples`: Changes to example code
- `utils`: Changes to utility functions
- `ci`: Changes to CI/CD configuration
- `deps`: Changes to dependencies

#### Examples

```bash
feat(events): add support for new CallAnalytics event type
fix(utils): correct validation logic for event source
docs: update README with new installation instructions
feat!: remove deprecated fields from CallSummary struct

BREAKING CHANGE: Legacy field removed from CallSummary struct
```

#### Development Tools

Install development tools:

```bash
make install-tools
```

Setup git commit message template:

```bash
make setup-git-hooks
```

Get help with commit format:

```bash
make commit-help
```

Validate your last commit:

```bash
make validate-commit
```

### Semantic Versioning

This project follows [Semantic Versioning](https://semver.org/):

- **MAJOR** version when you make incompatible API changes (breaking changes)
- **MINOR** version when you add functionality in a backwards compatible manner
- **PATCH** version when you make backwards compatible bug fixes

Releases are automatically created when commits are pushed to the main branch using [semantic-release](https://semantic-release.gitbook.io/).

### Release Process

1. Create a feature branch: `git checkout -b feat/new-feature`
2. Make your changes with conventional commit messages
3. Push the branch: `git push origin feat/new-feature`
4. Create a pull request
5. After merge to main, a release will be automatically created if appropriate

### CI/CD

The project includes GitHub Actions workflows for:

- **CI**: Runs tests, linting, and validation on PRs and pushes
- **Release**: Automatically creates releases and updates changelog
- **Commit Lint**: Validates commit messages follow conventional format

## License

See LICENSE file for details.
Operata helper library 
