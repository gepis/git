package git

import (
	"github.com/gepis/sm-gh-api/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdConfigRoot(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		SilenceErrors: true,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	cmd.SetOut(f.IOStreams.Out)
	cmd.SetErr(f.IOStreams.ErrOut)

	cmdutil.DisableAuthCheck(cmd)

	return cmd
}
