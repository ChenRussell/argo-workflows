package metrics

import (
	"context"

	utilmetrics "github.com/argoproj/argo-workflows/v3/util/metrics"
)

type ErrorCause string

const (
	nameErrorCount                                   = `error_count`
	ErrorCauseOperationPanic              ErrorCause = "OperationPanic"
	ErrorCauseCronWorkflowSubmissionError ErrorCause = "CronWorkflowSubmissionError"
	ErrorCauseCronWorkflowSpecError       ErrorCause = "CronWorkflowSpecError"
)

func addErrorCounter(ctx context.Context, m *Metrics) error {
	err := m.CreateInstrument(utilmetrics.Int64Counter,
		nameErrorCount,
		"Number of errors encountered by the controller by cause",
		"{error}",
		utilmetrics.WithAsBuiltIn(),
	)
	if err != nil {
		return err
	}
	// Initialise all values to zero
	for _, cause := range []ErrorCause{ErrorCauseOperationPanic, ErrorCauseCronWorkflowSubmissionError, ErrorCauseCronWorkflowSpecError} {
		m.AddInt(ctx, nameErrorCount, 0, utilmetrics.InstAttribs{{Name: utilmetrics.LabelErrorCause, Value: string(cause)}})
	}
	return nil
}

func (m *Metrics) OperationPanic(ctx context.Context) {
	m.AddInt(ctx, nameErrorCount, 1, utilmetrics.InstAttribs{{Name: utilmetrics.LabelErrorCause, Value: string(ErrorCauseOperationPanic)}})
}

func (m *Metrics) CronWorkflowSubmissionError(ctx context.Context) {
	m.AddInt(ctx, nameErrorCount, 1, utilmetrics.InstAttribs{{Name: utilmetrics.LabelErrorCause, Value: string(ErrorCauseCronWorkflowSubmissionError)}})
}

func (m *Metrics) CronWorkflowSpecError(ctx context.Context) {
	m.AddInt(ctx, nameErrorCount, 1, utilmetrics.InstAttribs{{Name: utilmetrics.LabelErrorCause, Value: string(ErrorCauseCronWorkflowSpecError)}})
}
