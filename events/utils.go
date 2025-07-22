package events

import (
	"encoding/json"
	"fmt"
)

// ParseEventBridgeEvent parses a generic EventBridge event and returns the appropriate typed event
func ParseEventBridgeEvent(data []byte) (interface{}, error) {
	var genericEvent EventBridgeEvent
	if err := json.Unmarshal(data, &genericEvent); err != nil {
		return nil, fmt.Errorf("failed to parse generic event: %w", err)
	}

	switch genericEvent.DetailType {
	case "CallSummary":
		var event CallSummaryEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, fmt.Errorf("failed to parse CallSummary event: %w", err)
		}
		return &event, nil

	case "InsightsSummary":
		var event InsightsSummaryEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, fmt.Errorf("failed to parse InsightsSummary event: %w", err)
		}
		return &event, nil

	case "AgentReportedIssue":
		var event AgentReportedIssueEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, fmt.Errorf("failed to parse AgentReportedIssue event: %w", err)
		}
		return &event, nil

	case "HeadsetSummary":
		var event HeadsetSummaryEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, fmt.Errorf("failed to parse HeadsetSummary event: %w", err)
		}
		return &event, nil

	default:
		return &genericEvent, nil
	}
}

// IsOperataEvent checks if an EventBridge event is from Operata based on the source field
func IsOperataEvent(source string) bool {
	// Operata events have source in format: aws.partner/operata.com/...
	prefix := "aws.partner/operata.com/"
	return len(source) > len(prefix) && source[:len(prefix)] == prefix
}

// GetEventTypeFromDetailType maps the detail-type to a more friendly event type name
func GetEventTypeFromDetailType(detailType string) string {
	switch detailType {
	case "CallSummary":
		return "Call Summary"
	case "InsightsSummary":
		return "Insights Summary"
	case "AgentReportedIssue":
		return "Agent Reported Issue"
	case "HeadsetSummary":
		return "Headset Summary"
	default:
		return detailType
	}
}

// CallQualityLevel represents the quality level based on MOS score
type CallQualityLevel string

const (
	QualityExcellent CallQualityLevel = "Excellent"
	QualityGood      CallQualityLevel = "Good"
	QualityFair      CallQualityLevel = "Fair"
	QualityPoor      CallQualityLevel = "Poor"
	QualityBad       CallQualityLevel = "Bad"
)

// GetCallQualityLevel returns the quality level based on MOS score
// MOS (Mean Opinion Score) scale:
// 4.3-5.0: Excellent
// 4.0-4.3: Good
// 3.6-4.0: Fair
// 3.1-3.6: Poor
// 1.0-3.1: Bad
func GetCallQualityLevel(mosScore float64) CallQualityLevel {
	switch {
	case mosScore >= 4.3:
		return QualityExcellent
	case mosScore >= 4.0:
		return QualityGood
	case mosScore >= 3.6:
		return QualityFair
	case mosScore >= 3.1:
		return QualityPoor
	default:
		return QualityBad
	}
}

// PacketLossLevel represents the severity of packet loss
type PacketLossLevel string

const (
	PacketLossMinimal    PacketLossLevel = "Minimal"
	PacketLossAcceptable PacketLossLevel = "Acceptable"
	PacketLossNoticeable PacketLossLevel = "Noticeable"
	PacketLossHigh       PacketLossLevel = "High"
	PacketLossSevere     PacketLossLevel = "Severe"
)

// GetPacketLossLevel returns the severity level based on packet loss percentage
func GetPacketLossLevel(lossPercentage float64) PacketLossLevel {
	switch {
	case lossPercentage < 0.1:
		return PacketLossMinimal
	case lossPercentage < 1.0:
		return PacketLossAcceptable
	case lossPercentage < 3.0:
		return PacketLossNoticeable
	case lossPercentage < 5.0:
		return PacketLossHigh
	default:
		return PacketLossSevere
	}
}

// GetCallDurationCategory categorizes call duration
func GetCallDurationCategory(durationSec int) string {
	switch {
	case durationSec < 30:
		return "Very Short"
	case durationSec < 120:
		return "Short"
	case durationSec < 600:
		return "Medium"
	case durationSec < 1800:
		return "Long"
	default:
		return "Very Long"
	}
}
