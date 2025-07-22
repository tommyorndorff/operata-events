package events

import (
	"encoding/json"
	"testing"
	"time"
)

// Test data based on the examples from Operata documentation
func TestCallSummaryEventUnmarshal(t *testing.T) {
	jsonData := `{
		"version": "0",
		"id": "530848f3-1111-2222-3333-b33ba70c19f0",
		"detail-type": "CallSummary",
		"source": "aws.partner/operata.com/a28453f9-1111-2222-3333-84d9e67ac297/andyEventBus",
		"account": "083560837128",
		"time": "2023-06-01T05:00:13Z",
		"region": "ap-southeast-2",
		"resources": [],
		"detail": {
			"accountProperties": {
				"operataGroupName": "Operata Demo",
				"operataGroupId": "a28453f9-1111-2222-3333-84d9e67ac297"
			},
			"contact": {
				"id": {
					"current": "ac7a6a89-1111-2222-3333-1e659475d24e",
					"previous": "",
					"next": ""
				},
				"direction": "Inbound",
				"events": {
					"connectingToAgent": "2023-06-01T04:59:51.523Z",
					"enqueued": "2023-06-01T04:59:30.795Z"
				},
				"endedBy": "Agent",
				"queueName": "Operata Prod Default Queue",
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
							"avg": 3
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
					"rtt": {
						"min": 0,
						"max": 145,
						"avg": 110
					},
					"jitter": {
						"min": 0,
						"max": 8,
						"avg": 3
					},
					"mos": {
						"min": 3.65,
						"max": 4.43,
						"avg": 4.25
					}
				},
				"serviceEndpoint": {
					"fqdn": "",
					"transportLifeTimeSeconds": 0,
					"expiry": "0001-01-01T00:00:00Z"
				},
				"mediaEndpoint": {
					"fqdn": "turnnlb-93f2de0c97c4316b.elb.ap-southeast-2.amazonaws.com.",
					"destinationPort": "3478",
					"sourcePort": "49985",
					"transport": "udp",
					"privateIp": "10.4.3.108"
				},
				"signallingEndpoint": {
					"fqdn": ""
				},
				"usedDevices": [
					{
						"timestamp": "2023-06-01T04:59:55.041Z",
						"deviceId": "default",
						"groupId": "293e6a9871f0d56112233445566773d1e36f9e6ed9f4926c5a9318336ecfeec",
						"kind": "audioinput",
						"label": "Default - Elgato Wave:3 (0fd9:0070)"
					}
				]
			},
			"serviceAgent": {
				"username": "andy",
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
					"isp": "Bcd Networks Pty Ltd",
					"geolocation": {
						"city": "Nutfield",
						"region": "Victoria",
						"country": "Australia"
					}
				},
				"browser": {
					"name": "Chrome",
					"version": "113.0.5672.126"
				},
				"softphone": {
					"softphoneUrl": "https://operata-prod.awsapps.com/connect/ccp-v2/softphone#ac7a6a89-43ff-4dce-8aa2-1e659475d24e",
					"softphoneContextUrl": "https://operata-prod.awsapps.com/connect/ccp-v2/softphone#ac7a6a89-43ff-4dce-8aa2-1e659475d24e"
				},
				"interaction": {
					"totalDurationSec": 14,
					"onHoldDurationSec": 0,
					"talkingDurationSec": 14,
					"onMuteDurationSec": 0
				},
				"friendlyName": "Andy"
			},
			"billing": {
				"durationRoundedMin": 1
			},
			"timestamp": "2023-06-01T05:00:11.871Z"
		}
	}`

	var event CallSummaryEvent
	err := json.Unmarshal([]byte(jsonData), &event)
	if err != nil {
		t.Fatalf("Failed to unmarshal CallSummaryEvent: %v", err)
	}

	// Verify basic event structure
	if event.Version != "0" {
		t.Errorf("Expected version '0', got '%s'", event.Version)
	}
	if event.DetailType != "CallSummary" {
		t.Errorf("Expected detail-type 'CallSummary', got '%s'", event.DetailType)
	}

	// Verify detail content
	if event.Detail.Contact.EndedBy != "Agent" {
		t.Errorf("Expected call ended by 'Agent', got '%s'", event.Detail.Contact.EndedBy)
	}
	if event.Detail.ServiceAgent.Username != "andy" {
		t.Errorf("Expected username 'andy', got '%s'", event.Detail.ServiceAgent.Username)
	}
	if event.Detail.ServiceAgent.Interaction.TotalDurationSec != 14 {
		t.Errorf("Expected total duration 14 seconds, got %d", event.Detail.ServiceAgent.Interaction.TotalDurationSec)
	}
}

