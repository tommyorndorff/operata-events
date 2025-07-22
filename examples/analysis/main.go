package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tommyorndorff/operata-events/events"
)

func main() {
	// Example of processing multiple event types
	eventJSONs := []string{
		// CallSummary event
		`{
			"version": "0",
			"id": "call-event-123",
			"detail-type": "CallSummary",
			"source": "aws.partner/operata.com/demo-group/eventBus",
			"account": "123456789",
			"time": "2023-06-01T10:30:00Z",
			"region": "us-east-1",
			"resources": [],
			"detail": {
				"accountProperties": {
					"operataGroupName": "Demo Company",
					"operataGroupId": "demo-group-123"
				},
				"contact": {
					"id": {
						"current": "contact-456"
					},
					"direction": "Inbound",
					"endedBy": "Customer",
					"queueName": "Support Queue",
					"callerId": "+1234567890"
				},
				"webRTCSession": {
					"metrics": {
						"inbound": {
							"packetsReceived": 1200,
							"packetsLost": 50,
							"packetsLostPercentage": 4.17,
							"audioLevel": {"min": 0, "max": 500, "avg": 120}
						},
						"outbound": {
							"packetsSent": 1150,
							"packetsLost": 25,
							"packetsLostPercentage": 2.17,
							"audioLevel": {"min": 0, "max": 800, "avg": 200}
						},
						"mos": {
							"min": 3.2,
							"max": 3.8,
							"avg": 3.5
						}
					}
				},
				"serviceAgent": {
					"username": "support_agent_1",
					"friendlyName": "Alice Support",
					"interaction": {
						"totalDurationSec": 480,
						"talkingDurationSec": 420,
						"onHoldDurationSec": 60,
						"onMuteDurationSec": 0
					}
				}
			}
		}`,

		// InsightsSummary event
		`{
			"version": "0",
			"id": "insight-event-456",
			"detail-type": "InsightsSummary",
			"source": "aws.partner/operata.com/demo-group/eventBus",
			"account": "123456789",
			"time": "2023-06-01T10:32:00Z",
			"region": "us-east-1",
			"resources": [],
			"detail": {
				"accountProperties": {
					"operataGroupId": "demo-group-123"
				},
				"contact": {
					"id": {
						"current": "contact-456"
					}
				},
				"insights": {
					"count": 3,
					"tags": [
						{"description": "High Packet Loss"},
						{"description": "Poor Audio Quality"},
						{"description": "Network Congestion"}
					]
				}
			}
		}`,

		// Non-Operata event
		`{
			"version": "0",
			"id": "other-event-789",
			"detail-type": "S3ObjectCreated",
			"source": "aws.s3",
			"account": "123456789",
			"time": "2023-06-01T10:35:00Z",
			"region": "us-east-1",
			"resources": [],
			"detail": {}
		}`,
	}

	fmt.Println("=== Operata Event Analysis Tool ===")
	fmt.Println()

	for i, eventJSON := range eventJSONs {
		fmt.Printf("--- Processing Event %d ---\n", i+1)

		// Parse the generic event first to check source
		var genericEvent events.EventBridgeEvent
		if err := json.Unmarshal([]byte(eventJSON), &genericEvent); err != nil {
			log.Printf("Failed to parse event %d: %v", i+1, err)
			continue
		}

		// Check if it's an Operata event
		if !events.IsOperataEvent(genericEvent.Source) {
			fmt.Printf("Event ID: %s\n", genericEvent.ID)
			fmt.Printf("Source: %s\n", genericEvent.Source)
			fmt.Printf("Type: %s\n", genericEvent.DetailType)
			fmt.Printf("Status: Not an Operata event - skipping detailed analysis\n\n")
			continue
		}

		// Parse the specific event type
		parsedEvent, err := events.ParseEventBridgeEvent([]byte(eventJSON))
		if err != nil {
			log.Printf("Failed to parse Operata event %d: %v", i+1, err)
			continue
		}

		// Display common information
		fmt.Printf("Event ID: %s\n", genericEvent.ID)
		fmt.Printf("Event Type: %s\n", events.GetEventTypeFromDetailType(genericEvent.DetailType))
		fmt.Printf("Source: %s\n", genericEvent.Source)
		fmt.Printf("Account: %s\n", genericEvent.Account)
		fmt.Printf("Time: %s\n", genericEvent.Time.Format("2006-01-02 15:04:05 MST"))

		// Process based on event type
		switch event := parsedEvent.(type) {
		case *events.CallSummaryEvent:
			analyzeCallSummary(event)

		case *events.InsightsSummaryEvent:
			analyzeInsightsSummary(event)

		case *events.AgentReportedIssueEvent:
			fmt.Printf("Agent Reported Issue Analysis (not implemented in this example)\n")

		case *events.HeadsetSummaryEvent:
			fmt.Printf("Headset Summary Analysis (not implemented in this example)\n")

		default:
			fmt.Printf("Unknown or unsupported Operata event type\n")
		}

		fmt.Println()
	}
}

