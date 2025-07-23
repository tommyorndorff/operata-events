package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	operataEvents "github.com/tommyorndorff/operata-events/events"
)

// KinesisLambdaHandler handles Kinesis events containing Operata EventBridge events
func KinesisLambdaHandler(ctx context.Context, kinesisEvent events.KinesisEvent) error {
	log.Printf("Processing %d Kinesis records", len(kinesisEvent.Records))

	for i, record := range kinesisEvent.Records {
		log.Printf("Processing record %d/%d", i+1, len(kinesisEvent.Records))

		// Decode the Kinesis record data
		recordData := record.Kinesis.Data

		// Parse as EventBridge event
		var eventBridgeEvent operataEvents.EventBridgeEvent
		if err := json.Unmarshal(recordData, &eventBridgeEvent); err != nil {
			log.Printf("Error parsing EventBridge event from record %d: %v", i+1, err)
			continue
		}

		// Check if it's an Operata event
		if !operataEvents.IsOperataEvent(eventBridgeEvent.Source) {
			log.Printf("Skipping non-Operata event from source: %s", eventBridgeEvent.Source)
			continue
		}

		// Parse the specific event type
		parsedEvent, err := operataEvents.ParseEventBridgeEvent(recordData)
		if err != nil {
			log.Printf("Error parsing Operata event: %v", err)
			continue
		}

		// Process based on event type
		if err := processOperataEvent(eventBridgeEvent, parsedEvent); err != nil {
			log.Printf("Error processing event: %v", err)
			continue
		}
	}

	return nil
}

// processOperataEvent processes a parsed Operata event and writes details to stdout
func processOperataEvent(genericEvent operataEvents.EventBridgeEvent, parsedEvent interface{}) error {
	// Common event information
	fmt.Printf("==========================================\n")
	fmt.Printf("Operata Event Received\n")
	fmt.Printf("==========================================\n")
	fmt.Printf("Event ID: %s\n", genericEvent.ID)
	fmt.Printf("Event Type: %s\n", operataEvents.GetEventTypeFromDetailType(genericEvent.DetailType))
	fmt.Printf("Source: %s\n", genericEvent.Source)
	fmt.Printf("Account: %s\n", genericEvent.Account)
	fmt.Printf("Region: %s\n", genericEvent.Region)
	fmt.Printf("Time: %s\n", genericEvent.Time.Format("2006-01-02 15:04:05 MST"))
	fmt.Printf("------------------------------------------\n")

	// Process specific event types
	switch event := parsedEvent.(type) {
	case *operataEvents.CallSummaryEvent:
		return processCallSummaryEvent(event)

	case *operataEvents.InsightsSummaryEvent:
		return processInsightsSummaryEvent(event)

	case *operataEvents.AgentReportedIssueEvent:
		return processAgentReportedIssueEvent(event)

	case *operataEvents.HeadsetSummaryEvent:
		return processHeadsetSummaryEvent(event)

	default:
		fmt.Printf("Unknown event type - raw data:\n")
		if jsonData, err := json.MarshalIndent(parsedEvent, "", "  "); err == nil {
			fmt.Printf("%s\n", jsonData)
		}
	}

	fmt.Printf("==========================================\n\n")
	return nil
}

