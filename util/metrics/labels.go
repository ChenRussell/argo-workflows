package metrics

const (
	LabelBuildVersion      string = `version`
	LabelBuildPlatform     string = `platform`
	LabelBuildGoVersion    string = `go_version`
	LabelBuildDate         string = `build_date`
	LabelBuildCompiler     string = `compiler`
	LabelBuildGitCommit    string = `git_commit`
	LabelBuildGitTreeState string = `git_treestate`
	LabelBuildGitTag       string = `git_tag`

	LabelCronWFName string = `name`

	LabelErrorCause string = "cause"

	LabelLogLevel string = `level`

	LabelNodePhase string = `node_phase`

	LabelPodPhase         string = `phase`
	LabelPodNamespace     string = `namespace`
	LabelPodPendingReason string = `reason`

	LabelQueueName string = `queue_name`

	LabelRecentlyStarted string = `recently_started`

	LabelRequestKind = `kind`
	LabelRequestVerb = `verb`
	LabelRequestCode = `status_code`

	LabelTemplateName      string = `name`
	LabelTemplateNamespace string = `namespace`
	LabelTemplateCluster   string = `cluster_scope`

	LabelWorkerType string = `worker_type`

	LabelWorkflowNamespace string = `namespace`
	LabelWorkflowPhase     string = `phase`
	LabelWorkflowStatus           = `status`
	LabelWorkflowType             = `type`
)
