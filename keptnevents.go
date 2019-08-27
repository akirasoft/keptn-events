package keptnevents

import (
	"encoding/json"
)

// DeploymentFinishedEvent Keptn event payload for completed deployment
type DeploymentFinishedEvent struct {
	Githuborg          string `json:"githuborg"`
	Project            string `json:"project"`
	Teststrategy       string `json:"teststrategy"`
	Deploymentstrategy string `json:"deploymentstrategy"`
	Stage              string `json:"stage"`
	Service            string `json:"service"`
	Image              string `json:"image"`
	Tag                string `json:"tag"`
}

// EvaluationDoneEvent Keptn event payload for completed Pitometer evaluation Note: many elements are not strongly typed
type EvaluationDoneEvent struct {
	Githuborg          string `json:"githuborg"`
	Project            string `json:"project"`
	Teststrategy       string `json:"teststrategy"`
	Deploymentstrategy string `json:"deploymentstrategy"`
	Stage              string `json:"stage"`
	Service            string `json:"service"`
	Image              string `json:"image"`
	Tag                string `json:"tag"`
	Evaluationpassed   bool   `json:"evaluationpassed"`
	Evaluationdetails  struct {
		Options struct {
			TimeStart int `json:"timeStart"`
			TimeEnd   int `json:"timeEnd"`
		} `json:"options"`
		TotalScore int `json:"totalScore"`
		Objectives struct {
			Pass    int `json:"pass"`
			Warning int `json:"warning"`
		} `json:"objectives"`
		// Data coming back from Prometheus sources is not strongly typed
		// especially within indicatorResults
		IndicatorResults []struct {
			ID         string `json:"id"`
			Violations []struct {
				Value interface{} `json:"value"`
				// we need to  take the key as raw json and parse it later
				Key       json.RawMessage `json:"key"`
				Breach    string          `json:"breach"`
				Threshold interface{}     `json:"threshold"`
			} `json:"violations"`
			Score int `json:"score"`
		} `json:"indicatorResults"`
		Result string `json:"result"`
	} `json:"evaluationdetails"`
}

// TestsFinishedEvent Keptn event payload for completed tests
type TestsFinishedEvent struct {
	Githuborg          string `json:"githuborg"`
	Project            string `json:"project"`
	Teststrategy       string `json:"teststrategy"`
	Deploymentstrategy string `json:"deploymentstrategy"`
	Stage              string `json:"stage"`
	Service            string `json:"service"`
	Image              string `json:"image"`
	Tag                string `json:"tag"`
}

// NewArtifactEvent Keptn event payload for receipt of new build artifact
type NewArtifactEvent struct {
	Githuborg          string `json:"githuborg"`
	Project            string `json:"project"`
	Teststrategy       string `json:"teststrategy"`
	Deploymentstrategy string `json:"deploymentstrategy"`
	Stage              string `json:"stage"`
	Service            string `json:"service"`
	Image              string `json:"image"`
	Tag                string `json:"tag"`
}

// ProblemEvent Keptn event payload primarily created via Dynatrace webhook integration, ProblemDetails and ImpactedEntities will be raw json to be marshalled later
type ProblemEvent struct {
	State            string          `json:"State"`
	ProblemID        string          `json:"ProblemID"`
	PID              string          `json:"PID"`
	ProblemTitle     string          `json:"ProblemTitle"`
	ProblemDetails   json.RawMessage `json:"ProblemDetails"`
	ImpactedEntities json.RawMessage `json:"ImpactedEntities"`
	ImpactedEntity   string          `json:"ImpactedEntity"`
}

// RcvConfig stores configuration elements for cloudevents listener
type RcvConfig struct {
	// Port on which to listen for cloudevents
	Port int    `envconfig:"RCV_PORT" default:"8080"`
	Path string `envconfig:"RCV_PATH" default:"/"`
}
