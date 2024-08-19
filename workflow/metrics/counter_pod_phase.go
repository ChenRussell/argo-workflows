package metrics

import (
	"context"

	utilmetrics "github.com/argoproj/argo-workflows/v3/util/metrics"
)

const (
	namePodPhase = `pods_total_count`
)

func addPodPhaseCounter(_ context.Context, m *Metrics) error {
	return m.CreateInstrument(utilmetrics.Int64Counter,
		namePodPhase,
		"Total number of Pods that have entered each phase",
		"{pod}",
		utilmetrics.WithAsBuiltIn(),
	)
}

func (m *Metrics) ChangePodPhase(ctx context.Context, phase, namespace string) {
	m.AddInt(ctx, namePodPhase, 1, utilmetrics.InstAttribs{
		{Name: utilmetrics.LabelPodPhase, Value: phase},
		{Name: utilmetrics.LabelPodNamespace, Value: namespace},
	})
}
