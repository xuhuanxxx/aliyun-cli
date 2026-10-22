package main

import (
	"github.com/aliyun/aliyun-cli/v3/cli"
	"github.com/aliyun/aliyun-cli/v3/config"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

func TestXRootIdentity(t *testing.T) {
	old := cli.CommandName
	cli.CommandName = "aliyunx"
	t.Cleanup(func() { cli.CommandName = old })
	root := newRootCommand(config.Profile{}, io.Discard)
	require.Equal(t, "aliyunx", root.Name)
	require.Contains(t, root.Usage, "aliyunx ")
}
