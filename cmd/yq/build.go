package main

import (
	"get.porter.sh/mixin/yq/pkg/yq"
	"github.com/spf13/cobra"
)

func buildBuildCommand(m *yq.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Generate Dockerfile lines for the bundle invocation image",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Build()
		},
	}
	return cmd
}
