// Package events provides Go structs for Operata's EventBridge event catalog.
//
// This package contains type definitions for all event types published by Operata
// to Amazon EventBridge, including:
//
//   - CallSummary events: Delivered after call completion with comprehensive
//     call metrics, WebRTC session data, and agent information
//   - InsightsSummary events: Generated when insights are established for a call
//   - AgentReportedIssue events: Created when agents report issues through the system
//   - HeadsetSummary events: Delivered when headset statistics collection is enabled
//   - HeartbeatWorkflow events: Generated from heartbeat test workflows
//
// All events follow the standard EventBridge event structure with a common header
// and event-specific detail payload.
//
// Example usage:
//
//	import "github.com/tommyorndorff/operata-events/events"
//
//	// Parse a CallSummary event
//	var event events.CallSummaryEvent
//	if err := json.Unmarshal(data, &event); err != nil {
//		// handle error
//	}
//
//	// Access event details
//	fmt.Printf("Call ended by: %s\n", event.Detail.Contact.EndedBy)
//	fmt.Printf("Call duration: %d seconds\n", event.Detail.ServiceAgent.Interaction.TotalDurationSec)
//
// For more information about Operata's event catalog, see:
// https://docs.operata.com/docs/event-catalog
package events