// processCallSummaryEvent processes a CallSummary event
func processCallSummaryEvent(event *operataEvents.CallSummaryEvent) error {
	detail := event.Detail

	fmt.Printf("Call Summary Details:\n")
	fmt.Printf("  Direction: %s\n", detail.Contact.Direction)
	fmt.Printf("  Ended By: %s\n", detail.Contact.EndedBy)
	fmt.Printf("  Queue: %s\n", detail.Contact.QueueName)
	fmt.Printf("  Caller ID: %s\n", detail.Contact.CallerID)

	// Agent information
	if detail.ServiceAgent.FriendlyName != "" {
		fmt.Printf("  Agent: %s (%s)\n", detail.ServiceAgent.FriendlyName, detail.ServiceAgent.Username)
	}

	// Call duration analysis
	duration := detail.ServiceAgent.Interaction.TotalDurationSec
	fmt.Printf("  Duration: %d seconds (%s)\n", duration, operataEvents.GetCallDurationCategory(duration))
	fmt.Printf("  Talk Time: %d seconds\n", detail.ServiceAgent.Interaction.TalkingDurationSec)

	if detail.ServiceAgent.Interaction.OnHoldDurationSec > 0 {
		fmt.Printf("  Hold Time: %d seconds\n", detail.ServiceAgent.Interaction.OnHoldDurationSec)
	}

	// Quality metrics
	metrics := detail.WebRTCSession.Metrics
	inboundLoss := metrics.Inbound.PacketsLostPercentage
	outboundLoss := metrics.Outbound.PacketsLostPercentage
	mosScore := metrics.MOS.Avg

	fmt.Printf("  Audio Quality:\n")
	fmt.Printf("    Inbound Packet Loss: %.2f%% (%s)\n",
		inboundLoss, operataEvents.GetPacketLossLevel(inboundLoss))
	fmt.Printf("    Outbound Packet Loss: %.2f%% (%s)\n",
		outboundLoss, operataEvents.GetPacketLossLevel(outboundLoss))
	fmt.Printf("    MOS Score: %.2f (%s)\n", mosScore, operataEvents.GetCallQualityLevel(mosScore))

	// Quality assessment
	quality := operataEvents.GetCallQualityLevel(mosScore)
	if quality == operataEvents.QualityExcellent &&
		operataEvents.GetPacketLossLevel(inboundLoss) == operataEvents.PacketLossMinimal &&
		operataEvents.GetPacketLossLevel(outboundLoss) == operataEvents.PacketLossMinimal {
		fmt.Printf("  ✅ Overall Quality: Excellent\n")
	} else if quality <= operataEvents.QualityPoor ||
		operataEvents.GetPacketLossLevel(inboundLoss) >= operataEvents.PacketLossNoticeable ||
		operataEvents.GetPacketLossLevel(outboundLoss) >= operataEvents.PacketLossNoticeable {
		fmt.Printf("  ⚠️  Overall Quality: Issues detected\n")
	} else {
		fmt.Printf("  ✓ Overall Quality: Good\n")
	}

	return nil
}

// processInsightsSummaryEvent processes an InsightsSummary event
func processInsightsSummaryEvent(event *operataEvents.InsightsSummaryEvent) error {
	detail := event.Detail

	fmt.Printf("Insights Summary Details:\n")
	fmt.Printf("  Contact ID: %s\n", detail.Contact.ID.Current)
	fmt.Printf("  Issues Detected: %d\n", detail.Insights.Count)

	if detail.Insights.Count > 0 {
		fmt.Printf("  Issues:\n")
		for i, tag := range detail.Insights.Tags {
			fmt.Printf("    %d. %s\n", i+1, tag.Description)
		}

		// Severity assessment
		if detail.Insights.Count >= 3 {
			fmt.Printf("  🔴 Severity: High - Multiple issues detected\n")
		} else if detail.Insights.Count == 2 {
			fmt.Printf("  🟡 Severity: Medium - Some issues detected\n")
		} else {
			fmt.Printf("  🟢 Severity: Low - Minor issue detected\n")
		}
	} else {
		fmt.Printf("  ✅ No issues detected\n")
	}

	return nil
}

