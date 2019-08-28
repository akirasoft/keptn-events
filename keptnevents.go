package keptnevents

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	cloudevents "github.com/cloudevents/sdk-go"
	keptnutils "github.com/keptn/go-utils/pkg/utils"
)

// ConfigurationChanged Keptn event payload for changed configuration
type ConfigurationChanged struct {
	Service            string `json:"service"`
	Image              string `json:"image"`
	Tag                string `json:"tag"`
	Project            string `json:"project"`
	Stage              string `json:"stage"`
	Githuborg          string `json:"githuborg"`
	Teststrategy       string `json:"teststrategy"`
	Deploymentstrategy string `json:"deploymentstrategy"`
}

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

// KeptnListener listens for Keptn events on the path and port defined via Rcv
func KeptnListener(Rcv RcvConfig) {

	ctx := context.Background()

	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithPort(Rcv.Port),
		cloudevents.WithPath(Rcv.Path),
	)
	if err != nil {
		log.Printf("failed to create transport, %v", err)
		return
	}
	c, err := cloudevents.NewClient(t)
	if err != nil {
		log.Printf("failed to create client, %v", err)
		return
	}

	log.Fatalf("failed to start receiver: %s", c.StartReceiver(ctx, KeptnHandler))
}

// KeptnHandler parses Keptn events and returns the Keptn event payload
func KeptnHandler(ctx context.Context, event cloudevents.Event) error {
	var shkeptncontext string
	event.Context.ExtensionAs("shkeptncontext", &shkeptncontext)

	logger := keptnutils.NewLogger(shkeptncontext, event.Context.GetID(), "jira-service")

	switch event.Type() {
	case "sh.keptn.events.configuration-changed":
		//return ConfigurationChanged
		log.Println("sh.keptn.events.configuration-changed")
	case "sh.keptn.events.deployment-finished":
		//return DeploymentFinishedEvent
		log.Println("sh.keptn.events.deployment-finished")
	case "sh.keptn.events.evaluation-done":
		//return EvaluationDoneEvent
		log.Println("sh.keptn.events.evaluation-done")
	case "sh.keptn.events.new-artifact":
		//return NewArtifactEvent
		log.Println("sh.keptn.events.new-artifact")
	case "sh.keptn.events.tests-finished":
		//return TestsFinishedEvent
		log.Println("sh.keptn.events.tests-finished")
	case "sh.keptn.events.problem":
		//return ProblemEvent
		log.Println("sh.keptn.events.problem")
	default:
		const errorMsg = "Received unexpected keptn event"
		logger.Error(errorMsg)
		return errors.New(errorMsg)
	}
	// data := &keptnevents.EvaluationDoneEvent{}
	// if err := event.DataAs(data); err != nil {
	// 	//TODO: replace with keptn logger
	// 	logger.Error(fmt.Sprintf("Got Data Error: %s", err.Error()))
	// 	return err
	// }

	return nil
}
