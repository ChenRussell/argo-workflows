package metrics

import (
	"context"
	"time"

	utilmetrics "github.com/argoproj/argo-workflows/v3/util/metrics"
)

const (
	nameWorkflowTemplateRuntime = `workflowtemplate_runtime`
)

func addWorkflowTemplateHistogram(_ context.Context, m *Metrics) error {
	return m.CreateInstrument(utilmetrics.Float64Histogram,
		nameWorkflowTemplateRuntime,
		"Duration of workflow template runs run through workflowTemplateRefs",
		"s",
		utilmetrics.WithAsBuiltIn(),
	)
}

func (m *Metrics) RecordWorkflowTemplateTime(ctx context.Context, duration time.Duration, name, namespace string, cluster bool) {
	m.Record(ctx, nameWorkflowTemplateRuntime, duration.Seconds(), templateLabels(name, namespace, cluster))
}