// processAgentReportedIssueEvent processes an AgentReportedIssue event
func processAgentReportedIssueEvent(event *operataEvents.AgentReportedIssueEvent) error {
	detail := event.Detail

	fmt.Printf("Agent Reported Issue Details:\n")
	fmt.Printf("  Agent: %s\n", detail.Agent)
	fmt.Printf("  Issue ID: %s\n", detail.ID)
	fmt.Printf("  State: %s\n", detail.State)
	fmt.Printf("  Timestamp: %s\n", detail.Timestamp.Format("2006-01-02 15:04:05 MST"))

	// Issue context
	fmt.Printf("  Issue Details:\n")
	fmt.Printf("    Category: %s\n", detail.Context.Category)
	fmt.Printf("    Cause: %s\n", detail.Context.Cause)
	fmt.Printf("    Severity: %s\n", detail.Context.Severity)
	fmt.Printf("    Message: %s\n", detail.Context.Message)

	// System information
	fmt.Printf("  System Info:\n")
	fmt.Printf("    CPU: %s (%.1f%% used)\n", detail.System.CPU.ModelName, detail.System.CPU.UsedPercentage)
	fmt.Printf("    Memory: %.1f GB available / %.1f GB total\n", detail.System.Memory.Available, detail.System.Memory.Total)
	fmt.Printf("    Browser: %s %s\n", detail.Browser.Name, detail.Browser.Version)

	// Softphone error
	if detail.SoftphoneError.Type != "" {
		fmt.Printf("  Softphone Error:\n")
		fmt.Printf("    Type: %s\n", detail.SoftphoneError.Type)
		if detail.SoftphoneError.Message != "" {
			fmt.Printf("    Message: %s\n", detail.SoftphoneError.Message)
		}
	}

	return nil
}

// processHeadsetSummaryEvent processes a HeadsetSummary event
func processHeadsetSummaryEvent(event *operataEvents.HeadsetSummaryEvent) error {
	detail := event.Detail

	fmt.Printf("Headset Summary Details:\n")
	fmt.Printf("  Contact ID: %s\n", detail.Contact.ID.Current)
	fmt.Printf("  Interaction Duration: %d seconds\n", detail.Contact.Interaction.TotalDurationSec)

	// Headset information
	headset := detail.Headset
	fmt.Printf("  Headset Info:\n")
	fmt.Printf("    Model: %s\n", headset.ModelName)
	fmt.Printf("    Firmware: %s\n", headset.FirmwareVersion)
	fmt.Printf("    Serial: %s\n", headset.SerialNumber)
	fmt.Printf("    API Version: %s\n", headset.APIVersion)

	// Speech metrics
	speech := headset.Metrics.Speech
	fmt.Printf("  Speech Metrics:\n")
	fmt.Printf("    Total Duration: %.1f seconds\n", speech.TotalSeconds)
	fmt.Printf("    TX Speech: %.1f seconds (%.1f%%)\n", speech.TxSpeechTotal, speech.TxSpeechTotalPct)
	fmt.Printf("    RX Speech: %.1f seconds (%.1f%%)\n", speech.RxSpeechTotal, speech.RxSpeechTotalPct)
	fmt.Printf("    Silence: %.1f seconds (%.1f%%)\n", speech.SilenceTotal, speech.SilenceTotalPct)
	fmt.Printf("    Cross Talk: %.1f seconds (%.1f%%)\n", speech.CrossTalkTotal, speech.CrossTalkTotalPct)

	// Audio quality
	fmt.Printf("  Audio Quality:\n")
	fmt.Printf("    Exposure: %.1f dB (avg)\n", headset.Metrics.ExposureDB.Avg)
	fmt.Printf("    Background Noise: %.1f dB (avg)\n", headset.Metrics.BackgroundNoiseDB.Avg)

	// Device usage
	fmt.Printf("  Device Usage:\n")
	fmt.Printf("    Misaligned Boom Arm: %d times\n", headset.Metrics.MisalignedBoomArmCount)
	fmt.Printf("    Device Mute: %d times\n", headset.Metrics.DeviceMuteCount)
	fmt.Printf("    Volume Adjustments: %d times\n", headset.Metrics.DeviceVolumeAdjustCount)

	return nil
}

func main() {
	lambda.Start(KinesisLambdaHandler)
}
