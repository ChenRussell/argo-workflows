package metrics

import (
	"context"

	"github.com/argoproj/argo-workflows/v3/config"
	utilmetrics "github.com/argoproj/argo-workflows/v3/util/metrics"

	metricsdk "go.opentelemetry.io/otel/sdk/metric"
)

type Metrics struct {
	utilmetrics.Metrics
}

func New(ctx context.Context, serviceName, prometheusName string, config *utilmetrics.Config, extraOpts ...metricsdk.Option) (*Metrics, error) {
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
		Metrics: *m,
	}

	if err != nil {
		return nil, err
	}
	return metrics, nil
}

func GetServerConfig(config *config.Config) *utilmetrics.Config {
	// Metrics config
	modifiers := make(map[string]utilmetrics.Modifier)
	for name, modifier := range config.ServerMetricsConfig.Modifiers {
		modifiers[name] = utilmetrics.Modifier{
			Disabled:           modifier.Disabled,
			DisabledAttributes: modifier.DisabledAttributes,
			HistogramBuckets:   modifier.HistogramBuckets,
		}
	}

	metricsConfig := utilmetrics.Config{
		Enabled:      config.ServerMetricsConfig.Enabled == nil || *config.ServerMetricsConfig.Enabled,
		Path:         config.ServerMetricsConfig.Path,
		Port:         config.ServerMetricsConfig.Port,
		IgnoreErrors: config.ServerMetricsConfig.IgnoreErrors,
		Secure:       config.ServerMetricsConfig.GetSecure(true),
		Modifiers:    modifiers,
		Temporality:  config.ServerMetricsConfig.Temporality,
	}
	return &metricsConfig
}
