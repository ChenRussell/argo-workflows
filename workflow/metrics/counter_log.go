package metrics

import (
	"context"

	utilmetrics "github.com/argoproj/argo-workflows/v3/util/metrics"

	log "github.com/sirupsen/logrus"
)

type logMetric struct {
	counter *utilmetrics.Instrument
}

func addLogCounter(ctx context.Context, m *Metrics) error {
	const nameLogMessages = `log_messages`
	err := m.CreateInstrument(utilmetrics.Int64Counter,
		nameLogMessages,
		"Total number of log messages.",
		"{message}",
		utilmetrics.WithAsBuiltIn(),
	)
	lm := logMetric{
		counter: m.AllInstruments[nameLogMessages],
	}
	log.AddHook(lm)
	for _, level := range lm.Levels() {
		m.AddInt(ctx, nameLogMessages, 0, utilmetrics.InstAttribs{
			{Name: utilmetrics.LabelLogLevel, Value: level.String()},
		})
	}

	return err
}

func (m logMetric) Levels() []log.Level {
	return []log.Level{log.InfoLevel, log.WarnLevel, log.ErrorLevel}
}

func (m logMetric) Fire(entry *log.Entry) error {
	(*m.counter).AddInt(entry.Context, 1, utilmetrics.InstAttribs{
		{Name: utilmetrics.LabelLogLevel, Value: entry.Level.String()},
	})
	return nil
}
