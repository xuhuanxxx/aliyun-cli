package cli

import (
	"bytes"
	"github.com/aliyun/aliyun-cli/v3/i18n"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestXDistributionIdentity(t *testing.T) {
	oldName, oldVersion, oldCommit := CommandName, Version, SourceCommit
	t.Cleanup(func() { CommandName, Version, SourceCommit = oldName, oldVersion, oldCommit })
	CommandName, Version, SourceCommit = "aliyunx", "3.5.1-x.1", "abc123"
	root := &Command{Name: CommandName, Usage: "aliyun <command>"}
	child := &Command{Name: "configure", Sample: "aliyun configure", Long: i18n.T("Run aliyun configure", "运行 aliyun configure")}
	root.AddSubCommand(child)
	BrandCommandHelp(root)
	require.Equal(t, "aliyunx <command>", root.Usage)
	require.Equal(t, "aliyunx configure", child.Sample)
	require.Equal(t, "Run aliyunx configure", child.Long.Get("en"))
	require.Equal(t, "aliyunx help", (*InvalidFlagError)(nil).AgentHelpCommand())
	out := &bytes.Buffer{}
	require.NoError(t, NewVersionCommand().Run(NewCommandContext(out, out), nil))
	require.Equal(t, "3.5.1-x.1 (abc123)\n", out.String())
}
