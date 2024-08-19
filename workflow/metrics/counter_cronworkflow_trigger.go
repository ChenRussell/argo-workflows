package metrics

import (
	"context"

	utilmetrics "github.com/argoproj/argo-workflows/v3/util/metrics"
)

const (
	nameCronTriggered = `cronworkflows_triggered_total`
)

func addCronWfTriggerCounter(_ context.Context, m *Metrics) error {
	return m.CreateInstrument(utilmetrics.Int64Counter,
		nameCronTriggered,
		"Total number of cron workflows triggered",
		"{cronworkflow}",
		utilmetrics.WithAsBuiltIn(),
	)
}

func (m *Metrics) CronWfTrigger(ctx context.Context, name, namespace string) {
	m.AddInt(ctx, nameCronTriggered, 1, utilmetrics.InstAttribs{
		{Name: utilmetrics.LabelCronWFName, Value: name},
		{Name: utilmetrics.LabelWorkflowNamespace, Value: namespace},
	})
}
