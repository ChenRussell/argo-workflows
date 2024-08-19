package metrics

import (
	"context"

	utilmetrics "github.com/argoproj/argo-workflows/v3/util/metrics"

	"go.opentelemetry.io/otel/metric"
)

type IsLeaderCallback func() bool

type leaderGauge struct {
	callback IsLeaderCallback
	gauge    *utilmetrics.Instrument
}

func addIsLeader(ctx context.Context, m *Metrics) error {
	const nameLeader = `is_leader`
	err := m.CreateInstrument(utilmetrics.Int64ObservableGauge,
		nameLeader,
		"Emits 1 if leader, 0 otherwise. Always 1 if leader election is disabled.",
		"{leader}",
		utilmetrics.WithAsBuiltIn(),
	)
	if err != nil {
		return err
	}
	if m.callbacks.IsLeader == nil {
		return nil
	}
	lGauge := leaderGauge{
		callback: m.callbacks.IsLeader,
		gauge:    m.AllInstruments[nameLeader],
	}
	return m.AllInstruments[nameLeader].RegisterCallback(&m.Metrics, lGauge.update)
}

func (l *leaderGauge) update(_ context.Context, o metric.Observer) error {
	var val int64 = 0
	if l.callback() {
		val = 1
	}
	l.gauge.ObserveInt(o, val, utilmetrics.InstAttribs{})
	return nil
}