func analyzeCallSummary(event *events.CallSummaryEvent) {
	detail := event.Detail

	fmt.Printf("\n--- Call Summary Analysis ---\n")

	// Basic call information
	fmt.Printf("Call Direction: %s\n", detail.Contact.Direction)
	fmt.Printf("Ended By: %s\n", detail.Contact.EndedBy)
	fmt.Printf("Queue: %s\n", detail.Contact.QueueName)
	fmt.Printf("Caller ID: %s\n", detail.Contact.CallerID)

	// Agent information
	if detail.ServiceAgent.FriendlyName != "" {
		fmt.Printf("Agent: %s (%s)\n", detail.ServiceAgent.FriendlyName, detail.ServiceAgent.Username)
	}

	// Call duration analysis
	duration := detail.ServiceAgent.Interaction.TotalDurationSec
	fmt.Printf("Call Duration: %d seconds (%s)\n", duration, events.GetCallDurationCategory(duration))
	fmt.Printf("Talk Time: %d seconds (%.1f%%)\n",
		detail.ServiceAgent.Interaction.TalkingDurationSec,
		float64(detail.ServiceAgent.Interaction.TalkingDurationSec)/float64(duration)*100)

	if detail.ServiceAgent.Interaction.OnHoldDurationSec > 0 {
		fmt.Printf("Hold Time: %d seconds (%.1f%%)\n",
			detail.ServiceAgent.Interaction.OnHoldDurationSec,
			float64(detail.ServiceAgent.Interaction.OnHoldDurationSec)/float64(duration)*100)
	}

	// WebRTC quality analysis
	metrics := detail.WebRTCSession.Metrics

	// Packet loss analysis
	inboundLoss := metrics.Inbound.PacketsLostPercentage
	outboundLoss := metrics.Outbound.PacketsLostPercentage

	fmt.Printf("\n--- Audio Quality Analysis ---\n")
	fmt.Printf("Inbound Packet Loss: %.2f%% (%s)\n",
		inboundLoss, events.GetPacketLossLevel(inboundLoss))
	fmt.Printf("Outbound Packet Loss: %.2f%% (%s)\n",
		outboundLoss, events.GetPacketLossLevel(outboundLoss))

	// MOS score analysis
	mosScore := metrics.MOS.Avg
	quality := events.GetCallQualityLevel(mosScore)
	fmt.Printf("MOS Score: %.2f (%s)\n", mosScore, quality)

	// Overall quality assessment
	fmt.Printf("\n--- Quality Assessment ---\n")
	if quality == events.QualityExcellent &&
		events.GetPacketLossLevel(inboundLoss) == events.PacketLossMinimal &&
		events.GetPacketLossLevel(outboundLoss) == events.PacketLossMinimal {
		fmt.Printf("Overall Quality: Excellent - No issues detected\n")
	} else if quality == events.QualityGood &&
		events.GetPacketLossLevel(inboundLoss) <= events.PacketLossAcceptable &&
		events.GetPacketLossLevel(outboundLoss) <= events.PacketLossAcceptable {
		fmt.Printf("Overall Quality: Good - Minor issues may be present\n")
	} else {
		fmt.Printf("Overall Quality: Issues detected - Review recommended\n")

		// Specific recommendations
		if events.GetPacketLossLevel(inboundLoss) >= events.PacketLossNoticeable {
			fmt.Printf("⚠️  High inbound packet loss detected\n")
		}
		if events.GetPacketLossLevel(outboundLoss) >= events.PacketLossNoticeable {
			fmt.Printf("⚠️  High outbound packet loss detected\n")
		}
		if quality <= events.QualityPoor {
			fmt.Printf("⚠️  Poor audio quality detected\n")
		}
	}
}

func analyzeInsightsSummary(event *events.InsightsSummaryEvent) {
	detail := event.Detail

	fmt.Printf("\n--- Insights Summary Analysis ---\n")
	fmt.Printf("Contact ID: %s\n", detail.Contact.ID.Current)
	fmt.Printf("Issues Detected: %d\n", detail.Insights.Count)

	if detail.Insights.Count > 0 {
		fmt.Printf("Issues:\n")
		for i, tag := range detail.Insights.Tags {
			fmt.Printf("  %d. %s\n", i+1, tag.Description)
		}

		// Categorize severity based on issue count and types
		if detail.Insights.Count >= 3 {
			fmt.Printf("Severity: High - Multiple issues detected\n")
		} else if detail.Insights.Count == 2 {
			fmt.Printf("Severity: Medium - Some issues detected\n")
		} else {
			fmt.Printf("Severity: Low - Minor issue detected\n")
		}
	} else {
		fmt.Printf("No issues detected\n")
	}
}
