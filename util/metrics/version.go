package metrics

import (
	"context"

	"github.com/argoproj/argo-workflows/v3"
)

func AddVersion(ctx context.Context, m *Metrics) error {
	const nameVersion = `version`
	err := m.CreateInstrument(Int64Counter,
		nameVersion,
		"Build metadata for this Controller",
		"{unused}",
		WithAsBuiltIn(),
	)
	if err != nil {
		return err
	}

	version := argo.GetVersion()
	m.AddInt(ctx, nameVersion, 1, InstAttribs{
		{Name: LabelBuildVersion, Value: version.Version},
		{Name: LabelBuildPlatform, Value: version.Platform},
		{Name: LabelBuildGoVersion, Value: version.GoVersion},
		{Name: LabelBuildDate, Value: version.BuildDate},
		{Name: LabelBuildCompiler, Value: version.Compiler},
		{Name: LabelBuildGitCommit, Value: version.GitCommit},
		{Name: LabelBuildGitTreeState, Value: version.GitTreeState},
		{Name: LabelBuildGitTag, Value: version.GitTag},
	})
	return nil
}
