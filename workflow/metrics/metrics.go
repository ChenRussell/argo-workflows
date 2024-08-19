package metrics

import (
	"context"

	utilmetrics "github.com/argoproj/argo-workflows/v3/util/metrics"

	metricsdk "go.opentelemetry.io/otel/sdk/metric"
)

type Metrics struct {
	utilmetrics.Metrics

	callbacks         Callbacks
	realtimeWorkflows map[string][]realtimeTracker
}

func New(ctx context.Context, serviceName, prometheusName string, config *utilmetrics.Config, callbacks Callbacks, extraOpts ...metricsdk.Option) (*Metrics, error) {
	m, err := utilmetrics.New(ctx, serviceName, prometheusName, config, extraOpts...)
	if err != nil {
		return nil, err
	}

	err = m.Populate(ctx,
		utilmetrics.AddVersion,
	)
	if err != nil {
		return nil, err
	}

	metrics := &Metrics{
		Metrics:           *m,
		callbacks:         callbacks,
		realtimeWorkflows: make(map[string][]realtimeTracker),
	}

	err = metrics.populate(ctx,
		addIsLeader,
		addPodPhaseGauge,
		addPodPhaseCounter,
		addPodMissingCounter,
		addPodPendingCounter,
		addWorkflowPhaseGauge,
		addCronWfTriggerCounter,
		addWorkflowPhaseCounter,
		addWorkflowTemplateCounter,
		addWorkflowTemplateHistogram,
		addOperationDurationHistogram,
		addErrorCounter,
		addLogCounter,
		addK8sRequests,
		addWorkflowConditionGauge,
		addWorkQueueMetrics,
	)
	if err != nil {
		return nil, err
	}

	go metrics.customMetricsGC(ctx, config.TTL)

	return metrics, nil
}

type addMetric func(context.Context, *Metrics) error

func (m *Metrics) populate(ctx context.Context, adders ...addMetric) error {
	for _, adder := range adders {
		if err := adder(ctx, m); err != nil {
			return err
		}
	}
	return nil
}
