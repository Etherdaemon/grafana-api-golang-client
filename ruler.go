package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

type RuleGroupConfig struct {
	Interval string  `json:"interval"`
	Name     string  `json:"name"`
	Rules    []*Rule `json:"rules,omitempty"`
}

type Rule struct {
	Alert        string                 `json:"alert"`
	Annotations  map[string]interface{} `json:"annotations"`
	Expression   string                 `json:"expr"`
	For          string                 `json:"for"`
	GrafanaAlert GrafanaAlertRule       `json:"grafana_alert"`
	Labels       map[string]interface{} `json:"labels"`
	Record       string                 `json:"record"`
}

type GrafanaAlertRule struct {
	Condition       string        `json:"condition"`
	Data            []*AlertQuery `json:"data"`
	ExecErrorState  string        `json:"exec_err_state"`
	ID              int64         `json:"id"`
	IntervalSeconds int64         `json:"intervalSeconds"`
	NamespaceID     int64         `json:"namespace_id"`
	NamespaceUID    string        `json:"namespace_uid"`
	NoDataState     string        `json:"no_data_state"`
	OrgID           int64         `json:"orgId"`
	RuleGroup       string        `json:"rule_group"`
	Title           string        `json:"title"`
	UID             string        `json:"uid"`
	Updated         string        `json:"updated"`
	Version         int64         `json:"version"`
}

type AlertQuery struct {
	// Grafana data source unique identifier; it should be '-100' for a Server Side Expression operation.
	DatasourceUID     string                 `json:"datasourceUid"`
	Model             map[string]interface{} `json:"model"`
	QueryType         string                 `json:"queryType"`
	RefID             string                 `json:"refId"`
	RelativeTimeRange AlertRelativeTimeRange `json:"relativeTimeRange"`
}

type AlertRelativeTimeRange struct {
	From int64 `json:"from"`
	To   int64 `json:"to"`
}

// Fetch all alert rules for recipient in all namespaces
// Recipient should be "grafana" for requests to be handled by grafana
// and the numeric datasource id for requests to be forwarded to a datasource
func (c *Client) GetAllAlertRules(recipient string) (map[string][]*RuleGroupConfig, error) {
	rules := make(map[string][]*RuleGroupConfig)

	request_uri := fmt.Sprintf("/api/ruler/%s/api/v1/rules", recipient)
	err := c.request("GET", request_uri, nil, nil, &rules)
	if err != nil {
		return rules, err
	}

	return rules, err
}

// Creates a new Namespace Alert Rule Group Config
func (c *Client) NewRuleGroupConfig(recipient string, namespace string, ruleGroupConfig RuleGroupConfig) (*RuleGroupConfig, error) {
	body, err := json.Marshal(ruleGroupConfig)
	if err != nil {
		return nil, err
	}

	rgc := &RuleGroupConfig{}
	request_uri := fmt.Sprintf("/api/ruler/%s/api/v1/rules/%s", recipient, url.QueryEscape(namespace))
	err = c.request("POST", request_uri, nil, bytes.NewBuffer(body), &rgc)
	if err != nil {
		return nil, err
	}

	return rgc, err
}
