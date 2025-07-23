package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/tommyorndorff/operata-events/events"
)

func main() {
	// Example CallSummary event JSON (simplified)
	callSummaryJSON := `{
		"version": "0",
		"id": "530848f3-1111-2222-3333-b33ba70c19f0",
		"detail-type": "CallSummary",
		"source": "aws.partner/operata.com/a28453f9-1111-2222-3333-84d9e67ac297/eventBus",
		"account": "083560837128",
		"time": "2023-06-01T05:00:13Z",
		"region": "ap-southeast-2",
		"resources": [],
		"detail": {
			"accountProperties": {
				"operataGroupName": "Demo Group",
				"operataGroupId": "a28453f9-1111-2222-3333-84d9e67ac297"
			},
			"contact": {
				"id": {
					"current": "ac7a6a89-1111-2222-3333-1e659475d24e"
				},
				"direction": "Inbound",
				"events": {
					"connectingToAgent": "2023-06-01T04:59:51.523Z",
					"enqueued": "2023-06-01T04:59:30.795Z"
				},
				"endedBy": "Agent",
				"queueName": "Default Queue",
				"callerId": "+61402960149"
			},
			"webRTCSession": {
				"metrics": {
					"inbound": {
						"packetsReceived": 636,
						"packetsLost": 12,
						"packetsLostPercentage": 1.85,
						"bytesReceived": 67178,
						"audioLevel": {
							"min": 0,
							"max": 400,
							"avg": 51.23
						},
						"jitterBufferMils": {
							"min": 0,
							"max": 8,
							"avg": 3.5
						}
					},
					"outbound": {
						"packetsSent": 786,
						"packetsLost": 10,
						"packetsLostPercentage": 1.27,
						"bytesSent": 65585,
						"audioLevel": {
							"min": 0,
							"max": 4515,
							"avg": 618.62
						},
						"jitterBufferMils": {
							"min": 4,
							"max": 26,
							"avg": 11.23
						}
					},
					"mos": {
						"min": 3.65,
						"max": 4.43,
						"avg": 4.25
					}
				},
				"usedDevices": []
			},
			"serviceAgent": {
				"username": "demo-agent",
				"machine": {
					"cpu": {
						"modelName": "Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz",
						"idlePercentage": {
							"avg": 69.89
						},
						"utilisedPercentage": {
							"min": 28.53,
							"max": 31.84,
							"avg": 30.11
						}
					},
					"memory": {
						"availableGb": 16,
						"utilisedPercentage": {
							"min": 92.27,
							"max": 92.49,
							"avg": 92.31
						}
					}
				},
				"network": {
					"internetGatewayIp": "103.120.49.101",
					"mediaIpAddress": "192.168.1.12",
					"type": "wlan",
					"isp": "Example ISP",
					"geolocation": {
						"city": "Sydney",
						"region": "NSW",
						"country": "Australia"
					}
				},
				"browser": {
					"name": "Chrome",
					"version": "113.0.5672.126"
				},
				"interaction": {
					"totalDurationSec": 45,
					"onHoldDurationSec": 0,
					"talkingDurationSec": 42,
					"onMuteDurationSec": 3
				},
				"friendlyName": "Demo Agent"
			},
			"billing": {
				"durationRoundedMin": 1
			},
			"timestamp": "2023-06-01T05:00:11.871Z"
		}
	}`

	// Parse the CallSummary event
	var callEvent events.CallSummaryEvent
	if err := json.Unmarshal([]byte(callSummaryJSON), &callEvent); err != nil {
		log.Fatalf("Failed to parse CallSummary event: %v", err)
	}

	// Display key information from the call
	fmt.Println("=== Call Summary Event ===")
	fmt.Printf("Event ID: %s\n", callEvent.ID)
	fmt.Printf("Event Type: %s\n", callEvent.DetailType)
	fmt.Printf("Account: %s\n", callEvent.Account)
	fmt.Printf("Time: %s\n", callEvent.Time.Format("2006-01-02 15:04:05"))

	// Call details
	fmt.Printf("\n--- Call Details ---\n")
	fmt.Printf("Direction: %s\n", callEvent.Detail.Contact.Direction)
	fmt.Printf("Ended by: %s\n", callEvent.Detail.Contact.EndedBy)
	fmt.Printf("Queue: %s\n", callEvent.Detail.Contact.QueueName)
	fmt.Printf("Caller ID: %s\n", callEvent.Detail.Contact.CallerID)

	// Agent information
	fmt.Printf("\n--- Agent Information ---\n")
	fmt.Printf("Agent: %s (%s)\n", callEvent.Detail.ServiceAgent.FriendlyName, callEvent.Detail.ServiceAgent.Username)
	fmt.Printf("Browser: %s %s\n", callEvent.Detail.ServiceAgent.Browser.Name, callEvent.Detail.ServiceAgent.Browser.Version)
	fmt.Printf("Location: %s, %s, %s\n",
		callEvent.Detail.ServiceAgent.Network.Geolocation.City,
		callEvent.Detail.ServiceAgent.Network.Geolocation.Region,
		callEvent.Detail.ServiceAgent.Network.Geolocation.Country)

	// Call metrics
	fmt.Printf("\n--- Call Metrics ---\n")
	fmt.Printf("Total Duration: %d seconds\n", callEvent.Detail.ServiceAgent.Interaction.TotalDurationSec)
	fmt.Printf("Talk Time: %d seconds\n", callEvent.Detail.ServiceAgent.Interaction.TalkingDurationSec)
	fmt.Printf("Hold Time: %d seconds\n", callEvent.Detail.ServiceAgent.Interaction.OnHoldDurationSec)
	fmt.Printf("Mute Time: %d seconds\n", callEvent.Detail.ServiceAgent.Interaction.OnMuteDurationSec)

	// WebRTC Quality metrics
	fmt.Printf("\n--- WebRTC Quality ---\n")
	inbound := callEvent.Detail.WebRTCSession.Metrics.Inbound
	outbound := callEvent.Detail.WebRTCSession.Metrics.Outbound
	mos := callEvent.Detail.WebRTCSession.Metrics.MOS

	fmt.Printf("Inbound Packet Loss: %.2f%% (%d/%d packets)\n",
		inbound.PacketsLostPercentage, inbound.PacketsLost, inbound.PacketsReceived)
	fmt.Printf("Outbound Packet Loss: %.2f%% (%d/%d packets)\n",
		outbound.PacketsLostPercentage, outbound.PacketsLost, outbound.PacketsSent)
	fmt.Printf("MOS Score: %.2f (min: %.2f, max: %.2f)\n",
		mos.Avg, mos.Min, mos.Max)

	// Machine performance
	fmt.Printf("\n--- Machine Performance ---\n")
	cpu := callEvent.Detail.ServiceAgent.Machine.CPU
	memory := callEvent.Detail.ServiceAgent.Machine.Memory
	fmt.Printf("CPU: %s\n", cpu.ModelName)
	fmt.Printf("CPU Utilization: %.1f%% (avg)\n", cpu.UtilisedPercentage.Avg)
	fmt.Printf("Memory Available: %.0f GB\n", memory.AvailableGB)
	fmt.Printf("Memory Utilization: %.1f%% (avg)\n", memory.UtilisedPercentage.Avg)

	// Example InsightsSummary event
	fmt.Println("\n" + strings.Repeat("=", 50))

	insightJSON := `{
		"version": "0",
		"id": "e063ddf9-ac05-08ee-138c-7105a7d3c04b",
		"detail-type": "InsightsSummary",
		"source": "aws.partner/operata.com/a28453f9-c9d3-4c48-a7cd-84d9e67ac297/eventBus",
		"account": "083560837128",
		"time": "2023-06-01T05:00:13Z",
		"region": "ap-southeast-2",
		"resources": [],
		"detail": {
			"accountProperties": {
				"operataGroupId": "a28453f9-c9d3-4c48-a7cd-84d9e67ac297"
			},
			"contact": {
				"id": {
					"current": "ac7a6a89-43ff-4dce-8aa2-1e659475d24e"
				}
			},
			"insights": {
				"count": 2,
				"tags": [
					{
						"description": "High Packet Loss"
					},
					{
						"description": "Poor Audio Quality"
					}
				]
			}
		}
	}`

	var insightEvent events.InsightsSummaryEvent
	if err := json.Unmarshal([]byte(insightJSON), &insightEvent); err != nil {
		log.Fatalf("Failed to parse InsightsSummary event: %v", err)
	}

	fmt.Println("=== Insights Summary Event ===")
	fmt.Printf("Event ID: %s\n", insightEvent.ID)
	fmt.Printf("Contact ID: %s\n", insightEvent.Detail.Contact.ID.Current)
	fmt.Printf("Insights Found: %d\n", insightEvent.Detail.Insights.Count)
	fmt.Println("Issues Detected:")
	for i, tag := range insightEvent.Detail.Insights.Tags {
		fmt.Printf("  %d. %s\n", i+1, tag.Description)
	}
}
