package metrics

import (
	"context"

	utilmetrics "github.com/argoproj/argo-workflows/v3/util/metrics"

	"go.opentelemetry.io/otel/metric"
)

// WorkflowPhaseCallback is the function prototype to provide this gauge with the phase of the pods
type WorkflowPhaseCallback func() map[string]int64

type workflowPhaseGauge struct {
	callback WorkflowPhaseCallback
	gauge    *utilmetrics.Instrument
}

func addWorkflowPhaseGauge(_ context.Context, m *Metrics) error {
	const nameWorkflowPhaseGauge = `gauge`
	err := m.CreateInstrument(utilmetrics.Int64ObservableGauge,
		nameWorkflowPhaseGauge,
		"number of Workflows currently accessible by the controller by status",
		"{workflow}",
		utilmetrics.WithAsBuiltIn(),
	)
	if err != nil {
		return err
	}

	if m.callbacks.WorkflowPhase != nil {
		wfpGauge := workflowPhaseGauge{
			callback: m.callbacks.WorkflowPhase,
			gauge:    m.AllInstruments[nameWorkflowPhaseGauge],
		}
		return m.AllInstruments[nameWorkflowPhaseGauge].RegisterCallback(&m.Metrics, wfpGauge.update)
	}
	return nil
	// TODO init all phases?
}

func (p *workflowPhaseGauge) update(_ context.Context, o metric.Observer) error {
	phases := p.callback()
	for phase, val := range phases {
		p.gauge.ObserveInt(o, val, utilmetrics.InstAttribs{{Name: utilmetrics.LabelWorkflowStatus, Value: phase}})
	}
	return nil
}
