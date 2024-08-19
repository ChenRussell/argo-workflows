package metrics

import (
	"context"

	utilmetrics "github.com/argoproj/argo-workflows/v3/util/metrics"
)

const (
	namePodMissing = `pod_missing`
)

func addPodMissingCounter(_ context.Context, m *Metrics) error {
	return m.CreateInstrument(utilmetrics.Int64Counter,
		namePodMissing,
		"Incidents of pod missing.",
		"{pod}",
		utilmetrics.WithAsBuiltIn(),
	)
}

func (m *Metrics) incPodMissing(ctx context.Context, val int64, recentlyStarted bool, phase string) {
	m.AddInt(ctx, namePodMissing, val, utilmetrics.InstAttribs{
		{Name: utilmetrics.LabelRecentlyStarted, Value: recentlyStarted},
		{Name: utilmetrics.LabelNodePhase, Value: phase},
	})
}

func (m *Metrics) PodMissingEnsure(ctx context.Context, recentlyStarted bool, phase string) {
	m.incPodMissing(ctx, 0, recentlyStarted, phase)
}

func (m *Metrics) PodMissingInc(ctx context.Context, recentlyStarted bool, phase string) {
	m.incPodMissing(ctx, 1, recentlyStarted, phase)
}
