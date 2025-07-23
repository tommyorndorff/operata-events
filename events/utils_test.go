package events

import (
	"testing"
)

func TestParseEventBridgeEvent(t *testing.T) {
	callSummaryJSON := `{
		"version": "0",
		"id": "test-id",
		"detail-type": "CallSummary",
		"source": "aws.partner/operata.com/test/eventBus",
		"account": "123456789",
		"time": "2023-06-01T05:00:13Z",
		"region": "us-east-1",
		"resources": [],
		"detail": {
			"accountProperties": {
				"operataGroupId": "test-group"
			},
			"contact": {
				"id": {
					"current": "test-contact"
				}
			},
			"serviceAgent": {
				"interaction": {
					"totalDurationSec": 60
				}
			}
		}
	}`

	event, err := ParseEventBridgeEvent([]byte(callSummaryJSON))
	if err != nil {
		t.Fatalf("Failed to parse event: %v", err)
	}

	callEvent, ok := event.(*CallSummaryEvent)
	if !ok {
		t.Fatalf("Expected CallSummaryEvent, got %T", event)
	}

	if callEvent.DetailType != "CallSummary" {
		t.Errorf("Expected detail-type 'CallSummary', got '%s'", callEvent.DetailType)
	}
}

func TestIsOperataEvent(t *testing.T) {
	tests := []struct {
		source   string
		expected bool
	}{
		{"aws.partner/operata.com/test/eventBus", true},
		{"aws.partner/operata.com/a28453f9-1111-2222-3333-84d9e67ac297/andyEventBus", true},
		{"aws.s3", false},
		{"aws.ec2", false},
		{"operata.com", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsOperataEvent(test.source)
		if result != test.expected {
			t.Errorf("IsOperataEvent(%q) = %v, expected %v", test.source, result, test.expected)
		}
	}
}

func TestGetEventTypeFromDetailType(t *testing.T) {
	tests := []struct {
		detailType string
		expected   string
	}{
		{"CallSummary", "Call Summary"},
		{"InsightsSummary", "Insights Summary"},
		{"AgentReportedIssue", "Agent Reported Issue"},
		{"HeadsetSummary", "Headset Summary"},
		{"UnknownType", "UnknownType"},
	}

	for _, test := range tests {
		result := GetEventTypeFromDetailType(test.detailType)
		if result != test.expected {
			t.Errorf("GetEventTypeFromDetailType(%q) = %q, expected %q", test.detailType, result, test.expected)
		}
	}
}

func TestGetCallQualityLevel(t *testing.T) {
	tests := []struct {
		mosScore float64
		expected CallQualityLevel
	}{
		{4.5, QualityExcellent},
		{4.3, QualityExcellent},
		{4.2, QualityGood},
		{4.0, QualityGood},
		{3.8, QualityFair},
		{3.6, QualityFair},
		{3.4, QualityPoor},
		{3.1, QualityPoor},
		{2.8, QualityBad},
		{1.0, QualityBad},
	}

	for _, test := range tests {
		result := GetCallQualityLevel(test.mosScore)
		if result != test.expected {
			t.Errorf("GetCallQualityLevel(%.1f) = %v, expected %v", test.mosScore, result, test.expected)
		}
	}
}

func TestGetPacketLossLevel(t *testing.T) {
	tests := []struct {
		lossPercentage float64
		expected       PacketLossLevel
	}{
		{0.05, PacketLossMinimal},
		{0.5, PacketLossAcceptable},
		{1.5, PacketLossNoticeable},
		{4.0, PacketLossHigh},
		{6.0, PacketLossSevere},
	}

	for _, test := range tests {
		result := GetPacketLossLevel(test.lossPercentage)
		if result != test.expected {
			t.Errorf("GetPacketLossLevel(%.1f) = %v, expected %v", test.lossPercentage, result, test.expected)
		}
	}
}

func TestGetCallDurationCategory(t *testing.T) {
	tests := []struct {
		durationSec int
		expected    string
	}{
		{15, "Very Short"},
		{60, "Short"},
		{300, "Medium"},
		{900, "Long"},
		{2000, "Very Long"},
	}

	for _, test := range tests {
		result := GetCallDurationCategory(test.durationSec)
		if result != test.expected {
			t.Errorf("GetCallDurationCategory(%d) = %q, expected %q", test.durationSec, result, test.expected)
		}
	}
}
