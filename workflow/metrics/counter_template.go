package metrics

import (
	"context"

	utilmetrics "github.com/argoproj/argo-workflows/v3/util/metrics"
)

const (
	nameWFTemplateTriggered = `workflowtemplate_triggered_total`
)

func addWorkflowTemplateCounter(_ context.Context, m *Metrics) error {
	return m.CreateInstrument(utilmetrics.Int64Counter,
		nameWFTemplateTriggered,
		"Total number of workflow templates triggered by workflowTemplateRef",
		"{workflow_template}",
		utilmetrics.WithAsBuiltIn(),
	)
}

func templateLabels(name, namespace string, cluster bool) utilmetrics.InstAttribs {
	return utilmetrics.InstAttribs{
		{Name: utilmetrics.LabelTemplateName, Value: name},
		{Name: utilmetrics.LabelTemplateNamespace, Value: namespace},
		{Name: utilmetrics.LabelTemplateCluster, Value: cluster},
	}
}

func (m *Metrics) CountWorkflowTemplate(ctx context.Context, phase MetricWorkflowPhase, name, namespace string, cluster bool) {
	labels := templateLabels(name, namespace, cluster)
	labels = append(labels, utilmetrics.InstAttrib{Name: utilmetrics.LabelWorkflowPhase, Value: string(phase)})

	m.AddInt(ctx, nameWFTemplateTriggered, 1, labels)
}