func TestInsightsSummaryEventUnmarshal(t *testing.T) {
	jsonData := `{
		"version": "0",
		"id": "e063ddf9-ac05-08ee-138c-7105a7d3c04b",
		"detail-type": "InsightsSummary",
		"source": "aws.partner/operata.com/a28453f9-c9d3-4c48-a7cd-84d9e67ac297/andyEventBus",
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
				"count": 1,
				"tags": [
					{
						"description": "High Packet Loss"
					}
				]
			}
		}
	}`

	var event InsightsSummaryEvent
	err := json.Unmarshal([]byte(jsonData), &event)
	if err != nil {
		t.Fatalf("Failed to unmarshal InsightsSummaryEvent: %v", err)
	}

	if event.DetailType != "InsightsSummary" {
		t.Errorf("Expected detail-type 'InsightsSummary', got '%s'", event.DetailType)
	}
	if event.Detail.Insights.Count != 1 {
		t.Errorf("Expected insights count 1, got %d", event.Detail.Insights.Count)
	}
	if len(event.Detail.Insights.Tags) != 1 {
		t.Errorf("Expected 1 insight tag, got %d", len(event.Detail.Insights.Tags))
	}
	if event.Detail.Insights.Tags[0].Description != "High Packet Loss" {
		t.Errorf("Expected tag description 'High Packet Loss', got '%s'", event.Detail.Insights.Tags[0].Description)
	}
}

func TestHeartbeatWorkflowEventUnmarshal(t *testing.T) {
	jsonData := `[
		{
			"agentId": "kanchan",
			"agentType": "",
			"axScore": 10,
			"createdOn": "2022-03-31T05:40:57.967565Z",
			"cxScore": 10,
			"diallerCallId": "CA355bf611223344557f2b9f420d77a52841",
			"groupId": "5bbbd17a-1111-2222-3333-c4674035f406",
			"heartbeatId": "7ac7ba31-1111-2222-3333-6796c5d0f2e5",
			"jobId": "1648705257967728403",
			"networkScore": 10,
			"receiverCallId": "5981c50a-1111-2222-3333-b65004a777f2",
			"routingProfileId": "",
			"status": "Call Not Accepted by Agent",
			"toNumber": "",
			"updatedOn": "2022-03-31T05:41:36.130184Z"
		}
	]`

	var events HeartbeatWorkflowEvents
	err := json.Unmarshal([]byte(jsonData), &events)
	if err != nil {
		t.Fatalf("Failed to unmarshal HeartbeatWorkflowEvents: %v", err)
	}

	if len(events) != 1 {
		t.Errorf("Expected 1 heartbeat event, got %d", len(events))
	}

	event := events[0]
	if event.AgentID != "kanchan" {
		t.Errorf("Expected agent ID 'kanchan', got '%s'", event.AgentID)
	}
	if event.Status != "Call Not Accepted by Agent" {
		t.Errorf("Expected status 'Call Not Accepted by Agent', got '%s'", event.Status)
	}
	if event.AxScore != 10 {
		t.Errorf("Expected AxScore 10, got %d", event.AxScore)
	}

	// Test time parsing
	expectedTime, _ := time.Parse(time.RFC3339, "2022-03-31T05:40:57.967565Z")
	if !event.CreatedOn.Equal(expectedTime) {
		t.Errorf("Expected created time %v, got %v", expectedTime, event.CreatedOn)
	}
}
