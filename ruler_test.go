package gapi

import (
	"testing"
)

const (
	getGrafanaAlertRulesResponse = `
{
	"Namespace 1": [
		{
			"name": "Kafka Alert",
			"interval": "1m",
			"rules": [
				{
					"expr": "",
					"for": "5m",
					"labels": {
						"environment": "ci",
						"severity": "Warning"
					},
					"grafana_alert": {
						"id": 5,
						"orgId": 1,
						"title": "Kafka Alert - CI",
						"condition": "kafka-count-condition",
						"data": [
							{
								"refId": "zookeeper-count",
								"queryType": "",
								"relativeTimeRange": {
									"from": 600,
									"to": 0
								},
								"datasourceUid": "wU--V0d7k",
								"model": {
									"exemplar": true,
									"expr": "count(zookeeper_status_quorumsize{job=\"zookeeper\",env=\"ci\"})",
									"hide": false,
									"interval": "",
									"intervalMs": 1000,
									"legendFormat": "zookeeper-node-count",
									"maxDataPoints": 43200,
									"refId": "zookeeper-count"
								}
							},
							{
								"refId": "kafka-count",
								"queryType": "",
								"relativeTimeRange": {
									"from": 600,
									"to": 0
								},
								"datasourceUid": "wU--V0d7k",
								"model": {
									"exemplar": true,
									"expr": "count(kafka_server_replicamanager_leadercount{job=\"kafka\",env=\"ci\"})",
									"hide": false,
									"interval": "",
									"intervalMs": 1000,
									"legendFormat": "kafka-node-count",
									"maxDataPoints": 43200,
									"refId": "kafka-count"
								}
							},
							{
								"refId": "kafka-count-condition",
								"queryType": "",
								"relativeTimeRange": {
									"from": 0,
									"to": 0
								},
								"datasourceUid": "-100",
								"model": {
									"conditions": [
										{
											"evaluator": {
												"params": [
													3,
													3
												],
												"type": "lt"
											},
											"operator": {
												"type": "and"
											},
											"query": {
												"params": [ "kafka-count" ]
											},
											"reducer": {
												"params": [],
												"type": "last"
											},
											"type": "query"
										},
										{
											"evaluator": {
												"params": [
													3,
													3
												],
												"type": "lt"
											},
											"operator": {
												"type": "and"
											},
											"query": {
												"params": [ "zookeeper-count" ]
											},
											"reducer": {
												"params": [],
												"type": "last"
											},
											"type": "query"
										}
									],
									"datasource": "__expr__",
									"hide": false,
									"intervalMs": 1000,
									"maxDataPoints": 43200,
									"refId": "kafka-count-condition",
									"type": "classic_conditions"
								}
							}
						],
						"updated": "2021-12-01T12:30:00+10:00",
						"intervalSeconds": 60,
						"version": 10,
						"uid": "VMeCNvcnn",
						"namespace_uid": "zZltVzOnm",
						"namespace_id": 16,
						"rule_group": "Kafka Alert",
						"no_data_state": "NoData",
						"exec_err_state": "Alerting"
					}
				}
			]
		}
	]
}
`
)

func TestGetGrafanaAlertRules(t *testing.T) {
	server, client := gapiTestTools(t, 200, getGrafanaAlertRulesResponse)
	t.Cleanup(func() {
		server.Close()
	})

	resp, err := client.GetAlertRules("grafana")

	if err != nil {
		t.Error(err)
	}

	if len(resp["Namespace 1"]) <= 0 {
		t.Error("Issue")
	}
}
